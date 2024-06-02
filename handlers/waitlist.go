package handlers

import (
	"net/http"
	"neurokin/views/waitlist"
)

func HandlePostWaitlist(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	// email := r.Form.Get("email")

	return Render(w, r, waitlist.Index())
}
