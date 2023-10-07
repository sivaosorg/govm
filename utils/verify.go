package utils

import (
	"fmt"
	"regexp"
)

func VerifyEmail(email string) (bool, error) {
	regex := `^[a-zA-Z0-9._%+-]+@([a-zA-Z0-9.-]+\.[a-zA-Z]{2,}|localhost)$`
	return regexp.MatchString(regex, email)
}

func VerifyPassword(password string, min, max int) (bool, error) {
	if len(password) < min || len(password) > max {
		return false, fmt.Errorf("Invalid length password (%v), length minimum (%v) and maximum (%v) required", len(password), min, max)
	}
	regex := fmt.Sprintf("^[a-zA-Z]*[0-9]+[a-zA-Z]*[@#$%%^&+=!]+[a-zA-Z0-9]*$")
	return regexp.MatchString(regex, password)
}
