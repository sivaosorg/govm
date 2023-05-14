package postgres

type PostgresConfig struct {
	IsEnabled   bool   `json:"enabled" yaml:"enabled"`
	Database    string `json:"database" binding:"required" yaml:"database"`
	Host        string `json:"host" binding:"required" yaml:"host"`
	Port        int    `json:"port" binding:"required" yaml:"port"`
	Username    string `json:"username" yaml:"username"`
	Password    string `json:"password" yaml:"password"`
	SSLMode     string `json:"ssl_mode" binding:"required" yaml:"ssl-mode"`
	MaxOpenConn int    `json:"max_open_conn" binding:"required" yaml:"max-open-conn"`
	MaxIdleConn int    `json:"max_idle_conn" binding:"required" yaml:"max-idle-conn"`
}
