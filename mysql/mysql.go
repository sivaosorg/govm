package mysql

import (
	"fmt"
	"log"
	"time"

	"github.com/sivaosorg/govm/utils"
)

func NewMysqlConfig() *MysqlConfig {
	m := &MysqlConfig{}
	m.SetTimeout(10 * time.Second)
	return m
}

func (m *MysqlConfig) SetEnabled(value bool) *MysqlConfig {
	m.IsEnabled = value
	return m
}

func (m *MysqlConfig) SetDatabase(value string) *MysqlConfig {
	if utils.IsEmpty(value) {
		log.Panicf("Database is required")
	}
	m.Database = utils.TrimSpaces(value)
	return m
}

func (m *MysqlConfig) SetHost(value string) *MysqlConfig {
	if utils.IsEmpty(value) {
		log.Panicf("Host is required")
	}
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

func (m *MysqlConfig) SetDebugMode(value bool) *MysqlConfig {
	m.DebugMode = value
	return m
}

func (m *MysqlConfig) SetTimeout(value time.Duration) *MysqlConfig {
	m.Timeout = value
	return m
}

func (m *MysqlConfig) Json() string {
	return utils.ToJson(m)
}

func (m *MysqlConfig) GetConnString() string {
	MysqlConfigValidator(m)
	hostname := fmt.Sprintf("%s:%d", m.Host, m.Port)
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", m.Username, m.Password, hostname, m.Database)
}

func MysqlConfigValidator(m *MysqlConfig) {
	m.SetPort(m.Port).
		SetDatabase(m.Database).
		SetMaxOpenConn(m.MaxOpenConn).
		SetMaxIdleConn(m.MaxIdleConn).
		SetMaxLifeTimeMinutesConn(m.MaxLifeTimeMinutesConn).
		SetHost(m.Host)
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

func NewMysqlOptionConfig() *mysqlOptionConfig {
	m := &mysqlOptionConfig{}
	return m
}

func NewMultiTenantMysqlConfig() *MultiTenantMysqlConfig {
	m := &MultiTenantMysqlConfig{}
	return m
}

func (m *MultiTenantMysqlConfig) SetKey(value string) *MultiTenantMysqlConfig {
	if utils.IsEmpty(value) {
		log.Panicf("Key is required")
	}
	m.Key = value
	return m
}

func (m *MultiTenantMysqlConfig) SetUsableDefault(value bool) *MultiTenantMysqlConfig {
	m.IsUsableDefault = value
	return m
}

func (m *MultiTenantMysqlConfig) SetConfig(value MysqlConfig) *MultiTenantMysqlConfig {
	m.Config = value
	return m
}

func (m *MultiTenantMysqlConfig) SetConfigCursor(value *MysqlConfig) *MultiTenantMysqlConfig {
	m.Config = *value
	return m
}

func (m *MultiTenantMysqlConfig) SetOption(value mysqlOptionConfig) *MultiTenantMysqlConfig {
	m.Option = value
	return m
}

func (m *MultiTenantMysqlConfig) Json() string {
	return utils.ToJson(m)
}

func MultiTenantMysqlConfigValidator(m *MultiTenantMysqlConfig) {
	m.SetKey(m.Key)
}

func GetMultiTenantMysqlConfigSample() *MultiTenantMysqlConfig {
	m := NewMultiTenantMysqlConfig().
		SetKey("tenant_1").
		SetUsableDefault(false).
		SetConfigCursor(GetMysqlConfigSample())
	return m
}

func NewClusterMultiTenantMysqlConfig() *ClusterMultiTenantMysqlConfig {
	c := &ClusterMultiTenantMysqlConfig{}
	return c
}

func (c *ClusterMultiTenantMysqlConfig) SetClusters(values []MultiTenantMysqlConfig) *ClusterMultiTenantMysqlConfig {
	c.Clusters = values
	return c
}

func (c *ClusterMultiTenantMysqlConfig) AppendClusters(values ...MultiTenantMysqlConfig) *ClusterMultiTenantMysqlConfig {
	c.Clusters = append(c.Clusters, values...)
	return c
}

func (c *ClusterMultiTenantMysqlConfig) Json() string {
	return utils.ToJson(c.Clusters)
}

func GetClusterMultiTenantMysqlConfigSample() *ClusterMultiTenantMysqlConfig {
	c := NewClusterMultiTenantMysqlConfig().
		AppendClusters(*GetMultiTenantMysqlConfigSample(), *GetMultiTenantMysqlConfigSample().SetKey("tenant_2"))
	return c
}

func (c *ClusterMultiTenantMysqlConfig) FindClusterBy(key string) (MultiTenantMysqlConfig, error) {
	if utils.IsEmpty(key) {
		return *NewMultiTenantMysqlConfig(), fmt.Errorf("Key is required")
	}
	if len(c.Clusters) == 0 {
		return *NewMultiTenantMysqlConfig(), fmt.Errorf("No mysql cluster")
	}
	for _, v := range c.Clusters {
		if v.Key == key {
			return v, nil
		}
	}
	return *NewMultiTenantMysqlConfig(), fmt.Errorf("The mysql cluster not found")
}
