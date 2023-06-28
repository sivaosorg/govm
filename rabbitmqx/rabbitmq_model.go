package rabbitmqx

type RabbitMqConfig struct {
	IsEnabled bool   `json:"enabled" yaml:"enabled"`
	DebugMode bool   `json:"debug_mode" yaml:"debug-mode"`
	UrlConn   string `json:"url_conn" binding:"required" yaml:"url-conn"`
	Username  string `json:"username" binding:"required" yaml:"username"`
	Password  string `json:"password" binding:"required" yaml:"password"`
	Host      string `json:"host" yaml:"host"`
	Port      int    `json:"port" binding:"required" yaml:"port"`
}
