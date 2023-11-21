package utils

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func TrimSpaces(s string) string {
	if IsEmpty(s) {
		return s
	}
	return strings.Join(strings.Fields(s), " ")
}

func TrimAllSpaces(s string) string {
	if IsEmpty(s) {
		return s
	}
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

// GenUUID returns a new UUID based on /dev/urandom (unix).
func GenUUID() (string, error) {
	file, err := os.Open("/dev/urandom")
	if err != nil {
		return "", fmt.Errorf("open /dev/urandom error:[%v]", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}()
	b := make([]byte, 16)

	_, err = file.Read(b)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}

func GenUUIDShorten() string {
	uuid, err := GenUUID()
	if err != nil {
		return ""
	}
	return uuid
}

func RepeatPlaceholders(format string, value interface{}) (string, error) {
	placeholders := strings.Count(format, "%")
	values := make([]interface{}, placeholders)
	for i := range values {
		values[i] = value
	}
	result := fmt.Sprintf(format, values...)
	return result, nil
}
