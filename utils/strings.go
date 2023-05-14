package utils

import (
	"strings"
	"unicode"
)

func TrimSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func IsSpace(str string) bool {
	for _, c := range str {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

func IsEmptyAbsolute(str string) bool {
	return len(str) == 0 ||
		str == "" ||
		strings.TrimSpace(str) == ""
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func RemovePrefix(str string, prefix ...string) string {
	if IsEmptyAbsolute(str) {
		return str
	}
	if len(prefix) == 0 {
		return str
	}
	for _, v := range prefix {
		str = strings.TrimPrefix(str, v)
	}
	return str
}
