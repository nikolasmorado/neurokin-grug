package handlers

import (
	"net/http"
	"neurokin/views/quiz"
)

func HandleQuiz(w http.ResponseWriter, r *http.Request) error {
  return Render(w, r, quiz.Index())
}
