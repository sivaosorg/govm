package queues

import (
	"github.com/sivaosorg/govm/builder"
	"github.com/sivaosorg/govm/utils"
)

func NewKafka() *Kafka {
	return &Kafka{}
}

func (k *Kafka) SetEnabled(value bool) *Kafka {
	k.IsEnabled = value
	return k
}

func (k *Kafka) SetTopics(values []KafkaTopicConfig) *Kafka {
	k.Topics = values
	return k
}

func (k *Kafka) AppendTopics(values ...KafkaTopicConfig) *Kafka {
	k.Topics = append(k.Topics, values...)
	return k
}

func (k *Kafka) SetAppliedAuth(value KafkaAuthConfig) *Kafka {
	k.AppliedAuth = value
	return k
}

func (k *Kafka) SetProducer(value KafkaProducerConfig) *Kafka {
	k.Producer = value
	return k
}

func (k *Kafka) SetConsumer(value KafkaConsumerConfig) *Kafka {
	k.Consumer = value
	return k
}

func (k *Kafka) Json() string {
	return utils.ToJson(k)
}

func (k *Kafka) AvailableTopics() bool {
	return len(k.Topics) > 0
}

func (k *Kafka) GetAvailableTopics() ([]KafkaTopicConfig, bool) {
	return GetAvailableTopics(k.Topics)
}

func (k *Kafka) GetAvailableTopicsString() ([]string, bool) {
	return GetAvailableTopicsString(k.Topics)
}

func (k *Kafka) GetTopicsAutoCreated() ([]KafkaTopicConfig, bool) {
	return GetTopicsAutoCreated(k.Topics)
}

func (k *Kafka) GetTopic(key string) (KafkaTopicConfig, bool) {
	return GetTopic(k.Topics, key)
}

func (k *Kafka) RemoveProducerPropsAuthKeys() *Kafka {
	k.Producer.SetProperties(k.removePropsAuthKeys(k.Producer.Props))
	return k
}

func (k *Kafka) RemoveConsumerPropsAuthKeys() *Kafka {
	k.Consumer.SetProperties(k.removePropsAuthKeys(k.Consumer.Props))
	return k
}

func (k *Kafka) GetAuthBasedProducer() KafkaAuthConfig {
	if k.Producer.AppliedAuth.IsEnabled {
		return k.Producer.AppliedAuth
	}
	return k.AppliedAuth
}

func (k *Kafka) GetAuthMapBasedProducer() map[string]interface{} {
	auth := k.GetAuthBasedProducer()
	builder := builder.NewMapBuilder()
	builder.Add(Bootstrap_Servers, auth.BootstrapServersString())
	builder.Add(Security_Protocol, auth.SecurityProtocol)
	builder.Add(Sasl_Mechanism, auth.SaslMechanism)
	builder.Add(Sasl_Username, auth.SaslUsername)
	builder.Add(Sasl_Password, auth.SaslPassword)
	if utils.IsNotEmpty(auth.SslCaLocation) {
		builder.Add(Ssl_Ca_Location, auth.SslCaLocation)
	}
	if utils.IsNotEmpty(auth.SslCertificateLocation) {
		builder.Add(Ssl_Certificate_Location, auth.SslCertificateLocation)
	}
	if utils.IsNotEmpty(auth.SslKeyLocation) {
		builder.Add(Ssl_Key_Location, auth.SslKeyLocation)
	}
	return builder.Build()
}

func (k *Kafka) GetAuthBasedConsumer() KafkaAuthConfig {
	if k.Consumer.AppliedAuth.IsEnabled {
		return k.Consumer.AppliedAuth
	}
	return k.AppliedAuth
}

func (k *Kafka) GetAuthMapBasedConsumer() map[string]interface{} {
	auth := k.GetAuthBasedConsumer()
	builder := builder.NewMapBuilder()
	builder.Add(Bootstrap_Servers, auth.BootstrapServersString())
	builder.Add(Security_Protocol, auth.SecurityProtocol)
	builder.Add(Sasl_Mechanism, auth.SaslMechanism)
	builder.Add(Sasl_Username, auth.SaslUsername)
	builder.Add(Sasl_Password, auth.SaslPassword)
	if utils.IsNotEmpty(auth.SslCaLocation) {
		builder.Add(Ssl_Ca_Location, auth.SslCaLocation)
	}
	if utils.IsNotEmpty(auth.SslCertificateLocation) {
		builder.Add(Ssl_Certificate_Location, auth.SslCertificateLocation)
	}
	if utils.IsNotEmpty(auth.SslKeyLocation) {
		builder.Add(Ssl_Key_Location, auth.SslKeyLocation)
	}
	return builder.Build()
}

func (k *Kafka) AvailableAuthBasedProducer() bool {
	return k.GetAuthBasedProducer().IsEnabled
}

func (k *Kafka) AvailableAuthBasedConsumer() bool {
	return k.GetAuthBasedConsumer().IsEnabled
}

func (k *Kafka) GetPropsBasedProducer() map[string]interface{} {
	k.RemoveProducerPropsAuthKeys()
	if !k.AvailableAuthBasedProducer() {
		return k.Producer.Props
	}
	_map := k.Producer.Props
	for key, value := range k.GetAuthMapBasedProducer() {
		_map[key] = value
	}
	return _map
}

func (k *Kafka) GetPropsBasedProducerJson() string {
	return utils.ToJson(k.GetPropsBasedProducer())
}

func (k *Kafka) GetPropsBasedConsumer() map[string]interface{} {
	k.RemoveConsumerPropsAuthKeys()
	if !k.AvailableAuthBasedConsumer() {
		return k.Consumer.Props
	}
	_map := k.Consumer.Props
	for key, value := range k.GetAuthMapBasedConsumer() {
		_map[key] = value
	}
	return _map
}

func (k *Kafka) GetPropsBasedConsumerJson() string {
	return utils.ToJson(k.GetPropsBasedConsumer())
}

func (k *Kafka) removePropsAuthKeys(props map[string]interface{}) map[string]interface{} {
	if len(props) == 0 {
		return props
	}
	delete(props, Bootstrap_Servers)
	delete(props, Security_Protocol)
	delete(props, Sasl_Mechanism)
	delete(props, Sasl_Username)
	delete(props, Sasl_Password)
	delete(props, Ssl_Ca_Location)
	delete(props, Ssl_Certificate_Location)
	delete(props, Ssl_Key_Location)
	return props
}

func GetKafkaSample() *Kafka {
	k := NewKafka().SetEnabled(false).
		SetAppliedAuth(*GetKafkaAuthConfigSample()).
		SetProducer(*GetKafkaProducerConfigSample()).
		SetConsumer(*GetKafkaConsumerConfigSample()).
		AppendTopics(*GetKafkaTopicConfigSample(), *GetKafkaTopicConfigSample().SetKey("topic-2"))
	return k
}
