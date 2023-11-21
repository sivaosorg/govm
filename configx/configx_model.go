package configx

import (
	"github.com/sivaosorg/govm/asterisk"
	"github.com/sivaosorg/govm/bot/slack"
	"github.com/sivaosorg/govm/bot/telegram"
	"github.com/sivaosorg/govm/cookies"
	"github.com/sivaosorg/govm/corsx"
	"github.com/sivaosorg/govm/logger"
	"github.com/sivaosorg/govm/mongodb"
	"github.com/sivaosorg/govm/mysql"
	"github.com/sivaosorg/govm/postgres"
	"github.com/sivaosorg/govm/rabbitmqx"
	"github.com/sivaosorg/govm/redisx"
	"github.com/sivaosorg/govm/server"
)

type FieldCommentConfig map[string]string

type TypeConfig string

type CommentedConfig struct {
	Data     interface{}        `json:"data" binding:"required" yaml:"-"`
	Comments FieldCommentConfig `json:"comments" yaml:"-"`
}

type KeysConfig struct {
	// Basic
	Asterisk asterisk.AsteriskConfig  `json:"asterisk,omitempty" yaml:"asterisk"`
	Mongodb  mongodb.MongodbConfig    `json:"mongodb,omitempty" yaml:"mongodb"`
	MySql    mysql.MysqlConfig        `json:"mysql,omitempty" yaml:"mysql"`
	Postgres postgres.PostgresConfig  `json:"postgres,omitempty" yaml:"postgres"`
	RabbitMq rabbitmqx.RabbitMqConfig `json:"rabbitmq,omitempty" yaml:"rabbitmq"`
	Redis    redisx.RedisConfig       `json:"redis,omitempty" yaml:"redis"`
	Telegram telegram.TelegramConfig  `json:"telegram,omitempty" yaml:"telegram"`
	Slack    slack.SlackConfig        `json:"slack,omitempty" yaml:"slack"`
	Cors     corsx.CorsConfig         `json:"cors,omitempty" yaml:"cors"`
	Server   server.Server            `json:"server,omitempty" yaml:"server"`
	Cookie   cookies.CookieConfig     `json:"cookie,omitempty" yaml:"cookie"`
	Logger   logger.Logger            `json:"logger,omitempty" yaml:"logger"`

	// Seekers
	TelegramSeekers []telegram.MultiTenantTelegramConfig  `json:"telegram_seekers,omitempty" yaml:"telegram-seekers"`
	SlackSeekers    []slack.MultiTenantSlackConfig        `json:"slack_seekers,omitempty" yaml:"slack-seekers"`
	AsteriskSeekers []asterisk.MultiTenantAsteriskConfig  `json:"asterisk_seekers,omitempty" yaml:"asterisk-seekers"`
	MongodbSeekers  []mongodb.MultiTenantMongodbConfig    `json:"mongodb_seekers,omitempty" yaml:"mongodb-seekers"`
	MySqlSeekers    []mysql.MultiTenantMysqlConfig        `json:"mysql_seekers,omitempty" yaml:"mysql-seekers"`
	PostgresSeekers []postgres.MultiTenantPostgresConfig  `json:"postgres_seekers,omitempty" yaml:"postgres-seekers"`
	RabbitMqSeekers []rabbitmqx.MultiTenantRabbitMqConfig `json:"rabbitmq_seekers,omitempty" yaml:"rabbitmq-seekers"`
	RedisSeekers    []redisx.MultiTenantRedisConfig       `json:"redis_seekers,omitempty" yaml:"redis-seekers"`
	CookieSeekers   []cookies.MultiTenantCookieConfig     `json:"cookie_seekers,omitempty" yaml:"cookie-seekers"`
	LoggerSeekers   []logger.MultiTenantLoggerConfig      `json:"logger_seekers,omitempty" yaml:"logger-seekers"`

	// Params unknown
	Param1 map[string]interface{} `json:"param1,omitempty" yaml:"param1"`
	Param2 map[string]interface{} `json:"param2,omitempty" yaml:"param2"`
	Param3 map[string]interface{} `json:"param3,omitempty" yaml:"param3"`
	Param4 map[string]interface{} `json:"param4,omitempty" yaml:"param4"`
	KV     interface{}            `json:"kv,omitempty" yaml:"kv"`
}

type MultiTenancyKeysConfig struct {
	Key             string     `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool       `json:"usable_default" yaml:"usable_default"`
	Config          KeysConfig `json:"config" yaml:"config"`
}

type ClusterMultiTenancyKeysConfig struct {
	Clusters []MultiTenancyKeysConfig `json:"clusters,omitempty" yaml:"clusters"`
}
