package quizzes

import (
	"encoding/json"
	"fmt"
	t "neurokin/types"
	"os"
	"path/filepath"
)

func GetQuiz(s string) (*t.Quiz, error) {
	if s == "aspie" {
		return &AspieQuiz, nil
	}

	quizzesPath, err := filepath.Abs("quizzes.json")
	if err != nil {
		return nil, fmt.Errorf("Quiz file not loaded")
	}

	var quizzes []t.Quiz
	quizzesData, err := os.ReadFile(quizzesPath)
	if err := json.Unmarshal(quizzesData, &quizzes); err != nil {
		return nil, fmt.Errorf("Error parsing quizzes file")
	}

	for _, quiz := range quizzes {
		if quiz.Slug == s {
			return &quiz, nil
		}
	}

	return nil, fmt.Errorf("Quiz not found")
}
