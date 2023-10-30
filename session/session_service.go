package session

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/sivaosorg/govm/common"
	"github.com/sivaosorg/govm/cookies"
	"github.com/sivaosorg/govm/entity"
	"github.com/sivaosorg/govm/utils"
)

type SessionManagerService interface {
	GenSessionId() string
	CreateSession(w http.ResponseWriter, data map[string]interface{}) *session
	GetSessionId(request *http.Request) string
	GetSession(request *http.Request) *session

	// Session Middlewares
	RequireSessionMiddleware(next http.Handler) http.Handler
}

type sessionManagerServiceImpl struct {
	conf SessionManager
}

func NewSessionManagerService(conf SessionManager) SessionManagerService {
	s := &sessionManagerServiceImpl{
		conf: conf,
	}
	return s
}

func (s *sessionManagerServiceImpl) GenSessionId() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (s *sessionManagerServiceImpl) CreateSession(w http.ResponseWriter, data map[string]interface{}) *session {
	sessionId := s.GenSessionId()
	expiresAt := time.Now().Add(s.conf.Expiration)
	session := NewSession().
		SetId(sessionId).
		SetCreatedAt(time.Now()).
		SetExpiresAt(expiresAt).
		SetData(data)
	s.conf.SessionStore[sessionId] = session

	// updating the cookie
	// set cookie value
	// set cookie timeout
	s.conf.Cookie.SetValue(sessionId)
	s.conf.Cookie.SetTimeout(s.conf.Expiration)
	svc := cookies.NewCookieService()

	// set cookie based on session
	svc.SetCookie(w, s.conf.Cookie)
	return session
}

func (s *sessionManagerServiceImpl) GetSessionId(request *http.Request) string {
	cookieValue := request.Cookies()[0].Value
	return cookieValue
}

// GetSession retrieves a session based on the session ID from a request.
func (s *sessionManagerServiceImpl) GetSession(request *http.Request) *session {
	cookieValue := s.GetSessionId(request)
	if utils.IsEmpty(cookieValue) {
		return nil
	}
	session, ok := s.conf.SessionStore[cookieValue]
	if !ok || session.ExpiresAt.Before(time.Now()) {
		delete(s.conf.SessionStore, cookieValue)
		return nil
	}
	return session
}

// VerifySessionId checks if a session is valid based on session ID.
func (s *sessionManagerServiceImpl) VerifySessionId(sessionId string) bool {
	session, ok := s.conf.SessionStore[sessionId]
	if !ok || session.ExpiresAt.Before(time.Now()) {
		delete(s.conf.SessionStore, sessionId)
		return false
	}
	return true
}

// Middleware function to validate the session for protected routes.
func (s *sessionManagerServiceImpl) RequireSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionId := s.GetSessionId(r)
		if utils.IsNotEmpty(sessionId) && s.VerifySessionId(sessionId) {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set(common.HeaderContentType, common.MediaTypeApplicationJSON)
			w.WriteHeader(http.StatusUnauthorized)
			e := entity.NewResponseEntity().Unauthorized("Session Unauthorized", nil)
			w.Write([]byte(e.Json()))
		}
	})
}
