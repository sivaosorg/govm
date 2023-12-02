package queues

import (
	"fmt"

	"github.com/sivaosorg/govm/utils"
)

func NewKafkaTopicConfig() *KafkaTopicConfig {
	return &KafkaTopicConfig{}
}

func (k *KafkaTopicConfig) SetEnabled(value bool) *KafkaTopicConfig {
	k.IsEnabled = value
	return k
}

func (k *KafkaTopicConfig) SetKey(value string) *KafkaTopicConfig {
	k.Key = value
	return k
}

func (k *KafkaTopicConfig) SetName(value string) *KafkaTopicConfig {
	k.Name = value
	return k
}

func (k *KafkaTopicConfig) SetDescription(value string) *KafkaTopicConfig {
	k.Description = value
	return k
}

func (k *KafkaTopicConfig) SetCreateEnabled(value bool) *KafkaTopicConfig {
	k.CreateEnabled = value
	return k
}

func (k *KafkaTopicConfig) SetDeleteEnabled(value bool) *KafkaTopicConfig {
	k.DeleteEnabled = value
	return k
}

func (k *KafkaTopicConfig) SetReplicationFactor(value int) *KafkaTopicConfig {
	k.ReplicationFactor = value
	return k
}

func (k *KafkaTopicConfig) SetPartitions(value int) *KafkaTopicConfig {
	k.Partitions = value
	return k
}

func (k *KafkaTopicConfig) SetRetentionMs(value int) *KafkaTopicConfig {
	k.RetentionMs = value
	return k
}

func (k *KafkaTopicConfig) Json() string {
	return utils.ToJson(k)
}

func (k *KafkaTopicConfig) Available() bool {
	return k.IsEnabled
}

func KafkaTopicConfigValidator(k KafkaTopicConfig) error {
	if k.Partitions <= 0 {
		return fmt.Errorf("Invalid partitions: %v", k.Partitions)
	}
	if k.ReplicationFactor <= 0 {
		return fmt.Errorf("Invalid replication.factor: %v", k.ReplicationFactor)
	}
	if utils.IsEmpty(k.Key) {
		return fmt.Errorf("Topic key is required")
	}
	if utils.IsEmpty(k.Name) {
		return fmt.Errorf("Topic name is required")
	}
	return nil
}

func GetAvailableTopics(topics []KafkaTopicConfig) ([]KafkaTopicConfig, bool) {
	var response []KafkaTopicConfig
	if len(topics) == 0 {
		return response, false
	}
	for _, v := range topics {
		if !v.IsEnabled {
			continue
		}
		response = append(response, v)
	}
	if len(response) == 0 {
		return response, false
	}
	return response, true
}

func GetAvailableTopicsString(topics []KafkaTopicConfig) ([]string, bool) {
	response, ok := GetAvailableTopics(topics)
	if !ok {
		return []string{}, ok
	}
	var t []string
	for _, v := range response {
		t = append(t, v.Name)
	}
	return t, true
}

func GetTopicsAutoCreated(topics []KafkaTopicConfig) ([]KafkaTopicConfig, bool) {
	var response []KafkaTopicConfig
	if len(topics) == 0 {
		return response, false
	}
	for _, v := range topics {
		if !v.IsEnabled {
			continue
		}
		if !v.CreateEnabled {
			continue
		}
		response = append(response, v)
	}
	if len(response) == 0 {
		return response, false
	}
	return response, true
}

func GetTopic(topics []KafkaTopicConfig, key string) (KafkaTopicConfig, bool) {
	if len(topics) == 0 {
		return *NewKafkaTopicConfig(), false
	}
	for _, v := range topics {
		if v.Key == key {
			return v, true
		}
	}
	return *NewKafkaTopicConfig(), false
}

func GetKafkaTopicConfigSample() *KafkaTopicConfig {
	k := NewKafkaTopicConfig()
	k.IsEnabled = false
	k.Key = "topic-1"
	k.CreateEnabled = true
	k.DeleteEnabled = false
	k.ReplicationFactor = 3
	k.Partitions = 5
	k.RetentionMs = -1
	return k
}
