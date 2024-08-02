package handlers

import (
	"net/http"
	"neurokin/views/verify"
)

func HandleVerify(w http.ResponseWriter, r *http.Request) error {
  return Render(w, r, verify.Index())
}
