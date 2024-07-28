package handlers

import (
	"net/http"
	"neurokin/views/dashboard"
)

func HandleDashboard(w http.ResponseWriter, r *http.Request) error {
  return Render(w, r, dashboard.Index())
}
