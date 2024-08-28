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
			if _, exists := bqns[a.QuestionId]; exists {
				return Render(w, r, quiz.QuestionSkip(
					a.Question,
					strconv.Itoa(i+1),
					strconv.Itoa(len(qz.Questions)),
					qz.Slug,
					a.QuestionId,
					a.Answers,
				))
			}
		}
	}

	return Render(w, r, quiz.QuestionSkip(
		qz.Questions[qIdx].Question,
		strconv.Itoa(qIdx+1),
		strconv.Itoa(len(qz.Questions)),
		qz.Slug,
		qz.Questions[qIdx].QuestionId,
		qz.Questions[qIdx].Answers,
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
		qz.Questions[0].QuestionId,
		qz.Questions[0].Answers,
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

	if qa == nil {
		return Render(w, r, quiz.Qbox(
			qz.Questions[0].Question,
			"1",
			strconv.Itoa(len(qz.Questions)),
			qz.Slug,
			qz.Questions[0].QuestionId,
			qz.Questions[0].Answers,
		))
	}

	selectedOption := r.FormValue("selectedOption")
	additionalContext := r.FormValue("additionalContext")
	bookmarked := r.FormValue("bookmarked") == "true"
	pQIdx := r.FormValue("questionId")

	// Use a map for faster lookups and replacement of existing answers
	answerMap := make(map[string]*t.Answer)
	for i := range qa.DecodedAnswers {
		answerMap[qa.DecodedAnswers[i].QuestionId] = &qa.DecodedAnswers[i]
	}

	// Replace or add the answer
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

	err = s.UpdateQuizAnswer(qa)
	if err != nil {
		return err
	}

	seenAll := len(qa.DecodedAnswers) >= len(qz.Questions)

	if seenAll {
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

			return Render(w, r, quiz.Meta(
				qz.Name,
				qz.Slug,
				"Return to Dashboard ->",
				true,
			))
		}

		for i, a := range qz.Questions {
			if _, exists := bqns[a.QuestionId]; exists {
				return Render(w, r, quiz.Qbox(
					a.Question,
					strconv.Itoa(i+1),
					strconv.Itoa(len(qz.Questions)),
					qz.Slug,
					a.QuestionId,
					a.Answers,
				))
			}
		}

		return nil
	}

	qIdx := len(qa.DecodedAnswers)
	return Render(w, r, quiz.Qbox(
		qz.Questions[qIdx].Question,
		strconv.Itoa(qIdx+1),
		strconv.Itoa(len(qz.Questions)),
		qz.Slug,
		qz.Questions[qIdx].QuestionId,
		qz.Questions[qIdx].Answers,
	))
}

