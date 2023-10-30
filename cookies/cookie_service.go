package cookies

import (
	"fmt"
	"net/http"
	"time"
)

type CookieService interface {
	CreateCookie(conf CookieConfig) *http.Cookie
	SetCookie(w http.ResponseWriter, conf CookieConfig) error
	GetCookie(r *http.Request, conf CookieConfig) (string, error)
	DeleteCookie(w http.ResponseWriter, conf CookieConfig) error
}

type cookieServiceImpl struct {
}

func NewCookieService() CookieService {
	s := &cookieServiceImpl{}
	return s
}

func (s *cookieServiceImpl) CreateCookie(conf CookieConfig) *http.Cookie {
	return &http.Cookie{
		Name:       conf.Name,
		Value:      conf.Value,
		Path:       conf.Path,
		Domain:     conf.Domain,
		Expires:    time.Now().Add(conf.Timeout),
		RawExpires: "",
		MaxAge:     conf.MaxAge,
		Secure:     conf.Secure,
		HttpOnly:   conf.HttpOnly,
		SameSite:   0,
	}
}

func (s *cookieServiceImpl) SetCookie(w http.ResponseWriter, conf CookieConfig) error {
	if !conf.IsEnabled {
		return fmt.Errorf("Cookie unavailable")
	}
	c := s.CreateCookie(conf)
	http.SetCookie(w, c)
	return nil
}

func (s *cookieServiceImpl) GetCookie(r *http.Request, conf CookieConfig) (string, error) {
	if !conf.IsEnabled {
		return "", fmt.Errorf("Cookie unavailable")
	}
	c, err := r.Cookie(conf.Name)
	if err != nil {
		return "", err
	}
	return c.Value, nil
}

func (s *cookieServiceImpl) DeleteCookie(w http.ResponseWriter, conf CookieConfig) error {
	if !conf.IsEnabled {
		return fmt.Errorf("Cookie unavailable")
	}
	c := s.CreateCookie(conf)
	c.MaxAge = -1
	c.Value = ""
	http.SetCookie(w, c)
	return nil
}
