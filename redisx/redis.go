package redisx

import (
	"log"

	"github.com/sivaosorg/govm/utils"
)

func NewRedisConfig() *RedisConfig {
	r := &RedisConfig{}
	return r
}

func (r *RedisConfig) SetEnabled(value bool) *RedisConfig {
	r.IsEnabled = value
	return r
}

func (r *RedisConfig) SetUrlConn(value string) *RedisConfig {
	r.UrlConn = utils.TrimSpaces(value)
	return r
}

func (r *RedisConfig) SetPassword(value string) *RedisConfig {
	r.Password = value
	return r
}

func (r *RedisConfig) SetDatabase(value string) *RedisConfig {
	r.Database = utils.TrimSpaces(value)
	return r
}

func (r *RedisConfig) SetDebugMode(value bool) *RedisConfig {
	r.DebugMode = value
	return r
}

func (r *RedisConfig) Json() string {
	return utils.ToJson(r)
}

func GetRedisConfigSample() *RedisConfig {
	r := NewRedisConfig()
	r.SetEnabled(true)
	r.SetPassword("redis.pwd")
	r.SetDatabase("database_stable")
	r.SetUrlConn("localhost:6379")
	return r
}

func NewRedisOptionConfig() *redisOptionConfig {
	r := &redisOptionConfig{}
	return r
}

func NewMultiTenantRedisConfig() *MultiTenantRedisConfig {
	m := &MultiTenantRedisConfig{}
	return m
}

func (m *MultiTenantRedisConfig) SetKey(value string) *MultiTenantRedisConfig {
	if utils.IsEmpty(value) {
		log.Panicf("Key is required")
	}
	m.Key = value
	return m
}

func (m *MultiTenantRedisConfig) SetUsableDefault(value bool) *MultiTenantRedisConfig {
	m.IsUsableDefault = value
	return m
}

func (m *MultiTenantRedisConfig) SetConfig(value RedisConfig) *MultiTenantRedisConfig {
	m.Config = value
	return m
}

func (m *MultiTenantRedisConfig) SetConfigCursor(value *RedisConfig) *MultiTenantRedisConfig {
	m.Config = *value
	return m
}

func (m *MultiTenantRedisConfig) SetOption(value redisOptionConfig) *MultiTenantRedisConfig {
	m.Option = value
	return m
}

func (m *MultiTenantRedisConfig) Json() string {
	return utils.ToJson(m)
}

func MultiTenantRedisConfigValidator(m *MultiTenantRedisConfig) {
	m.SetKey(m.Key)
}

func GetMultiTenantRedisConfigSample() *MultiTenantRedisConfig {
	m := NewMultiTenantRedisConfig().
		SetKey("tenant_1").
		SetUsableDefault(false).
		SetConfigCursor(GetRedisConfigSample()).
		SetOption(*NewRedisOptionConfig())
	return m
}

func NewClusterMultiTenantRedisConfig() *ClusterMultiTenantRedisConfig {
	c := &ClusterMultiTenantRedisConfig{}
	return c
}

func (c *ClusterMultiTenantRedisConfig) SetClusters(values []MultiTenantRedisConfig) *ClusterMultiTenantRedisConfig {
	c.Clusters = values
	return c
}

func (c *ClusterMultiTenantRedisConfig) AppendClusters(values ...MultiTenantRedisConfig) *ClusterMultiTenantRedisConfig {
	c.Clusters = append(c.Clusters, values...)
	return c
}

func (c *ClusterMultiTenantRedisConfig) Json() string {
	return utils.ToJson(c.Clusters)
}

func GetClusterMultiTenantRedisConfigSample() *ClusterMultiTenantRedisConfig {
	c := NewClusterMultiTenantRedisConfig().
		AppendClusters(*GetMultiTenantRedisConfigSample(), *GetMultiTenantRedisConfigSample().SetKey("tenant_2"))
	return c
}
