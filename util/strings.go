package util

import (
	"regexp"
	"strings"
	"unicode"
)

func ValidateEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}

	if len(s) == 1 {
		return strings.ToUpper(s)
	}

	f := []rune(s)[0]
	if unicode.IsLetter(f) {
		return strings.ToUpper(string(f)) + s[1:]
	}

	return s
}
