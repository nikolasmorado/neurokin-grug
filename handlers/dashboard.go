package handlers

import (
	"net/http"
	"neurokin/views/dashboard"

	a "neurokin/auth"
)

func HandleDashboard(w http.ResponseWriter, r *http.Request) error {
  
  aC, err := r.Cookie("Authorization")

  if err != nil {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return nil
  }

  email, err := a.GetEmailFromJWT(aC.Value)

  if err != nil {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return nil
  }

  return Render(w, r, dashboard.Index(email))
}
