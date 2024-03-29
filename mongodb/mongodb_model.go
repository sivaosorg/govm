package mongodb

import "time"

type MongodbConfig struct {
	IsEnabled          bool          `json:"enabled" yaml:"enabled"`
	DebugMode          bool          `json:"debug_mode" yaml:"debug_mode"`
	UrlConn            string        `json:"url_conn" yaml:"url_conn"`
	Host               string        `json:"host" binding:"required" yaml:"host"`
	Port               int           `json:"port" binding:"required" yaml:"port"`
	Database           string        `json:"database" binding:"required" yaml:"database"`
	Username           string        `json:"username" yaml:"username"`
	Password           string        `json:"-" yaml:"password"`
	TimeoutSecondsConn int           `json:"timeout_second_conn" yaml:"timeout_second_conn"`
	AllowConnSync      bool          `json:"allow_conn_sync" yaml:"allow_conn_sync"`
	Timeout            time.Duration `json:"timeout" yaml:"timeout"` // for context
}

type mongodbOptionConfig struct {
}

type MultiTenantMongodbConfig struct {
	Key             string              `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool                `json:"usable_default" yaml:"usable_default"`
	Config          MongodbConfig       `json:"config" yaml:"config"`
	Option          mongodbOptionConfig `json:"option,omitempty" yaml:"option"`
}

type ClusterMultiTenantMongodbConfig struct {
	Clusters []MultiTenantMongodbConfig `json:"clusters,omitempty" yaml:"clusters"`
}
