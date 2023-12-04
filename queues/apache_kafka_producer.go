package queues

import (
	"fmt"

	"github.com/sivaosorg/govm/utils"
)

func NewKafkaProducerConfig() *KafkaProducerConfig {
	return &KafkaProducerConfig{
		Props: make(map[string]interface{}),
	}
}

func (k *KafkaProducerConfig) SetEnabled(value bool) *KafkaProducerConfig {
	k.IsEnabled = value
	return k
}

func (k *KafkaProducerConfig) SetAppliedAuth(value KafkaAuthConfig) *KafkaProducerConfig {
	k.AppliedAuth = value
	return k
}

func (k *KafkaProducerConfig) SetProperties(value map[string]interface{}) *KafkaProducerConfig {
	k.Props = value
	return k
}

func (k *KafkaProducerConfig) AppendProperty(key string, value interface{}) *KafkaProducerConfig {
	if k.Props == nil {
		k.Props = make(map[string]interface{})
	}
	k.Props[key] = value
	return k
}

func (k *KafkaProducerConfig) LenProperties() int {
	return len(k.Props)
}

func (k *KafkaProducerConfig) AvailableProperties() bool {
	return k.LenProperties() > 0
}

func (k *KafkaProducerConfig) Json() string {
	return utils.ToJson(k)
}

func GetKafkaProducerConfigSample() *KafkaProducerConfig {
	k := NewKafkaProducerConfig().
		SetEnabled(true).
		SetAppliedAuth(*GetKafkaAuthConfigSample()).
		AppendProperty("client.id", "producer-client-id").
		AppendProperty("acks", "all").
		AppendProperty("retries", 3).AppendProperty(Bootstrap_Servers, "kafka-broker-3:9092")
	return k
}

func NewKafkaPublisherRequest() *KafkaPublisherRequest {
	return &KafkaPublisherRequest{}
}

func (k *KafkaPublisherRequest) SetTopicKey(value string) *KafkaPublisherRequest {
	k.TopicKey = value
	return k
}

func (k *KafkaPublisherRequest) SetTenantKey(value string) *KafkaPublisherRequest {
	k.TenantKey = value
	return k
}

func (k *KafkaPublisherRequest) SetPayload(value map[string]interface{}) *KafkaPublisherRequest {
	k.Payload = value
	return k
}

func (k *KafkaPublisherRequest) Json() string {
	return utils.ToJson(k)
}

func KafkaPublisherRequestValidator(k KafkaPublisherRequest) error {
	if utils.IsEmpty(k.TopicKey) {
		return fmt.Errorf("Topic key is required")
	}
	if utils.IsEmpty(k.TenantKey) {
		return fmt.Errorf("Tenant key is required")
	}
	return nil
}
