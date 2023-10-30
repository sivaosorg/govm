package session

import (
	"time"

	"github.com/sivaosorg/govm/cookies"
	"github.com/sivaosorg/govm/timex"
	"github.com/sivaosorg/govm/utils"
)

func NewSession() *session {
	return &session{
		Data: make(map[string]interface{}),
	}
}

func (s *session) SetId(value string) *session {
	s.Id = value
	return s
}

func (s *session) SetCreatedAt(value time.Time) *session {
	s.CreatedAt = value
	return s
}

func (s *session) SetExpiresAt(value time.Time) *session {
	s.ExpiresAt = value
	return s
}

func (s *session) SetData(value map[string]interface{}) *session {
	s.Data = value
	return s
}

func (s *session) AppendData(key string, value interface{}) *session {
	s.Data[key] = value
	return s
}

func (s *session) Json() string {
	return utils.ToJson(s)
}

func (s *session) Available() bool {
	return s != nil
}

func GetSessionSample() *session {
	s := NewSession().
		SetCreatedAt(time.Now()).
		SetExpiresAt(timex.AddDays(time.Now(), 1)).
		AppendData("session_data_key", 123).
		SetId(utils.GenUUIDShorten())
	return s
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		SessionStore: make(map[string]*session),
	}
}

func (s *SessionManager) SetSecretKey(value string) *SessionManager {
	s.SecretKey = value
	return s
}

func (s *SessionManager) SetExpiration(value time.Duration) *SessionManager {
	s.Expiration = value
	return s
}

func (s *SessionManager) SetCookie(value cookies.CookieConfig) *SessionManager {
	s.Cookie = value
	return s
}

func (s *SessionManager) SetStore(value map[string]*session) *SessionManager {
	s.SessionStore = value
	return s
}

func (s *SessionManager) AppendStore(key string, value *session) *SessionManager {
	s.SessionStore[key] = value
	return s
}

func (s *SessionManager) Json() string {
	return utils.ToJson(s)
}

func GetSessionManagerSample() *SessionManager {
	s := NewSessionManager().
		SetSecretKey(utils.GenUUIDShorten()).
		SetExpiration(10*time.Minute).
		SetCookie(*cookies.GetCookieConfigSample()).
		AppendStore(utils.GenUUIDShorten(), GetSessionSample())
	return s
}
