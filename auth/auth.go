package auth

import (
	"fmt"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	t "neurokin/types"
	u "neurokin/util"
)

func WithJWT(handlerFunc http.HandlerFunc, s t.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("Authorization")
		if err != nil {
			u.WriteJSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		tokenString := cookie.Value

		token, err := validateJWT(tokenString)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if !token.Valid {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		email := claims["email"].(string)

		_, err = s.GetAccountByEmail(email)

		if err != nil {
			u.WriteJSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		handlerFunc(w, r)
	}
}

func GetEmailFromJWT(tokenString string) (string, error) {
	token, err := validateJWT(tokenString)

	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims["email"].(string), nil
}

func GetIDFromJWT(tokenString string) (uuid.UUID, error) {
	token, err := validateJWT(tokenString)

	if err != nil {
		return uuid.UUID{}, err
	}

	claims := token.Claims.(jwt.MapClaims)

	id, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
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
