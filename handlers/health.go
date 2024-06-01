package handlers

import (
  u "neurokin/util"
	"net/http"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) error {
	return u.WriteJSON(w, http.StatusOK, "Server OK")
}
