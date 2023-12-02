package queues

import "github.com/sivaosorg/govm/utils"

func NewKafkaConsumerConfig() *KafkaConsumerConfig {
	return &KafkaConsumerConfig{
		Props: make(map[string]interface{}),
	}
}

func (k *KafkaConsumerConfig) SetEnabled(value bool) *KafkaConsumerConfig {
	k.IsEnabled = value
	return k
}

func (k *KafkaConsumerConfig) SetAppliedAuth(value KafkaAuthConfig) *KafkaConsumerConfig {
	k.AppliedAuth = value
	return k
}

func (k *KafkaConsumerConfig) SetProperties(value map[string]interface{}) *KafkaConsumerConfig {
	k.Props = value
	return k
}

func (k *KafkaConsumerConfig) AppendProperty(key string, value interface{}) *KafkaConsumerConfig {
	if k.Props == nil {
		k.Props = make(map[string]interface{})
	}
	k.Props[key] = value
	return k
}

func (k *KafkaConsumerConfig) LenProperties() int {
	return len(k.Props)
}

func (k *KafkaConsumerConfig) AvailableProperties() bool {
	return k.LenProperties() > 0
}

func (k *KafkaConsumerConfig) GroupId() (string, bool) {
	if !k.AvailableProperties() {
		return "", false
	}
	v, ok := k.Props[GroupId]
	return v.(string), ok
}

func (k *KafkaConsumerConfig) Json() string {
	return utils.ToJson(k)
}

func GetKafkaConsumerConfigSample() *KafkaConsumerConfig {
	k := NewKafkaConsumerConfig().
		SetEnabled(false).
		SetAppliedAuth(*GetKafkaAuthConfigSample()).
		AppendProperty("group.id", "consumer-group-id").
		AppendProperty("client.id", "consumer-client-id").
		AppendProperty("enable.auto.commit", true)
	return k
}
