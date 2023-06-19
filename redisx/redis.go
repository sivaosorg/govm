package redisx

import "github.com/sivaosorg/govm/utils"

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
