package rabbitmqx

import (
	"fmt"
	"log"

	"github.com/sivaosorg/govm/utils"
)

func NewRabbitMqConfig() *RabbitMqConfig {
	r := &RabbitMqConfig{}
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

func GetRabbitMqConfigSample() *RabbitMqConfig {
	r := NewRabbitMqConfig().
		SetEnabled(true).
		SetPort(5672).
		SetUsername("guest").
		SetPassword("guest").
		SetHost("127.0.0.1").
		SetUrlConn("amqp://guest:guest@localhost:5672/")
	return r
}
