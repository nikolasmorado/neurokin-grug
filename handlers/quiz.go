package handlers

import (
	"net/http"
	"strconv"

	a "neurokin/auth"
	q "neurokin/quizzes"
	t "neurokin/types"
	u "neurokin/util"
	"neurokin/views/quiz"

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
		return err
	}

	id, err := a.GetIDFromJWT(aC.Value)
	if err != nil {
		return err
	}

	slug := chi.URLParam(r, "slug")

	qz, err := q.GetQuiz(slug)
	if err != nil {
		return err
	}

	qa, err := s.GetQuizAnswer(slug, id)
	if err != nil {
		return err
	}

	// Check if the quiz is already completed
	if qa != nil && qa.Completed {
		return Render(w, r, quiz.Question(
			qz.Name,
			qz.Slug,
			"Return to Dashboard ->",
			true,
		))
	}

	if qa == nil {
		return Render(w, r, quiz.Question(qz.Name, qz.Slug, "Start Quiz ->", false))
	}

	qIdx := len(qa.DecodedAnswers)

	if qIdx >= len(qz.Questions) {
		// Use a map for faster lookups of bookmarked or unanswered questions
		bqns := make(map[string]t.Answer)

		for _, a := range qa.DecodedAnswers {
			if a.Bookmarked || a.Answer == "" {
				bqns[a.QuestionId] = a
			}
		}

		if len(bqns) == 0 {
			qa.Completed = true
			err = s.UpdateQuizAnswer(qa)
			if err != nil {
				return err
			}

			return Render(w, r, quiz.Question(
				qz.Name,
				qz.Slug,
				"Return to Dashboard ->",
				true,
			))
		}

		for i, a := range qz.Questions {
			if _, exists := bqns[a.Id]; exists {
				return Render(w, r, quiz.QuestionSkip(
					a.Question,
					strconv.Itoa(i+1),
					strconv.Itoa(len(qz.Questions)),
					qz.Slug,
					a.Id,
					a.Answers,
          string(a.Type),
				))
			}
		}
	}

	return Render(w, r, quiz.QuestionSkip(
		qz.Questions[qIdx].Question,
		strconv.Itoa(qIdx+1),
		strconv.Itoa(len(qz.Questions)),
		qz.Slug,
		qz.Questions[qIdx].Id,
		qz.Questions[qIdx].Answers,
    string(qz.Questions[qIdx].Type),
	))
}

func StartQuiz(w http.ResponseWriter, r *http.Request, s t.Storage) error {
	aC, err := r.Cookie("Authorization")
	if err != nil {
		return err
	}

	id, err := a.GetIDFromJWT(aC.Value)
	if err != nil {
		return err
	}

	slug := chi.URLParam(r, "slug")

	qz, err := q.GetQuiz(slug)
	if err != nil {
		return err
	}

	err = s.CreateQuizAnswer(id.String(), slug)
	if err != nil {
		return err
	}

	return Render(w, r, quiz.Qbox(
		qz.Questions[0].Question,
		"1",
		strconv.Itoa(len(qz.Questions)),
		qz.Slug,
		qz.Questions[0].Id,
		qz.Questions[0].Answers,
		string(qz.Questions[0].Type),
	))
}

func ProgressQuiz(w http.ResponseWriter, r *http.Request, s t.Storage) error {
	aC, err := r.Cookie("Authorization")
	if err != nil {
		return err
	}

	id, err := a.GetIDFromJWT(aC.Value)
	if err != nil {
		return err
	}

	slug := chi.URLParam(r, "slug")

	qz, err := q.GetQuiz(slug)
	if err != nil {
		return err
	}

	qa, err := s.GetQuizAnswer(slug, id)
	if err != nil {
		return err
	}

	// Start quiz if no quiz answer exists
	if qa == nil {
		return Render(w, r, quiz.Qbox(
			qz.Questions[0].Question,
			"1",
			strconv.Itoa(len(qz.Questions)),
			qz.Slug,
			qz.Questions[0].Id,
			qz.Questions[0].Answers,
      string(qz.Questions[0].Type),
		))
	}

	// Retrieve form values
	selectedOption := r.FormValue("selectedOption")
	additionalContext := r.FormValue("additionalContext")
	bookmarked := r.FormValue("bookmarked") == "true"
	pQIdx := r.FormValue("questionId")

	// Use a map for faster lookups
	answerMap := make(map[string]*t.Answer)
	for i := range qa.DecodedAnswers {
		answerMap[qa.DecodedAnswers[i].QuestionId] = &qa.DecodedAnswers[i]
	}

	// Update or add the answer
	if answer, exists := answerMap[pQIdx]; exists {
		answer.Answer = selectedOption
		answer.Context = additionalContext
		answer.Bookmarked = bookmarked
	} else {
		qa.DecodedAnswers = append(qa.DecodedAnswers, t.Answer{
			QuestionId: pQIdx,
			Answer:     selectedOption,
			Context:    additionalContext,
			Bookmarked: bookmarked,
		})
	}

	// Update the quiz answer in storage
	err = s.UpdateQuizAnswer(qa)
	if err != nil {
		return err
	}

	// Check if all questions have been answered
	seenAll := len(qa.DecodedAnswers) >= len(qz.Questions)

	if seenAll {
		// Check for any bookmarked or unanswered questions
		bqns := make(map[string]t.Answer)
		for _, a := range qa.DecodedAnswers {
			if a.Bookmarked || a.Answer == "" {
				bqns[a.QuestionId] = a
			}
		}

		// If no bookmarked/unanswered questions, mark as completed and return to dashboard
		if len(bqns) == 0 {
			qa.Completed = true
			err = s.UpdateQuizAnswer(qa)
			if err != nil {
				return err
			}

			// Render completion message
			return Render(w, r, quiz.Meta(
				qz.Name,
				qz.Slug,
				"Quiz Completed! Return to Dashboard ->",
				true,
			))
		}

		// Render the next bookmarked question if any
		for i, a := range qz.Questions {
			if _, exists := bqns[a.Id]; exists {
				return Render(w, r, quiz.Qbox(
					a.Question,
					strconv.Itoa(i+1),
					strconv.Itoa(len(qz.Questions)),
					qz.Slug,
					a.Id,
					a.Answers,
          string(a.Type),
				))
			}
		}

		return nil
	}

	// Render the next question
	qIdx := len(qa.DecodedAnswers)
	return Render(w, r, quiz.Qbox(
		qz.Questions[qIdx].Question,
		strconv.Itoa(qIdx+1),
		strconv.Itoa(len(qz.Questions)),
		qz.Slug,
		qz.Questions[qIdx].Id,
		qz.Questions[qIdx].Answers,
		string(qz.Questions[qIdx].Type),
	))
}
