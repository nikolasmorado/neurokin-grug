package handlers

import (
	"net/http"

	a "neurokin/auth"
	q "neurokin/quizzes"
	t "neurokin/types"
	u "neurokin/util"
	"neurokin/views/quiz"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HandleQuizSlug(deps *u.HandlerDependencies) t.HTTPHandler {
	s := deps.Store

	return func(w http.ResponseWriter, r *http.Request) error {
		switch r.Method {
		case "GET":
			return RenderQuizSlug(w, r, s)
		case "POST":
      return StartQuiz(w, r, s)
    case "PUT":
      return ProgressQuiz(w, r, s)
		default:
			return RenderQuizSlug(w, r, s)
		}
	}
}

func RenderQuizSlug(w http.ResponseWriter, r *http.Request, s t.Storage) error {
	aC, err := r.Cookie("Authorization")

	if err != nil {
		return nil
	}

	id, err := a.GetIDFromJWT(aC.Value)

	if err != nil {
		return nil
	}

	slug := chi.URLParam(r, "slug")

	qz, err := q.GetQuiz(slug)

	if err != nil {
		return nil
	}

  qa, err := s.GetQuizAnswer(slug, id)

  if err != nil {
    return nil
  }


  if qa == nil {
    return Render(w, r, quiz.Question(qz.Name, qz.Slug))
  }

  qIdx := len(qa.DecodedAnswers)

  return Render(w, r, quiz.QuestionSkip(
    qz.Questions[qIdx].Question,
    strconv.Itoa(qIdx + 1),
    strconv.Itoa(len(qz.Questions)),
    qz.Slug,
    qz.Questions[qIdx].Answers,
  ))
}

func StartQuiz(w http.ResponseWriter, r *http.Request, s t.Storage) error {
  aC, err := r.Cookie("Authorization")

  if err != nil {
    return nil
  }

  id, err := a.GetIDFromJWT(aC.Value)

  if err != nil {
    return nil
  }

  slug := chi.URLParam(r, "slug")

  qz, err := q.GetQuiz(slug)

  if err != nil {
    return nil
  }

  err = s.CreateQuizAnswer(id.String(), slug)

  if err != nil {
    return nil
  }

  return Render(w, r, quiz.Qbox(
    qz.Questions[0].Question,
    "1",
    strconv.Itoa(len(qz.Questions)),
    qz.Slug,
    qz.Questions[0].Answers,
  ))
}

func ProgressQuiz(w http.ResponseWriter, r *http.Request, s t.Storage) error {
	aC, err := r.Cookie("Authorization")

	if err != nil {
		return nil
	}

	id, err := a.GetIDFromJWT(aC.Value)

	if err != nil {
		return nil
	}

	slug := chi.URLParam(r, "slug")

	qz, err := q.GetQuiz(slug)

	if err != nil {
		return nil
	}

  qa, err := s.GetQuizAnswer(slug, id)

  if err != nil {
    return nil
  }

  if qa == nil {
    return Render(w, r, quiz.Qbox(
      qz.Questions[0].Question,
      "1",
      strconv.Itoa(len(qz.Questions)),
      qz.Slug,
      qz.Questions[0].Answers,
    ))
  }

  selectedOption := r.FormValue("selectedOption")
  additionalContext := r.FormValue("additionalContext")
  bookmarked := r.FormValue("bookmarked")
  
  qa.DecodedAnswers = append(qa.DecodedAnswers, t.Answer{
    QuestionId: qz.Questions[len(qa.DecodedAnswers)].QuestionId,
    Answer: selectedOption,
    Context: additionalContext,
    Bookmarked: bookmarked == "true",
  })

  err = s.UpdateQuizAnswer(qa)

  qIdx := len(qa.DecodedAnswers)

	return Render(w, r, quiz.Qbox(
    qz.Questions[qIdx].Question,
    strconv.Itoa(qIdx + 1),
		strconv.Itoa(len(qz.Questions)),
    qz.Slug,
    qz.Questions[qIdx].Answers,
	))
}
