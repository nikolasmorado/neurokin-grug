package auth

import (
	"fmt"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt/v4"

	t "neurokin/types"
	u "neurokin/util"
)

func WithJWT(handlerFunc http.HandlerFunc, s t.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		token, err := validateJWT(tokenString)

		if err != nil {
			u.WriteJSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		if !token.Valid {
			u.WriteJSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// claims := token.Claims.(jwt.MapClaims)

		// uID := claims["accountId"].(float64)

		// _, err = s.GetAccount(int(uID))

		// if err != nil {
		// 	u.WriteJSON(w, http.StatusUnauthorized, "Unauthorized")
		// 	return
		// }

		handlerFunc(w, r)
	}
}

func GetIdFromJWT(tokenString string) (int, error) {
	token, err := validateJWT(tokenString)

	if err != nil {
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)

	return int(claims["accountId"].(float64)), nil
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("SUPA_JWT")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}
