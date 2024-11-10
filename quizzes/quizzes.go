package quizzes

import (
	"encoding/json"
	"fmt"
	t "neurokin/types"
	"os"
	"path/filepath"
	"strconv"
)

func GetQuiz(s string) (*t.Quiz, error) {
	if s == "aspie" {
		return &AspieQuiz, nil
	}

	quizzesPath, err := filepath.Abs("merged_quizzes.json")
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for quizzes.json: %v", err)
	}

	quizzesData, err := os.ReadFile(quizzesPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read quizzes.json file: %v", err)
	}

	// Load raw JSON data into an interface{} for type casting
	var rawData []map[string]interface{}
	if err := json.Unmarshal(quizzesData, &rawData); err != nil {
		return nil, fmt.Errorf("error parsing quizzes file: %v", err)
	}

	// Convert raw data to typed structs
	var quizzes []t.Quiz
	for _, quizData := range rawData {
		quiz := t.Quiz{
			Id:   int(quizData["id"].(float64)),
			Name: quizData["name"].(string),
			Slug: quizData["slug"].(string),
		}

		// Handle questions with type casting for id field
		questionsData, ok := quizData["questions"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid questions format for quiz '%s'", quiz.Slug)
		}

		for _, q := range questionsData {
			questionMap, ok := q.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid question format in quiz '%s'", quiz.Slug)
			}

			// Convert question id to string
			idValue := questionMap["id"]
			var idStr string
			switch v := idValue.(type) {
			case float64:
				idStr = strconv.Itoa(int(v))
			case string:
				idStr = v
			default:
				return nil, fmt.Errorf("unexpected type for question id in quiz '%s'", quiz.Slug)
			}

			question := t.QuizQuestion{
				Id:       idStr,
				Question: questionMap["question"].(string),
				Answers:  toStringSlice(questionMap["answers"]),
				Type:     t.QuestionType(questionMap["type"].(string)),
			}

			quiz.Questions = append(quiz.Questions, question)
		}

		quizzes = append(quizzes, quiz)
	}

	// Search for the quiz with the matching slug
	for _, quiz := range quizzes {
		if quiz.Slug == s {
			return &quiz, nil
		}
	}

	return nil, fmt.Errorf("quiz with slug '%s' not found", s)
}

// Helper function to convert interface{} to []string
func toStringSlice(data interface{}) []string {
	var result []string
	if items, ok := data.([]interface{}); ok {
		for _, item := range items {
			if str, ok := item.(string); ok {
				result = append(result, str)
			}
		}
	}
	return result
}

