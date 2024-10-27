package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"neurokin/views/dashboard"
	"os"
	"path/filepath"

	a "neurokin/auth"
	t "neurokin/types"
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


  tasksPath, err := filepath.Abs("tasks.json")
  if err != nil {
    fmt.Errorf("Tasks file not found")
  }

  var tasks []t.Task 
  tasksData, err := os.ReadFile(tasksPath)
  if err := json.Unmarshal(tasksData, &tasks); err != nil {
    fmt.Errorf("Error loading tasks file")
  }

  return Render(w, r, dashboard.Index(email, tasks))
}
