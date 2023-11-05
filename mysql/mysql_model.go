package mysql

import "time"

type MysqlConfig struct {
	IsEnabled              bool          `json:"enabled" yaml:"enabled"`
	DebugMode              bool          `json:"debug_mode" yaml:"debug_mode"`
	Database               string        `json:"database" binding:"required" yaml:"database"`
	Host                   string        `json:"host" binding:"required" yaml:"host"`
	Port                   int           `json:"port" binding:"required" yaml:"port"`
	Username               string        `json:"username" yaml:"username"`
	Password               string        `json:"-" yaml:"password"`
	MaxOpenConn            int           `json:"max_open_conn" binding:"required" yaml:"max-open-conn"`
	MaxIdleConn            int           `json:"max_idle_conn" binding:"required" yaml:"max-idle-conn"`
	MaxLifeTimeMinutesConn int           `json:"max_life_time_minutes_conn" binding:"required" yaml:"max-life-time-minutes-conn"`
	Timeout                time.Duration `json:"-" yaml:"-"`
}

type mysqlOptionConfig struct {
}

type MultiTenantMysqlConfig struct {
	Key             string            `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool              `json:"usable_default" yaml:"usable_default"`
	Config          MysqlConfig       `json:"config" yaml:"config"`
	Option          mysqlOptionConfig `json:"option,omitempty" yaml:"option"`
}

type ClusterMultiTenantMysqlConfig struct {
	Clusters []MultiTenantMysqlConfig `json:"clusters,omitempty" yaml:"clusters"`
}
