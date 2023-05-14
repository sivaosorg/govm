package mysql

type MysqlConfig struct {
	IsEnabled              bool   `json:"enabled" yaml:"enabled"`
	Database               string `json:"database" binding:"required" yaml:"database"`
	Host                   string `json:"host" binding:"required" yaml:"host"`
	Port                   int    `json:"port" binding:"required" yaml:"port"`
	Username               string `json:"username" yaml:"username"`
	Password               string `json:"password" yaml:"password"`
	MaxOpenConn            int    `json:"max_open_conn" binding:"required" yaml:"max-open-conn"`
	MaxIdleConn            int    `json:"max_idle_conn" binding:"required" yaml:"max-idle-conn"`
	MaxLifeTimeMinutesConn int    `json:"max_life_time_minutes_conn" binding:"required" yaml:"max-life-time-minutes-conn"`
}
