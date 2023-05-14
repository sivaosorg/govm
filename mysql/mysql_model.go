package mysql

type MysqlConfig struct {
	IsEnabled              bool   `json:"enabled" yaml:"enabled"`
	Database               string `json:"database" yaml:"database"`
	Host                   string `json:"host" yaml:"host"`
	Port                   int    `json:"port" yaml:"port"`
	Username               string `json:"username" yaml:"username"`
	Password               string `json:"password" yaml:"password"`
	MaxOpenConn            int    `json:"max_open_conn" yaml:"max-open-conn"`
	MaxIdleConn            int    `json:"max_idle_conn" yaml:"max-idle-conn"`
	MaxLifeTimeMinutesConn int    `json:"max_life_time_minutes_conn" yaml:"max-life-time-minutes-conn"`
}
