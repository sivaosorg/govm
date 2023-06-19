package postgres

import (
	"log"

	"github.com/sivaosorg/govm/utils"
)

func NewPostgresConfig() *PostgresConfig {
	p := &PostgresConfig{}
	return p
}

func (p *PostgresConfig) SetEnabled(value bool) *PostgresConfig {
	p.IsEnabled = value
	return p
}

func (p *PostgresConfig) SetDatabase(value string) *PostgresConfig {
	p.Database = utils.TrimSpaces(value)
	return p
}

func (p *PostgresConfig) SetHost(value string) *PostgresConfig {
	p.Host = utils.TrimSpaces(value)
	return p
}

func (p *PostgresConfig) SetPort(value int) *PostgresConfig {
	if value <= 0 {
		log.Panic("Invalid port")
	}
	p.Port = value
	return p
}

func (p *PostgresConfig) SetUsername(value string) *PostgresConfig {
	p.Username = value
	return p
}

func (p *PostgresConfig) SetPassword(value string) *PostgresConfig {
	p.Password = value
	return p
}

func (p *PostgresConfig) SetSslMode(value string) *PostgresConfig {
	if _, ok := PostgresSslModes[value]; !ok {
		log.Panic(PostgresSslModeError)
	}
	p.SSLMode = value
	return p
}

func (p *PostgresConfig) SetMaxOpenConn(value int) *PostgresConfig {
	if value <= 0 {
		log.Panicf("Invalid max-open-conn")
	}
	p.MaxOpenConn = value
	return p
}

func (p *PostgresConfig) SetMaxIdleConn(value int) *PostgresConfig {
	if value <= 0 {
		log.Panicf("Invalid max-idle-conn")
	}
	p.MaxIdleConn = value
	return p
}

func (p *PostgresConfig) SetDebugMode(value bool) *PostgresConfig {
	p.DebugMode = value
	return p
}

func (p *PostgresConfig) Json() string {
	return utils.ToJson(p)
}

func PostgresConfigValidator(p *PostgresConfig) {
	p.SetPort(p.Port).
		SetMaxOpenConn(p.MaxOpenConn).
		SetMaxIdleConn(p.MaxIdleConn)
}

func GetPostgresConfigSample() *PostgresConfig {
	p := NewPostgresConfig()
	p.SetEnabled(true)
	p.SetDatabase("u_db")
	p.SetHost("127.0.0.1")
	p.SetPort(5432)
	p.SetUsername("u@root")
	p.SetPassword("pwd")
	p.SetSslMode("disable")
	p.SetMaxOpenConn(5)
	p.SetMaxIdleConn(3)
	return p
}
