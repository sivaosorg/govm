package rabbitmqx

import (
	"fmt"
	"log"

	"github.com/sivaosorg/govm/coltx"
	"github.com/sivaosorg/govm/utils"
)

func NewRabbitMqConfig() *RabbitMqConfig {
	r := &RabbitMqConfig{}
	return r
}

func NewRabbitMqExchangeConfig() *RabbitMqExchangeConfig {
	r := &RabbitMqExchangeConfig{}
	return r
}

func NewRabbitMqQueueConfig() *RabbitMqQueueConfig {
	r := &RabbitMqQueueConfig{}
	return r
}

func NewRabbitMqMessageConfig() *RabbitMqMessageConfig {
	r := &RabbitMqMessageConfig{}
	return r
}

func (r *RabbitMqConfig) SetEnabled(value bool) *RabbitMqConfig {
	r.IsEnabled = value
	return r
}

func (r *RabbitMqConfig) SetDebugMode(value bool) *RabbitMqConfig {
	r.DebugMode = value
	return r
}

func (r *RabbitMqConfig) SetUrlConn(value string) *RabbitMqConfig {
	r.UrlConn = value
	return r
}

func (r *RabbitMqConfig) SetUsername(value string) *RabbitMqConfig {
	if utils.IsEmpty(value) {
		log.Panic("Username is required")
	}
	r.Username = utils.TrimSpaces(value)
	return r
}

func (r *RabbitMqConfig) SetPassword(value string) *RabbitMqConfig {
	if utils.IsEmpty(r.UrlConn) {
		if utils.IsEmpty(value) {
			log.Panic("Password is required")
		}
	}
	r.Password = value
	return r
}

func (r *RabbitMqConfig) SetPort(value int) *RabbitMqConfig {
	if value <= 0 {
		log.Panicf("Invalid port: %d", value)
	}
	r.Port = value
	return r
}

func (r *RabbitMqConfig) SetHost(value string) *RabbitMqConfig {
	if utils.IsEmpty(r.UrlConn) {
		if utils.IsEmpty(value) {
			log.Panic("Host is required")
		}
	}
	if utils.IsNotEmpty(value) {
		value = utils.RemovePrefix(value, "http://", "https://")
		r.Host = utils.TrimSpaces(value)
	}
	return r
}

func (r *RabbitMqConfig) SetMessage(value RabbitMqMessageConfig) *RabbitMqConfig {
	r.Message = value
	return r
}

func (r *RabbitMqConfig) SetClusters(values map[string]RabbitMqMessageConfig) *RabbitMqConfig {
	if len(values) > 0 {
		r.Clusters = values
	}
	return r
}

func (r *RabbitMqConfig) Json() string {
	return utils.ToJson(r)
}

func (r *RabbitMqConfig) ToUrlConn() string {
	RabbitMqConfigValidator(r)
	if utils.IsNotEmpty(r.UrlConn) {
		return r.UrlConn
	}
	conn := fmt.Sprintf("%s://%s:%s@%s:%d/", RabbitMqPrefix, r.Username, r.Password, r.Host, r.Port)
	r.SetUrlConn(conn)
	return conn
}

func RabbitMqConfigValidator(r *RabbitMqConfig) {
	r.SetUrlConn(r.UrlConn).
		SetHost(r.Host).
		SetPort(r.Port).
		SetUsername(r.Username).
		SetPassword(r.Password)
}

func RabbitMqExchangeConfigValidator(r *RabbitMqExchangeConfig) {
	r.SetName(r.Name).SetKind(r.Kind)
}

func RabbitMqQueueConfigValidator(r *RabbitMqQueueConfig) {
	r.SetName(r.Name)
}

func GetRabbitMqConfigSample() *RabbitMqConfig {
	r := NewRabbitMqConfig().
		SetEnabled(true).
		SetPort(5672).
		SetUsername("guest").
		SetPassword("guest").
		SetHost("127.0.0.1").
		SetUrlConn("amqp://guest:guest@localhost:5672/").
		SetMessage(*GetRabbitMqMessageConfigSample())
	return r
}

func GetRabbitMqExchangeConfigSample() *RabbitMqExchangeConfig {
	r := NewRabbitMqExchangeConfig().
		SetName("guest_exchange").
		SetKind(ExchangeFanout).
		SetDurable(true)
	return r
}

func GetRabbitMqQueueConfigSample() *RabbitMqQueueConfig {
	r := NewRabbitMqQueueConfig().
		SetName("guest_queue").
		SetDurable(true)
	return r
}

func GetRabbitMqMessageConfigSample() *RabbitMqMessageConfig {
	r := NewRabbitMqMessageConfig().
		SetEnabled(true).
		SetExchange(*GetRabbitMqExchangeConfigSample()).
		SetQueue(*GetRabbitMqQueueConfigSample())
	return r
}

func RabbitMqExchangesString() string {
	return coltx.JoinMapKeys(Exchanges, ",")
}

func (r *RabbitMqExchangeConfig) SetName(value string) *RabbitMqExchangeConfig {
	if utils.IsEmpty(value) {
		log.Panic("Name exchange is required")
	}
	r.Name = value
	return r
}

func (r *RabbitMqExchangeConfig) SetKind(value string) *RabbitMqExchangeConfig {
	_, ok := Exchanges[value]
	if !ok {
		log.Panicf("Invalid exchange kind, RabbitMQ only supported exchange kind: %v", RabbitMqExchangesString())
	}
	r.Kind = value
	return r
}

func (r *RabbitMqExchangeConfig) SetDurable(value bool) *RabbitMqExchangeConfig {
	r.Durable = value
	return r
}

func (r *RabbitMqQueueConfig) SetName(value string) *RabbitMqQueueConfig {
	if utils.IsEmpty(value) {
		log.Panic("Name queue is required")
	}
	r.Name = value
	return r
}

func (r *RabbitMqQueueConfig) SetDurable(value bool) *RabbitMqQueueConfig {
	r.Durable = value
	return r
}

func (r *RabbitMqMessageConfig) SetEnabled(value bool) *RabbitMqMessageConfig {
	r.IsEnabled = value
	return r
}

func (r *RabbitMqMessageConfig) SetExchange(value RabbitMqExchangeConfig) *RabbitMqMessageConfig {
	r.Exchange = value
	return r
}

func (r *RabbitMqMessageConfig) SetQueue(value RabbitMqQueueConfig) *RabbitMqMessageConfig {
	r.Queue = value
	return r
}
