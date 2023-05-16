package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func Base64Encode(v interface{}) string {
	d := ToJson(v)
	return base64.StdEncoding.EncodeToString([]byte(d))
}

func Base64EncodeString(v string) string {
	return base64.StdEncoding.EncodeToString([]byte(v))
}

func Base64EncodeByte(v []byte) string {
	return base64.StdEncoding.EncodeToString(v)
}

func Base64Decode(encoded string) string {
	if IsEmptyAbsolute(encoded) {
		return encoded
	}
	d, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return ""
	}
	return string(d)
}

// Generate sign key base on HMAC256 algorithm, after that base64 encode
// We are going to sign a message with our secret key to create a HMAC (Keyed-Hash Message Authentication Code) hash and give it to client.
// The secret is shared between both client and server applications.
// If you give a different message or hash as opposed to original ones, the communication will be treated as tampered.
func Base64EncodeHmac256(message, secret []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(message)
	return Base64EncodeByte(h.Sum(nil))
}

// Generate sign key base on HMAC256 algorithm, after that base64 encode
func Base64EncodeHmac256With(message, secret string) string {
	return Base64EncodeHmac256([]byte(message), []byte(secret))
}
