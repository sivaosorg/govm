package redisx

type RedisConfig struct {
	IsEnabled bool   `json:"enabled" yaml:"enabled"`
	DebugMode bool   `json:"debug_mode" yaml:"debug-mode"`
	UrlConn   string `json:"url_conn" binding:"required" yaml:"url-conn"`
	Password  string `json:"password" binding:"required" yaml:"password"`
	Database  string `json:"database" binding:"required" yaml:"database"`
}
