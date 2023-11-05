package server

import "time"

type Server struct {
	Host    string  `json:"host" yaml:"host"`
	Port    int     `json:"port" binding:"required" yaml:"port"`
	Mode    string  `json:"mode" yaml:"mode"`
	Timeout Timeout `json:"timeout" yaml:"timeout"`
	Attr    Attr    `json:"attr" yaml:"attr"`
	SSL     SSL     `json:"ssl" yaml:"ssl"`
	SP      Pprof   `json:"serve_proxy" yaml:"serve_proxy"`
}

type Timeout struct {
	Serve time.Duration `json:"serve" yaml:"serve"`
	Read  time.Duration `json:"read" yaml:"read"`
	Write time.Duration `json:"write" yaml:"write"`
	Idle  time.Duration `json:"idle" yaml:"idle"`
}

type Attr struct {
	MaxHeaderBytes int `json:"max_header_bytes" yaml:"max_header_bytes"`
}

type SSL struct {
	IsEnabled bool   `json:"enabled" yaml:"enabled"`
	CertFile  string `json:"cert_file" yaml:"cert_file"`
	KeyFile   string `json:"key_file" yaml:"key_file"`
}

type Pprof struct {
	IsEnabled bool    `json:"enabled" yaml:"enabled"`
	Port      int     `json:"port" binding:"required" yaml:"port"`
	Timeout   Timeout `json:"timeout" yaml:"timeout"`
	Attr      Attr    `json:"attr" yaml:"attr"`
}
