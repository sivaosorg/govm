package session

import (
	"time"

	"github.com/sivaosorg/govm/cookies"
)

type session struct {
	Id        string                 `json:"_id" yaml:"id"`
	CreatedAt time.Time              `json:"created_at,omitempty" yaml:"-"`
	ExpiresAt time.Time              `json:"expires_at,omitempty" yaml:"-"`
	Data      map[string]interface{} `json:"data,omitempty" yaml:"data"`
}

type SessionManager struct {
	SecretKey    string               `json:"secret_key" yaml:"secret_key"`
	Expiration   time.Duration        `json:"expires_at" yaml:"expires_at"`
	Cookie       cookies.CookieConfig `json:"cookie" yaml:"cookie"`
	SessionStore map[string]*session  `json:"store" yaml:"store"`
}
