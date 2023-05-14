package mongodb

type MongodbConfig struct {
	IsEnabled          bool   `json:"enabled" yaml:"enabled"`
	UrlConn            string `json:"url_conn" yaml:"url-conn"`
	Host               string `json:"host" yaml:"host"`
	Port               int    `json:"port" yaml:"port"`
	Database           string `json:"database" yaml:"database"`
	Username           string `json:"username" yaml:"username"`
	Password           string `json:"password" yaml:"password"`
	TimeoutSecondsConn int    `json:"timeout_seconds_conn" yaml:"timeout-seconds-conn"`
	AllowConnSync      bool   `json:"allow_conn_sync" yaml:"allow-conn-sync"`
}
