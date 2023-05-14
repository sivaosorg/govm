package redisx

type RedisConfig struct {
	IsEnabled bool   `json:"enabled" yaml:"enabled"`
	UrlConn   string `json:"url_conn" binding:"required" yaml:"url-conn"`
	Password  string `json:"password" binding:"required" yaml:"password"`
	Database  string `json:"database" binding:"required" yaml:"database"`
}
