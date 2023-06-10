package utils

import (
	"strings"
	"unicode"
)

func TrimSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func TrimAllSpaces(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	return ReplaceAllSpecialCharacters(s)
}

func IsSpace(s string) bool {
	for _, c := range s {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

func IsEmpty(s string) bool {
	return len(s) == 0 ||
		s == "" ||
		strings.TrimSpace(s) == "" ||
		len(strings.TrimSpace(s)) == 0
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func RemovePrefix(s string, prefix ...string) string {
	if IsEmpty(s) {
		return s
	}
	if len(prefix) == 0 {
		return s
	}
	for _, v := range prefix {
		s = strings.TrimPrefix(s, v)
	}
	return s
}

func ReplaceAllSpecialCharacters(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\t", "")
	return s
}
