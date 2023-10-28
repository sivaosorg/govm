package server

import (
	"log"
	"time"

	"github.com/sivaosorg/govm/utils"
)

func NewServer() *Server {
	s := &Server{}
	return s
}

func (s *Server) SetHost(value string) *Server {
	s.Host = value
	return s
}

func (s *Server) SetPort(value int) *Server {
	if value <= 0 {
		log.Panicf("Invalid port: %v", value)
	}
	s.Port = value
	return s
}

func (s *Server) SetTimeout(value Timeout) *Server {
	s.Timeout = value
	return s
}

func (s *Server) SetMode(value string) *Server {
	s.Mode = value
	return s
}

func (s *Server) Json() string {
	return utils.ToJson(s)
}

func ServerValidator(s *Server) {
	s.SetPort(s.Port)
}

func GetServerSample() *Server {
	s := NewServer().
		SetHost("127.0.0.1").
		SetPort(8083).
		SetTimeout(*GetTimeoutSample())
	return s
}

func NewTimeout() *Timeout {
	t := &Timeout{}
	return t
}

func (t *Timeout) SetServe(value time.Duration) *Timeout {
	t.Serve = value
	return t
}

func (t *Timeout) SetRead(value time.Duration) *Timeout {
	t.Read = value
	return t
}

func (t *Timeout) SetWrite(value time.Duration) *Timeout {
	t.Write = value
	return t
}

func (t *Timeout) SetIdle(value time.Duration) *Timeout {
	t.Idle = value
	return t
}

func GetTimeoutSample() *Timeout {
	t := NewTimeout().
		SetIdle(10 * time.Second).
		SetRead(10 * time.Second).
		SetWrite(10 * time.Second).
		SetServe(10 * time.Second)
	return t
}
