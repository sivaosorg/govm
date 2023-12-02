package queues

type KafkaTopicConfig struct {
	IsEnabled         bool   `json:"enabled" yaml:"enabled"`
	Key               string `json:"key" binding:"required" yaml:"key"`            // Topic specified key for identifier
	Name              string `json:"name" yaml:"name"`                             // Topic name
	Description       string `json:"desc" yaml:"desc"`                             // Topic description
	CreateEnabled     bool   `json:"create_enabled" yaml:"create.enabled"`         // If set to true, allows automatic creation of topics.
	DeleteEnabled     bool   `json:"delete_enabled" yaml:"delete.enabled"`         // If set to true, allows topic deletion.
	ReplicationFactor int    `json:"replication_factor" yaml:"replication.factor"` // The replication factor for automatically created topics.
	Partitions        int    `json:"partitions" yaml:"partitions"`                 // The number of partitions for automatically created topics.
	RetentionMs       int    `json:"retention_ms" yaml:"retention.ms"`             // The retention time for messages in the topic. Set to -1 to retain messages indefinitely.
}

type KafkaAuthConfig struct {
	IsEnabled              bool     `json:"enabled" yaml:"enabled"`
	BootstrapServers       []string `json:"bootstrap_servers" yaml:"bootstrap.servers"`
	SecurityProtocol       string   `json:"security_protocol" yaml:"security.protocol"`
	SaslMechanism          string   `json:"sasl_mechanism" yaml:"sasl.mechanism"`
	SaslUsername           string   `json:"sasl_username" yaml:"sasl.username"`
	SaslPassword           string   `json:"-" yaml:"sasl.password"`
	SslCaLocation          string   `json:"ssl_ca_location" yaml:"ssl.ca.location"`
	SslCertificateLocation string   `json:"ssl_certificate_location" yaml:"ssl.certificate.location"`
	SslKeyLocation         string   `json:"ssl_key_location" yaml:"ssl.key.location"`
}

type KafkaConsumerConfig struct {
	IsEnabled   bool                   `json:"enabled" yaml:"enabled"`
	AppliedAuth KafkaAuthConfig        `json:"applied_auth" yaml:"applied_auth"`
	Props       map[string]interface{} `json:"properties" yaml:"properties"`
}

type KafkaProducerConfig struct {
	IsEnabled   bool                   `json:"enabled" yaml:"enabled"`
	AppliedAuth KafkaAuthConfig        `json:"applied_auth" yaml:"applied_auth"`
	Props       map[string]interface{} `json:"properties" yaml:"properties"`
}

type Kafka struct {
	IsEnabled   bool                `json:"enabled" yaml:"enabled"`
	Topics      []KafkaTopicConfig  `json:"topics" yaml:"topics"`
	AppliedAuth KafkaAuthConfig     `json:"applied_auth" yaml:"applied_auth"`
	Producer    KafkaProducerConfig `json:"producer" yaml:"producer"`
	Consumer    KafkaConsumerConfig `json:"consumer" yaml:"consumer"`
}

type MultiTenantKafkaConfig struct {
	Key             string `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool   `json:"usable_default" yaml:"usable_default"`
	Config          Kafka  `json:"config" yaml:"config"`
}

type ClusterMultiTenantKafkaConfig struct {
	Clusters []MultiTenantKafkaConfig `json:"clusters,omitempty" yaml:"clusters"`
}
