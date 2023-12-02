package queues

import (
	"fmt"

	"github.com/sivaosorg/govm/builder"
	"github.com/sivaosorg/govm/utils"
)

func NewKafka() *KafkaConfig {
	return &KafkaConfig{}
}

func (k *KafkaConfig) SetEnabled(value bool) *KafkaConfig {
	k.IsEnabled = value
	return k
}

func (k *KafkaConfig) SetTopics(values []KafkaTopicConfig) *KafkaConfig {
	k.Topics = values
	return k
}

func (k *KafkaConfig) AppendTopics(values ...KafkaTopicConfig) *KafkaConfig {
	k.Topics = append(k.Topics, values...)
	return k
}

func (k *KafkaConfig) SetAppliedAuth(value KafkaAuthConfig) *KafkaConfig {
	k.AppliedAuth = value
	return k
}

func (k *KafkaConfig) SetProducer(value KafkaProducerConfig) *KafkaConfig {
	k.Producer = value
	return k
}

func (k *KafkaConfig) SetConsumer(value KafkaConsumerConfig) *KafkaConfig {
	k.Consumer = value
	return k
}

func (k *KafkaConfig) Json() string {
	return utils.ToJson(k)
}

func (k *KafkaConfig) AvailableTopics() bool {
	return len(k.Topics) > 0
}

func (k *KafkaConfig) GetAvailableTopics() ([]KafkaTopicConfig, bool) {
	return GetAvailableTopics(k.Topics)
}

func (k *KafkaConfig) GetAvailableTopicsString() ([]string, bool) {
	return GetAvailableTopicsString(k.Topics)
}

func (k *KafkaConfig) GetTopicsAutoCreated() ([]KafkaTopicConfig, bool) {
	return GetTopicsAutoCreated(k.Topics)
}

func (k *KafkaConfig) GetTopic(key string) (KafkaTopicConfig, bool) {
	return GetTopic(k.Topics, key)
}

func (k *KafkaConfig) RemoveProducerPropsAuthKeys() *KafkaConfig {
	k.Producer.SetProperties(k.removePropsAuthKeys(k.Producer.Props))
	return k
}

func (k *KafkaConfig) RemoveConsumerPropsAuthKeys() *KafkaConfig {
	k.Consumer.SetProperties(k.removePropsAuthKeys(k.Consumer.Props))
	return k
}

func (k *KafkaConfig) GetAuthBasedProducer() KafkaAuthConfig {
	if k.Producer.AppliedAuth.IsEnabled {
		return k.Producer.AppliedAuth
	}
	return k.AppliedAuth
}

func (k *KafkaConfig) GetAuthMapBasedProducer() map[string]interface{} {
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

func (k *KafkaConfig) GetAuthBasedConsumer() KafkaAuthConfig {
	if k.Consumer.AppliedAuth.IsEnabled {
		return k.Consumer.AppliedAuth
	}
	return k.AppliedAuth
}

func (k *KafkaConfig) GetAuthMapBasedConsumer() map[string]interface{} {
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

func (k *KafkaConfig) AvailableAuthBasedProducer() bool {
	return k.GetAuthBasedProducer().IsEnabled
}

func (k *KafkaConfig) AvailableAuthBasedConsumer() bool {
	return k.GetAuthBasedConsumer().IsEnabled
}

func (k *KafkaConfig) GetPropsBasedProducer() map[string]interface{} {
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

func (k *KafkaConfig) GetPropsBasedProducerJson() string {
	return utils.ToJson(k.GetPropsBasedProducer())
}

func (k *KafkaConfig) GetPropsBasedConsumer() map[string]interface{} {
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

func (k *KafkaConfig) GetPropsBasedConsumerJson() string {
	return utils.ToJson(k.GetPropsBasedConsumer())
}

func (k *KafkaConfig) removePropsAuthKeys(props map[string]interface{}) map[string]interface{} {
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

func GetKafkaSample() *KafkaConfig {
	k := NewKafka().SetEnabled(false).
		SetAppliedAuth(*GetKafkaAuthConfigSample()).
		SetProducer(*GetKafkaProducerConfigSample()).
		SetConsumer(*GetKafkaConsumerConfigSample()).
		AppendTopics(*GetKafkaTopicConfigSample(), *GetKafkaTopicConfigSample().SetKey("topic-2"))
	return k
}

func NewMultiTenantKafkaConfig() *MultiTenantKafkaConfig {
	return &MultiTenantKafkaConfig{}
}

func (k *MultiTenantKafkaConfig) SetKey(value string) *MultiTenantKafkaConfig {
	k.Key = value
	return k
}

func (k *MultiTenantKafkaConfig) SetUsableDefault(value bool) *MultiTenantKafkaConfig {
	k.IsUsableDefault = value
	return k
}

func (k *MultiTenantKafkaConfig) SetConfig(value KafkaConfig) *MultiTenantKafkaConfig {
	k.Config = value
	return k
}

func (k *MultiTenantKafkaConfig) SetConfigCursor(value *KafkaConfig) *MultiTenantKafkaConfig {
	k.Config = *value
	return k
}

func (k *MultiTenantKafkaConfig) Json() string {
	return utils.ToJson(k)
}

func GetMultiTenantKafkaConfigSample() *MultiTenantKafkaConfig {
	k := NewMultiTenantKafkaConfig().
		SetKey("tenant_1").
		SetConfig(*GetKafkaSample()).
		SetUsableDefault(false)
	return k
}

func NewClusterMultiTenantKafkaConfig() *ClusterMultiTenantKafkaConfig {
	return &ClusterMultiTenantKafkaConfig{
		Clusters: make([]MultiTenantKafkaConfig, 0),
	}
}

func (c *ClusterMultiTenantKafkaConfig) SetClusters(values []MultiTenantKafkaConfig) *ClusterMultiTenantKafkaConfig {
	c.Clusters = values
	return c
}

func (c *ClusterMultiTenantKafkaConfig) AppendClusters(values ...MultiTenantKafkaConfig) *ClusterMultiTenantKafkaConfig {
	c.Clusters = append(c.Clusters, values...)
	return c
}

func (c *ClusterMultiTenantKafkaConfig) Json() string {
	return utils.ToJson(c.Clusters)
}

func (c *ClusterMultiTenantKafkaConfig) FindClusterBy(key string) (MultiTenantKafkaConfig, error) {
	if utils.IsEmpty(key) {
		return *NewMultiTenantKafkaConfig(), fmt.Errorf("Key is required")
	}
	if len(c.Clusters) == 0 {
		return *NewMultiTenantKafkaConfig(), fmt.Errorf("No kafka cluster")
	}
	for _, v := range c.Clusters {
		if v.Key == key {
			return v, nil
		}
	}
	return *NewMultiTenantKafkaConfig(), fmt.Errorf("The kafka cluster not found")
}

func GetClusterMultiTenantKafkaConfigSample() *ClusterMultiTenantKafkaConfig {
	c := NewClusterMultiTenantKafkaConfig().
		AppendClusters(*GetMultiTenantKafkaConfigSample(), *GetMultiTenantKafkaConfigSample().SetKey("tenant_2"))
	return c
}
