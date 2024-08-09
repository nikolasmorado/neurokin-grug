package quizzes

import (
	"fmt"
	t "neurokin/types"
)

func GetQuiz(s string) (*t.Quiz, error) {
	if s == "aspie" {
		return &AspieQuiz, nil
	}

	return nil, fmt.Errorf("Quiz not found")
}
