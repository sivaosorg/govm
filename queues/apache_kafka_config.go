package queues

const (
	// YAML configurations for Kafka Producer, along with explanations for some common and important configurations:

	// Comma-separated list of broker addresses.
	Bootstrap_Servers = "bootstrap.servers"
	// An ID string to pass to the server when making requests.
	Client_Id = "client.id"
	// Acknowledgment mode. "all" ensures that the record is successfully written to all replicas.
	Acks = "acks"
	// Number of retries before giving up sending a message.
	Retries = "retries"
	// Maximum time, in milliseconds, the producer will wait for an acknowledgment.
	Delivery_Timeout_Millisecond = "delivery.timeout.ms"
	// Time to wait for additional messages before sending a batch.
	Linger_Millisecond = "linger.ms"
	// Compression algorithm for message payloads.
	Compression_Type = "compression.type"
	// Maximum size of message batches.
	Batch_Size = "batch.size"
	// max.in.flight.requests.per.connection
	// Maximum number of in-flight requests per broker connection.
	Max_In_Flight_Requests_Per_Conn = "max.in.flight.requests.per.connection"
	// max.retries
	// Maximum number of message send retries.
	Max_Retries = "max.retries"
	// queue.buffering.max.messages
	// Maximum number of messages allowed in the producer queue.
	Queue_Buffering_Max_Messages = "queue.buffering.max.messages"
	// queue.buffering.max.ms
	// Maximum time, in milliseconds, a message can spend in the producer queue.
	Queue_Buffering_Max_Millisecond = "queue.buffering.max.ms"
	// message.send.max.retries
	// Maximum number of retries for a failed message.
	Message_Send_Max_Retries = "message.send.max.retries"
	// socket.keepalive.enable
	// Enable or disable socket keep-alive.
	Socket_Keepalive_Enabled = "socket.keepalive.enable"
	// enable.idempotence
	// Enable or disable idempotent producer.
	Enabled_Idempotence = "enable.idempotence"
	// security.protocol
	// Security protocol for communication with brokers.
	Security_Protocol = "security.protocol"
	// sasl.mechanism
	// SASL mechanism for authentication.
	Sasl_Mechanism = "sasl.mechanism"
	// sasl.username
	// SASL username
	Sasl_Username = "sasl.username"
	// sasl.password
	// SASL password.
	Sasl_Password = "sasl.password"
	// ssl.ca.location
	// Path to the CA certificate file for SSL.
	Ssl_Ca_Location = "ssl.ca.location"
	// ssl.certificate.location
	// Path to the client certificate file for SSL.
	Ssl_Certificate_Location = "ssl.certificate.location"
	// ssl.key.location
	// Path to the client private key file for SSL.
	Ssl_Key_Location = "ssl.key.location"

	// YAML configurations for Kafka Consumer, along with explanations for some common and important configurations:

	// group.id
	// Consumer group ID.
	GroupId = "group.id"
	// enable.auto.commit
	// If true, consumer offset is committed automatically at regular intervals.
	Enabled_Auto_Commit = "enable.auto.commit"
	// auto.commit.interval.ms
	// Interval at which offsets are committed when enable.auto.commit is true.
	Auto_Commit_Interval_Millisecond = "auto.commit.interval.ms"
	// session.timeout.ms
	// Maximum time, in milliseconds, spent in group re-balancing.
	Session_Timeout_Millisecond = "session.timeout.ms"
	// auto.offset.reset
	// Action to take when there is no initial offset in the offset store or the current offset no longer exists.
	Auto_Offset_Reset = "auto.offset.reset"
	// max.poll.interval.ms
	// Maximum time, in milliseconds, between two consecutive poll invocations.
	Max_Poll_Interval_Millisecond = "max.poll.interval.ms"
	// fetch.max.bytes
	// Maximum amount of data the server should return for a fetch request.
	Fetch_Max_Bytes = "fetch.max.bytes"
	// max.poll.records
	// Maximum number of records returned in a single call to poll.
	Max_Poll_Records = "max.poll.records"

	// Additional Notes:
	// 1. SASL Authentication: If you are using SASL for authentication, make sure to set the appropriate values for sasl.mechanism, sasl.username, and sasl.password.
	// 2. SSL Encryption: If SSL encryption is enabled, provide the paths to the CA certificate, client certificate, and client private key using ssl.ca.location, ssl.certificate.location, and ssl.key.location respectively.
	// 3. Producer Compression: compression.type allows you to choose the compression algorithm for message payloads. Options include "none," "gzip," "snappy," and "lz4."
	// 4. Consumer Offset Reset: auto.offset.reset determines what happens if there is no committed offset for a consumer group or if the committed offset is invalid. Options include "earliest," "latest," and "none."
	// 5. Consumer Polling: max.poll.records sets the maximum number of records returned in a single call to poll(). Adjust this based on your application's processing capacity.
	// 6. Consumer Re-balancing: session.timeout.ms is the maximum time a consumer in a group can go without sending a heartbeat to the broker. It plays a role in the group re-balancing process.
	// 7. Producer Idempotence: Enabling enable.idempotence ensures that messages sent to a partition are assigned increasing sequence numbers, providing stronger semantics guarantees.
	// 8. Error Handling: Implement robust error handling in your Golang code to handle various Kafka-related errors, retries, and logging.
	// 9. Security Protocols: security.protocol specifies the security protocol used for communication. Common options include "plaintext," "ssl," "sasl_plaintext," and "sasl_ssl."

	// Advanced Configurations:
	// 1. Producer Retries and Retry Backoff:
	// retries determines how many times the producer will attempt to resend a message in case of delivery failure.
	// Combine this with retry.backoff.ms (not explicitly mentioned in the previous response) to set the backoff time between retries.
	// retry.backoff.ms
	// Example:
	// retries: 3
	// retry.backoff.ms: 100
	Retry_Backoff_Millisecond = "retry.backoff.ms"

	// 2. Consumer Committing Offsets:
	// If enable.auto.commit is set to true, the consumer will automatically commit offsets at regular intervals.
	// Adjust auto.commit.interval.ms to control how often these commits occur.
	// Example:
	// enable.auto.commit: true
	// auto.commit.interval.ms: 1000

	// 3. Consumer Session Timeout:
	// The session.timeout.ms configuration in the consumer is critical for group management.
	// Ensure it is set to a value that allows sufficient time for rebalancing without causing unnecessary delays.

	// 4. Consumer Max Poll Interval:
	// max.poll.interval.ms is the maximum delay between consecutive calls to poll in the consumer.
	// Adjust this value based on the expected processing time for each poll.

	// 5. Producer Delivery Timeout:
	// delivery.timeout.ms is the maximum time, in milliseconds, that a producer will wait for a message to be acknowledged by the broker.
	// Ensure this value is appropriately set based on your application's requirements.

	// 6. Producer Batch Size:
	// batch.size is the maximum size of batches to use when sending messages.
	// Adjust this based on the expected size of your messages and the desired trade-off between latency and throughput.

	// 7. Consumer Fetch Configuration:
	// fetch.max.bytes determines the maximum amount of data the server should return for a fetch request.
	// Adjust this based on the expected size of your messages and the capacity of your consumers.
	// fetch.max.bytes: 1048576
	// fetch.max.wait.ms: 500

	// 8. Kafka Configuration Overrides:
	// You can also specify any Kafka configuration property directly in the YAML file.
	// For instance, if you need to set a specific Kafka property that is not explicitly provided in the library, you can do so like this:
	// kafka.property.name: "property-value"
	Kafka_Property_Name = "kafka.property.name"

	// 9. Partition Assignment Strategy:
	// When working with consumer groups, the partition.assignment.strategy configuration can be used to specify the strategy for assigning partitions to consumers.
	// The default is range, but you might want to explore other options like roundrobin based on your use case.
	// partition.assignment.strategy: "roundrobin"
	Partition_Assignment_Strategy = "partition.assignment.strategy"

	// 10.Consumer Heartbeat Configuration:
	// The heartbeat.interval.ms and heartbeat.timeout.ms configurations are related to the consumer's heartbeat mechanism.
	// Adjust these values to ensure timely heartbeats and avoid unnecessary rebalances.
	// heartbeat.interval.ms: 3000
	// heartbeat.timeout.ms: 10000
	Heartbeat_Interval_Millisecond = "heartbeat.interval.ms"
	Heartbeat_Timeout_Millisecond  = "heartbeat.timeout.ms"

	// 11. Producer Request Timeout:
	// The request.timeout.ms configuration in the producer determines the maximum time the producer will wait for an acknowledgement from the broker when sending a message.
	// request.timeout.ms: 5000
	Request_Timeout_Millisecond = "request.timeout.ms"

	// 12. Consumer Fetch Min Bytes:
	// fetch.min.bytes is the minimum amount of data the server should return for a fetch request.
	// Adjust this to control the trade-off between latency and efficiency in fetching messages.
	// fetch.min.bytes: 1
	Fetch_Min_Bytes = "fetch.min.bytes"

	// 13. Kafka Compression Configuration:
	// If you're using compression, such as snappy or gzip, you may want to adjust the compression level for better compression ratios or lower latencies.
	// compression.type: "snappy"
	// compression.level: 3
	Compression_Level = "compression.level"

	// 14. Producer Message Timeout:
	// The message.timeout.ms configuration in the producer is the maximum time a produced message is allowed to spend in the producer queue waiting for transmission.
	// message.timeout.ms: 3000
	Message_Timeout_Millisecond = "message.timeout.ms"

	// 15. Consumer Enable Partition EOF:
	// enable.partition.eof allows consumers to receive end-of-file (EOF) marker events when they reach the end of a partition.
	// enable.partition.eof: true
	Enabled_Partition_Eof = "enable.partition.eof"

	// 16. Consumer Max Session Duration:
	// max.session.timeout.ms is the upper limit on the amount of time a consumer's session can be in maintenance before it is considered failed.
	// max.session.timeout.ms: 300000
	Max_Session_Timeout_Millisecond = "max.session.timeout.ms"

	// 17. Producer Delivery Report Channel Size:
	// The go.produce.channel.size configuration sets the size of the Go channel used for delivering delivery reports from the producer.
	// go.produce.channel.size: 100
	Go_Produce_Channel_Size = "go.produce.channel.size"

	// 18. Consumer Fetch Max Wait Time:
	// fetch.wait.max.ms is the maximum amount of time the broker will wait for additional fetch response data before responding to the consumer.
	// fetch.wait.max.ms: 500
	Fetch_Wait_Max_Millisecond = "fetch.wait.max.ms"

	// 19. Producer Message Key and Value Serialization:
	// When producing messages, ensure that you configure the appropriate serializers for message keys and values.
	// Use key.serializer and value.serializer to specify the serializer classes.
	// key.serializer: "org.apache.kafka.common.serialization.StringSerializer"
	// value.serializer: "org.apache.kafka.common.serialization.StringSerializer"
	Key_Serializer   = "key.serializer"
	Value_Serializer = "value.serializer"

	// 20. Consumer Message Key and Value Deserialization:
	// Similarly, when consuming messages, configure the deserializer for message keys and values using key.deserializer and value.deserializer.
	// key.deserializer: "org.apache.kafka.common.serialization.StringDeserializer"
	// value.deserializer: "org.apache.kafka.common.serialization.StringDeserializer"
	Key_Deserializer   = "key.deserializer"
	Value_Deserializer = "value.deserializer"

	// 21. Consumer Seek to Beginning or End:
	// In certain scenarios, you may need to seek a consumer to the beginning or end of a partition. Use the seek.to.end or seek.to.beginning configuration.
	// seek.to.end: true
	Seek_To_End = "seek.to.end"

	// 22. Consumer Assign or Subscribe:
	// When manually assigning partitions to a consumer, use assign.
	// When using consumer groups and subscribing to topics, use subscribe.
	// mode: "assign"
	Mode = "mode"

	// 23. Message Timestamps:
	// Enable message timestamps to record the time a message is created or received.
	// Use message.timestamp.type and message.timestamp.difference.max.ms.
	// message.timestamp.type: "CreateTime"
	Message_Timestamp_Type                     = "message.timestamp.type"
	Message_Timeout_Difference_Max_Millisecond = "message.timestamp.difference.max.ms"

	// 24. Transaction Configuration:
	// If you're working with Kafka transactions, configure the transactional properties, such as transactional.id for producers and isolation.level for consumers.
	// transactional.id: "your-transactional-id"
	// isolation.level: "read_committed"
	Transaction_Id  = "transactional.id"
	Isolation_Level = "isolation.level"

	// 25. Producer Record Batching:
	// Fine-tune how records are batched using linger.ms (already mentioned) and batch.size.
	// Adjust these parameters for optimal batching behavior.
	// linger.ms: 10
	// batch.size: 16384

	// 26. Producer Max In-flight Requests:
	// The max.in.flight.requests.per.connection configuration in the producer specifies the maximum number of unacknowledged requests the producer will send on a single connection.
	// max.in.flight.requests.per.connection: 1

	// 27. Consumer Exclude Internal Topics:
	// Exclude internal topics from being subscribed by consumers using exclude.internal.topics.
	// exclude.internal.topics: true
	Exclude_Interval_Topics = "exclude.internal.topics"

	// 28. Producer Message Queue Buffering:
	// Adjust the producer configurations related to message buffering in the queue. These include queue.buffering.max.messages and queue.buffering.max.ms.
	// queue.buffering.max.messages: 100000
	// queue.buffering.max.ms: 100

	// 29. Security Protocol and SASL Configuration:
	// Ensure that you set the appropriate security protocol and SASL configurations based on your Kafka cluster's security settings.
	// security.protocol: "sasl_ssl"
	// sasl.mechanism: "PLAIN"
)
