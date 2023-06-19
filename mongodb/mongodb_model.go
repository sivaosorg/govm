package mongodb

type MongodbConfig struct {
	IsEnabled          bool   `json:"enabled" yaml:"enabled"`
	DebugMode          bool   `json:"debug_mode" yaml:"debug-mode"`
	UrlConn            string `json:"url_conn" yaml:"url-conn"`
	Host               string `json:"host" binding:"required" yaml:"host"`
	Port               int    `json:"port" binding:"required" yaml:"port"`
	Database           string `json:"database" binding:"required" yaml:"database"`
	Username           string `json:"username" yaml:"username"`
	Password           string `json:"password" yaml:"password"`
	TimeoutSecondsConn int    `json:"timeout_seconds_conn" yaml:"timeout-seconds-conn"`
	AllowConnSync      bool   `json:"allow_conn_sync" yaml:"allow-conn-sync"`
}
