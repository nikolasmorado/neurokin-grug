package handlers

import (
	"net/http"
	t "neurokin/types"
	u "neurokin/util"
  q "neurokin/quizzes"
	"neurokin/views/quiz"

	"github.com/go-chi/chi/v5"
	// q "neurokin/quizzes"
)

func HandleQuiz(w http.ResponseWriter, r *http.Request) error {

	return Render(w, r, quiz.Meta("", ""))
}


func HandleQuizSlug(deps *u.HandlerDependencies) t.HTTPHandler {
	s := deps.Store

	return func(w http.ResponseWriter, r *http.Request) error {
		switch r.Method {
		case "GET":
			return RenderQuizSlug(w, r, s)
    case "POST":
      return ProgressQuiz(w, r, s)
		default:
			return HandleQuiz(w, r)
		}
	}
}

func RenderQuizSlug(w http.ResponseWriter, r *http.Request, s t.Storage) error {
  slug := chi.URLParam(r, "slug")

  qz, err := q.GetQuiz(slug)

  if err != nil {
    http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
    return nil
  }

	return Render(w, r, quiz.Question(qz.Name, qz.Slug))
}

func ProgressQuiz(w http.ResponseWriter, r *http.Request, s t.Storage) error {
  slug := chi.URLParam(r, "slug")

  _, err := q.GetQuiz(slug)

  if err != nil {
    return nil
  }

  return Render(w, r, quiz.Qbox(
    "0",
    "10",
    "10",
    []string{"a", "b", "c"},
  ))
}
