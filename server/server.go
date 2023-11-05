package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sivaosorg/govm/logger"
	"github.com/sivaosorg/govm/utils"
)

func NewServer() *Server {
	s := &Server{}
	s.SetAttr(*NewAttribute())
	s.SetTimeout(*NewTimeout().SetRead(15 * time.Second).SetWrite(15 * time.Second))
	s.SetSSL(*GetSSLSample())
	s.SetMode("debug")
	s.SetSP(*GetPprofSample())
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

func (s *Server) SetAttr(value Attr) *Server {
	s.Attr = value
	return s
}

func (s *Server) SetSSL(value SSL) *Server {
	s.SSL = value
	return s
}

func (s *Server) SetSP(value Pprof) *Server {
	s.SP = value
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

func NewAttribute() *Attr {
	return &Attr{
		MaxHeaderBytes: 1 << 20,
	}
}

func (a *Attr) SetMaxHeaderBytes(value int) *Attr {
	if value < 0 {
		log.Panicf("Invalid max_header_bytes: %v", value)
	}
	a.MaxHeaderBytes = value
	return a
}

func (s *Server) CreateAppServer(handler http.Handler) *http.Server {
	h := &http.Server{
		Addr:           fmt.Sprintf(":%v", s.Port),
		ReadTimeout:    s.Timeout.Read,
		WriteTimeout:   s.Timeout.Write,
		MaxHeaderBytes: s.Attr.MaxHeaderBytes,
		Handler:        handler,
	}
	return h
}

func NewSsl() *SSL {
	return &SSL{}
}

func (s *SSL) SetEnabled(value bool) *SSL {
	s.IsEnabled = value
	return s
}

func (s *SSL) SetCertFile(value string) *SSL {
	s.CertFile = value
	return s
}

func (s *SSL) SetKeyFile(value string) *SSL {
	s.KeyFile = value
	return s
}

func (s *SSL) Json() string {
	return utils.ToJson(s)
}

func GetSSLSample() *SSL {
	s := NewSsl().
		SetEnabled(false).
		SetCertFile("./keys/ssl/cert.crt").
		SetKeyFile("./keys/ssl/key.pem")
	return s
}

func NewPprof() *Pprof {
	return &Pprof{}
}

func (p *Pprof) SetEnabled(value bool) *Pprof {
	p.IsEnabled = value
	return p
}

func (p *Pprof) SetPort(value int) *Pprof {
	p.Port = value
	return p
}

func (p *Pprof) SetTimeout(value Timeout) *Pprof {
	p.Timeout = value
	return p
}

func (p *Pprof) SetAttr(value Attr) *Pprof {
	p.Attr = value
	return p
}

func (p *Pprof) Json() string {
	return utils.ToJson(p)
}

func (p *Pprof) CreateAppServer(handler http.Handler) *http.Server {
	h := &http.Server{
		Addr:           fmt.Sprintf(":%v", p.Port),
		ReadTimeout:    p.Timeout.Read,
		WriteTimeout:   p.Timeout.Write,
		MaxHeaderBytes: p.Attr.MaxHeaderBytes,
		Handler:        handler,
	}
	return h
}

func GetPprofSample() *Pprof {
	p := NewPprof().
		SetEnabled(false).
		SetPort(8101).
		SetAttr(*NewAttribute()).
		SetTimeout(*GetTimeoutSample())
	return p
}

func StartServer(h *http.Server) {
	err := h.ListenAndServe()
	if err != nil {
		logger.Errorf("Start server got an error: %v", err, h.Addr)
		panic(err)
	}
}

func StartServerSecure(h *http.Server, cert, key string) {
	err := h.ListenAndServeTLS(cert, key)
	if err != nil {
		logger.Errorf("Start server TLS got an error: %v", err, h.Addr)
		panic(err)
	}
}
