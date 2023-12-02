package queues

import "github.com/sivaosorg/govm/utils"

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

func GetKafkaSample() *Kafka {
	k := NewKafka().SetEnabled(false).
		SetAppliedAuth(*GetKafkaAuthConfigSample()).
		SetProducer(*GetKafkaProducerConfigSample()).
		SetConsumer(*GetKafkaConsumerConfigSample()).
		AppendTopics(*GetKafkaTopicConfigSample(), *GetKafkaTopicConfigSample().SetKey("topic-2"))
	return k
}
