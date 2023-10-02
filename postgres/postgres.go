package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/sivaosorg/govm/utils"
)

func NewPostgresConfig() *PostgresConfig {
	p := &PostgresConfig{}
	p.SetTimeout(10 * time.Second)
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

func (p *PostgresConfig) SetTimeout(value time.Duration) *PostgresConfig {
	p.Timeout = value
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

func NewPostgresOptionConfig() *postgresOptionConfig {
	p := &postgresOptionConfig{}
	return p
}

func NewMultiTenantPostgresConfig() *MultiTenantPostgresConfig {
	m := &MultiTenantPostgresConfig{}
	return m
}

func (m *MultiTenantPostgresConfig) SetKey(value string) *MultiTenantPostgresConfig {
	if utils.IsEmpty(value) {
		log.Panicf("Key is required")
	}
	m.Key = value
	return m
}

func (m *MultiTenantPostgresConfig) SetUsableDefault(value bool) *MultiTenantPostgresConfig {
	m.IsUsableDefault = value
	return m
}

func (m *MultiTenantPostgresConfig) SetConfig(value PostgresConfig) *MultiTenantPostgresConfig {
	m.Config = value
	return m
}

func (m *MultiTenantPostgresConfig) SetConfigCursor(value *PostgresConfig) *MultiTenantPostgresConfig {
	m.Config = *value
	return m
}

func (m *MultiTenantPostgresConfig) SetOption(value postgresOptionConfig) *MultiTenantPostgresConfig {
	m.Option = value
	return m
}

func (m *MultiTenantPostgresConfig) Json() string {
	return utils.ToJson(m)
}

func MultiTenantPostgresConfigValidator(m *MultiTenantPostgresConfig) {
	m.SetKey(m.Key)
}

func GetMultiTenantPostgresConfigSample() *MultiTenantPostgresConfig {
	m := NewMultiTenantPostgresConfig().
		SetKey("tenant_1").
		SetUsableDefault(false).
		SetOption(*NewPostgresOptionConfig()).
		SetConfigCursor(GetPostgresConfigSample())
	return m
}

func NewClusterMultiTenantPostgresConfig() *ClusterMultiTenantPostgresConfig {
	c := &ClusterMultiTenantPostgresConfig{}
	return c
}

func (c *ClusterMultiTenantPostgresConfig) SetClusters(values []MultiTenantPostgresConfig) *ClusterMultiTenantPostgresConfig {
	c.Clusters = values
	return c
}

func (c *ClusterMultiTenantPostgresConfig) AppendClusters(values ...MultiTenantPostgresConfig) *ClusterMultiTenantPostgresConfig {
	c.Clusters = append(c.Clusters, values...)
	return c
}

func (c *ClusterMultiTenantPostgresConfig) Json() string {
	return utils.ToJson(c.Clusters)
}

func GetClusterMultiTenantPostgresConfigSample() *ClusterMultiTenantPostgresConfig {
	c := NewClusterMultiTenantPostgresConfig().
		AppendClusters(*GetMultiTenantPostgresConfigSample(), *GetMultiTenantPostgresConfigSample().SetKey("tenant_2"))
	return c
}

func (c *ClusterMultiTenantPostgresConfig) FindClusterBy(key string) (MultiTenantPostgresConfig, error) {
	if utils.IsEmpty(key) {
		return *NewMultiTenantPostgresConfig(), fmt.Errorf("Key is required")
	}
	if len(c.Clusters) == 0 {
		return *NewMultiTenantPostgresConfig(), fmt.Errorf("No postgres cluster")
	}
	for _, v := range c.Clusters {
		if v.Key == key {
			return v, nil
		}
	}
	return *NewMultiTenantPostgresConfig(), fmt.Errorf("The postgres cluster not found")
}
