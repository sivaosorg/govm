package mongodb

import (
	"log"

	"github.com/sivaosorg/govm/utils"
)

func NewMongodbConfig() *MongodbConfig {
	m := &MongodbConfig{}
	return m
}

func (m *MongodbConfig) SetEnabled(value bool) *MongodbConfig {
	m.IsEnabled = value
	return m
}

func (m *MongodbConfig) SetUrlConn(value string) *MongodbConfig {
	m.UrlConn = utils.TrimSpaces(value)
	return m
}

func (m *MongodbConfig) SetHost(value string) *MongodbConfig {
	m.Host = utils.TrimSpaces(value)
	return m
}

func (m *MongodbConfig) SetPort(value int) *MongodbConfig {
	if value <= 0 {
		log.Panic("Invalid port")
	}
	m.Port = value
	return m
}

func (m *MongodbConfig) SetDatabase(value string) *MongodbConfig {
	m.Database = utils.TrimSpaces(value)
	return m
}

func (m *MongodbConfig) SetUsername(value string) *MongodbConfig {
	m.Username = value
	return m
}

func (m *MongodbConfig) SetPassword(value string) *MongodbConfig {
	m.Password = value
	return m
}

func (m *MongodbConfig) SetTimeoutSecondsConn(value int) *MongodbConfig {
	if value < 0 {
		log.Panic("Invalid timeout-seconds-conn")
	}
	m.TimeoutSecondsConn = value
	return m
}

func (m *MongodbConfig) SetAllowConnSync(value bool) *MongodbConfig {
	m.AllowConnSync = value
	return m
}

func (m *MongodbConfig) SetDebugMode(value bool) *MongodbConfig {
	m.DebugMode = value
	return m
}

func (m *MongodbConfig) Json() string {
	return utils.ToJson(m)
}

func MongodbConfigValidator(m *MongodbConfig) {
	m.SetTimeoutSecondsConn(m.TimeoutSecondsConn).
		SetPort(m.Port)
}

func GetMongodbConfigSample() *MongodbConfig {
	m := NewMongodbConfig()
	m.SetEnabled(true)
	m.SetHost("127.0.0.1")
	m.SetPort(27017)
	m.SetDatabase("u_db")
	m.SetUsername("u@root")
	m.SetPassword("pwd")
	m.SetAllowConnSync(true)
	m.SetTimeoutSecondsConn(30)
	m.SetUrlConn("mongodb://127.0.0.1:27017/u_db")
	return m
}
