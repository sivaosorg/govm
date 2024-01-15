package rabbitmqx

import "time"

type clusters map[string]RabbitMqMessageConfig

type RabbitMqConfig struct {
	IsEnabled bool                  `json:"enabled" yaml:"enabled"`
	DebugMode bool                  `json:"debug_mode" yaml:"debug_mode"`
	UrlConn   string                `json:"url_conn" binding:"required" yaml:"url_conn"`
	Username  string                `json:"username" binding:"required" yaml:"username"`
	Password  string                `json:"-" binding:"required" yaml:"password"`
	Host      string                `json:"host" yaml:"host"`
	Port      int                   `json:"port" binding:"required" yaml:"port"`
	Message   RabbitMqMessageConfig `json:"message,omitempty" yaml:"message"`
	Clusters  clusters              `json:"clusters,omitempty" yaml:"clusters"`
	Timeout   time.Duration         `json:"timeout" yaml:"timeout"`
}

type RabbitMqExchangeConfig struct {
	Name    string `json:"name" yaml:"name"`
	Kind    string `json:"kind" yaml:"kind"`
	Durable bool   `json:"durable" yaml:"durable"`
}

type RabbitMqQueueConfig struct {
	Name    string `json:"name" yaml:"name"`
	Durable bool   `json:"durable" yaml:"durable"`
}

type RabbitMqMessageConfig struct {
	IsEnabled bool                   `json:"enabled" yaml:"enabled"`
	Exchange  RabbitMqExchangeConfig `json:"exchange,omitempty" yaml:"exchange"`
	Queue     RabbitMqQueueConfig    `json:"queue,omitempty" yaml:"queue"`
}

type rabbitMqOptionConfig struct {
}

type MultiTenantRabbitMqConfig struct {
	Key             string               `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool                 `json:"usable_default" yaml:"usable_default"`
	Config          RabbitMqConfig       `json:"config" yaml:"config"`
	Option          rabbitMqOptionConfig `json:"option,omitempty" yaml:"option"`
}

type ClusterMultiTenantRabbitMqConfig struct {
	Clusters []MultiTenantRabbitMqConfig `json:"clusters,omitempty" yaml:"clusters"`
}
