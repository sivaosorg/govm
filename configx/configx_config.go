package configx

const (
	AsteriskConf             TypeConfig = "asterisk_conf"
	MongodbConf              TypeConfig = "mongodb_conf"
	MysqlConf                TypeConfig = "mysql_conf"
	PostgresConf             TypeConfig = "postgres_conf"
	RabbitMqConf             TypeConfig = "rabbit_mq_conf"
	RedisConf                TypeConfig = "redis_conf"
	WsConnOptionConf         TypeConfig = "ws_conn_option_conf"
	WsConnMessagePayloadConf TypeConfig = "ws_conn_message_payload_conf"
	WsConnSubscriptionConf   TypeConfig = "ws_conn_subs_conf"
)

var (
	ClustersConf map[TypeConfig]bool = map[TypeConfig]bool{
		AsteriskConf:             true,
		MongodbConf:              true,
		MysqlConf:                true,
		PostgresConf:             true,
		RabbitMqConf:             true,
		RedisConf:                true,
		WsConnOptionConf:         true,
		WsConnMessagePayloadConf: true,
		WsConnSubscriptionConf:   true,
	}
)

const (
	FilenameDefaultConf                   string = "./keys/default_conf.yaml"
	FilenameDefaultMultiTenantConf        string = "./keys/default_multi_tenant_conf.yaml"
	FilenameDefaultClusterMultiTenantConf string = "./keys/default_cluster_multi_tenant_conf.yaml"
)
