package handlers

import (
	"fmt"
	"net/http"
	t "neurokin/types"
	u "neurokin/util"
	"neurokin/views/waitlist"
)

func HandleWaitlist(deps *u.HandlerDependencies) t.HTTPHandler {
	s := deps.Store
	return func(w http.ResponseWriter, r *http.Request) error {
		if err := r.ParseForm(); err != nil {
			return err
		}

		email := r.Form.Get("email")

    fmt.Println("email: ", email)

    s.CreateWaitlist(email)

		return Render(w, r, waitlist.Index())
	}
}

func HandlePostWaitlist(w http.ResponseWriter, r *http.Request, s t.Storage) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	return Render(w, r, waitlist.Index())
}
