package postgres

import "time"

type PostgresConfig struct {
	IsEnabled   bool          `json:"enabled" yaml:"enabled"`
	DebugMode   bool          `json:"debug_mode" yaml:"debug_mode"`
	Database    string        `json:"database" binding:"required" yaml:"database"`
	Host        string        `json:"host" binding:"required" yaml:"host"`
	Port        int           `json:"port" binding:"required" yaml:"port"`
	Username    string        `json:"username" yaml:"username"`
	Password    string        `json:"-" yaml:"password"`
	SSLMode     string        `json:"ssl_mode" binding:"required" yaml:"ssl-mode"`
	MaxOpenConn int           `json:"max_open_conn" binding:"required" yaml:"max-open-conn"`
	MaxIdleConn int           `json:"max_idle_conn" binding:"required" yaml:"max-idle-conn"`
	Timeout     time.Duration `json:"timeout" yaml:"timeout"`
}

type postgresOptionConfig struct {
}

type MultiTenantPostgresConfig struct {
	Key             string               `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool                 `json:"usable_default" yaml:"usable_default"`
	Config          PostgresConfig       `json:"config" yaml:"config"`
	Option          postgresOptionConfig `json:"option,omitempty" yaml:"option"`
}

type ClusterMultiTenantPostgresConfig struct {
	Clusters []MultiTenantPostgresConfig `json:"clusters,omitempty" yaml:"clusters"`
}
