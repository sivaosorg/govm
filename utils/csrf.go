package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

// Create CSRF token
func CreateCSRFToken(sid, salt string) (string, error) {
	hash := sha256.New()
	_, err := io.WriteString(hash, fmt.Sprintf("%s%s", salt, sid))
	if err != nil {
		return "", err
	}
	token := base64.RawStdEncoding.EncodeToString(hash.Sum(nil))
	return token, nil
}

// Validate CSRF token
func VerifyCSRFToken(token string, sid, salt string) bool {
	t, _ := CreateCSRFToken(sid, salt)
	return token == t
}
