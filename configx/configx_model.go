package configx

import (
	"github.com/sivaosorg/govm/asterisk"
	"github.com/sivaosorg/govm/bot/slack"
	"github.com/sivaosorg/govm/bot/telegram"
	"github.com/sivaosorg/govm/corsx"
	"github.com/sivaosorg/govm/mongodb"
	"github.com/sivaosorg/govm/mysql"
	"github.com/sivaosorg/govm/postgres"
	"github.com/sivaosorg/govm/rabbitmqx"
	"github.com/sivaosorg/govm/redisx"
)

type FieldCommentConfig map[string]string

type TypeConfig string

type CommentedConfig struct {
	Data     interface{}        `json:"data" binding:"required" yaml:"-"`
	Comments FieldCommentConfig `json:"comments" yaml:"-"`
}

type KeysConfig struct {
	Asterisk asterisk.AsteriskConfig  `json:"asterisk,omitempty" yaml:"asterisk"`
	Mongodb  mongodb.MongodbConfig    `json:"mongodb,omitempty" yaml:"mongodb"`
	MySql    mysql.MysqlConfig        `json:"mysql,omitempty" yaml:"mysql"`
	Postgres postgres.PostgresConfig  `json:"postgres,omitempty" yaml:"postgres"`
	RabbitMq rabbitmqx.RabbitMqConfig `json:"rabbitmq,omitempty" yaml:"rabbitmq"`
	Redis    redisx.RedisConfig       `json:"redis,omitempty" yaml:"redis"`
	Telegram telegram.TelegramConfig  `json:"telegram,omitempty" yaml:"telegram"`
	Slack    slack.SlackConfig        `json:"slack,omitempty" yaml:"slack"`
	Cors     corsx.CorsConfig         `json:"cors,omitempty" yaml:"cors"`
}

type MultiTenancyKeysConfig struct {
	Key             string     `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool       `json:"usable_default" yaml:"usable_default"`
	Config          KeysConfig `json:"config" yaml:"config"`
}

type ClusterMultiTenancyKeysConfig struct {
	Clusters []MultiTenancyKeysConfig `json:"clusters,omitempty" yaml:"clusters"`
}
