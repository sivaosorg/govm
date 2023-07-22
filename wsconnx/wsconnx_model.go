package wsconnx

import "time"

type WsConnOptionConfig struct {
	IsEnabled      bool          `json:"enabled" yaml:"enabled"`
	MaxMessageSize int           `json:"max_message_size" binding:"required" yaml:"max-message-size"`
	WriteWait      time.Duration `json:"write_wait"`
	PongWait       time.Duration `json:"pong_wait"`
	PingPeriod     time.Duration `json:"ping_period"`
}

type WsConnMessagePayload struct {
	Topic            string      `json:"topic" binding:"required"`
	Content          interface{} `json:"content,omitempty"`
	GenesisTimestamp time.Time   `json:"genesis_timestamp" binding:"required"`
}

// Subscription represents a user's subscription to a WebSocket topic.
type WsConnSubscription struct {
	// Topic is the name of the WebSocket topic the user is subscribed to.
	Topic string `json:"topic" binding:"required"`

	// UserId is the unique identifier of the user associated with this subscription.
	UserId string `json:"user_id,omitempty"`

	// Subscribed is the timestamp when the user subscribed to the topic.
	SubscribedAt time.Time `json:"subscribed_at,omitempty"`

	// Expiration is the timestamp when the subscription will expire.
	ExpiredAt time.Time `json:"expired_at,omitempty"`

	// IsPersistent indicates if the subscription persists even after the user disconnects.
	IsPersistent bool `json:"is_persistent,omitempty"`

	// Metadata is a map containing additional metadata associated with the subscription.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Scope defines the connection scope (e.g., "public", "private") of the subscription.
	Scope string `json:"scope"`

	// Status represents the status of the subscription (e.g., "active", "inactive", "expired").
	Status string `json:"status"`

	// ContentType specifies the content type of the subscription data (e.g., "json", "text").
	ContentType string `json:"content_type,omitempty"`

	// Frequency is an optional field that defines the frequency of receiving notifications.
	Frequency int `json:"frequency,omitempty"`

	// Priority is an optional field that defines the priority of the subscription.
	Priority float64 `json:"priority,omitempty"`

	// Events is a list of specific events the user is interested in receiving.
	Events []string `json:"events,omitempty"`

	// Data is an optional array of complex data objects associated with the subscription.
	Data []map[string]interface{} `json:"data,omitempty"`
}
