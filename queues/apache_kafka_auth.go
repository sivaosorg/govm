package queues

import (
	"strings"

	"github.com/sivaosorg/govm/utils"
)

func NewKafkaAuthConfig() *KafkaAuthConfig {
	return &KafkaAuthConfig{}
}

func (k *KafkaAuthConfig) SetEnabled(value bool) *KafkaAuthConfig {
	k.IsEnabled = value
	return k
}

func (k *KafkaAuthConfig) SetBootstrapServers(values []string) *KafkaAuthConfig {
	k.BootstrapServers = values
	return k
}

func (k *KafkaAuthConfig) AppendBootstrapServers(values ...string) *KafkaAuthConfig {
	k.BootstrapServers = append(k.BootstrapServers, values...)
	return k
}

func (k *KafkaAuthConfig) BootstrapServersString() string {
	return strings.Join(k.BootstrapServers, ",")
}

func (k *KafkaAuthConfig) SetSecurityProtocol(value string) *KafkaAuthConfig {
	k.SecurityProtocol = value
	return k
}

func (k *KafkaAuthConfig) SetSaslMechanism(value string) *KafkaAuthConfig {
	k.SaslMechanism = value
	return k
}

func (k *KafkaAuthConfig) SetSaslUsername(value string) *KafkaAuthConfig {
	k.SaslUsername = value
	return k
}

func (k *KafkaAuthConfig) SetSaslPassword(value string) *KafkaAuthConfig {
	k.SaslPassword = value
	return k
}

func (k *KafkaAuthConfig) SetSslCaLocation(value string) *KafkaAuthConfig {
	k.SslCaLocation = value
	return k
}

func (k *KafkaAuthConfig) SetSslCertificateLocation(value string) *KafkaAuthConfig {
	k.SslCertificateLocation = value
	return k
}

func (k *KafkaAuthConfig) SetSslKeyLocation(value string) *KafkaAuthConfig {
	k.SslKeyLocation = value
	return k
}

func (k *KafkaAuthConfig) Json() string {
	return utils.ToJson(k)
}

func GetKafkaAuthConfigSample() *KafkaAuthConfig {
	k := NewKafkaAuthConfig()
	k.IsEnabled = false
	k.AppendBootstrapServers("kafka-broker-1:9092", "kafka-broker-2:9092")
	k.SecurityProtocol = "sasl_ssl"
	k.SaslMechanism = "PLAIN"
	k.SaslUsername = "username"
	k.SaslPassword = "pwd"
	k.SslCaLocation = "/path/to/ca-certificate.pem"
	k.SslCertificateLocation = "/path/to/client-certificate.pem"
	k.SslKeyLocation = "/path/to/client-key.pem"
	return k
}
