package rabbitmqx

const (
	RabbitMqPrefix = "amqp"
)

const (
	ExchangeDirect  = "direct"
	ExchangeFanout  = "fanout"
	ExchangeTopic   = "topic"
	ExchangeHeaders = "headers"
)

var (
	Exchanges map[string]bool = map[string]bool{
		ExchangeDirect:  true,
		ExchangeFanout:  true,
		ExchangeTopic:   true,
		ExchangeHeaders: true,
	}
)
