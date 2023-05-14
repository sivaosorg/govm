package mysql

import (
	"log"

	"github.com/sivaosorg/govm/utils"
)

func NewMysqlConfig() *MysqlConfig {
	m := &MysqlConfig{}
	return m
}

func (m *MysqlConfig) SetEnabled(value bool) *MysqlConfig {
	m.IsEnabled = value
	return m
}

func (m *MysqlConfig) SetDatabase(value string) *MysqlConfig {
	m.Database = utils.TrimSpaces(value)
	return m
}

func (m *MysqlConfig) SetHost(value string) *MysqlConfig {
	m.Host = utils.TrimSpaces(value)
	return m
}

func (m *MysqlConfig) SetPort(value int) *MysqlConfig {
	if value <= 0 {
		log.Panic("Invalid port")
	}
	m.Port = value
	return m
}

func (m *MysqlConfig) SetUsername(value string) *MysqlConfig {
	m.Username = value
	return m
}

func (m *MysqlConfig) SetPassword(value string) *MysqlConfig {
	m.Password = value
	return m
}

func (m *MysqlConfig) SetMaxOpenConn(value int) *MysqlConfig {
	if value <= 0 {
		log.Panic("Invalid max-open-conn")
	}
	m.MaxOpenConn = value
	return m
}

func (m *MysqlConfig) SetMaxIdleConn(value int) *MysqlConfig {
	if value <= 0 {
		log.Panic("Invalid max-idle-conn")
	}
	m.MaxIdleConn = value
	return m
}

func (m *MysqlConfig) SetMaxLifeTimeMinutesConn(values int) *MysqlConfig {
	if values < 0 {
		log.Panic("Invalid max-life-time-minutes-conn")
	}
	m.MaxLifeTimeMinutesConn = values
	return m
}

func (m *MysqlConfig) Json() string {
	return utils.ToJson(m)
}

func MysqlConfigValidator(m *MysqlConfig) {
	m.SetPort(m.Port).
		SetMaxOpenConn(m.MaxOpenConn).
		SetMaxIdleConn(m.MaxIdleConn).
		SetMaxLifeTimeMinutesConn(m.MaxLifeTimeMinutesConn)
}

func GetMysqlConfigSample() *MysqlConfig {
	m := NewMysqlConfig()
	m.SetEnabled(true)
	m.SetDatabase("u_db")
	m.SetHost("127.0.0.1")
	m.SetPort(3306)
	m.SetUsername("u@root")
	m.SetPassword("pwd")
	m.SetMaxIdleConn(2)
	m.SetMaxOpenConn(10)
	m.SetMaxLifeTimeMinutesConn(10)
	return m
}
