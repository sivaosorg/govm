package wsconnx

import (
	"log"
	"time"

	"github.com/sivaosorg/govm/utils"
)

func NewWsConnOptionConfig() *WsConnOptionConfig {
	w := &WsConnOptionConfig{}
	return w
}

func NewWsConnMessagePayload() *WsConnMessagePayload {
	w := &WsConnMessagePayload{}
	return w
}

func NewWsConnSubscription() *WsConnSubscription {
	w := &WsConnSubscription{}
	return w
}

func (w *WsConnOptionConfig) SetEnabled(value bool) *WsConnOptionConfig {
	w.IsEnabled = value
	return w
}

func (w *WsConnOptionConfig) SetMaxMessageSize(value int) *WsConnOptionConfig {
	if value <= 0 {
		log.Panicf("Invalid max-message-size: %d", value)
	}
	w.MaxMessageSize = value
	return w
}

func (w *WsConnOptionConfig) SetWriteWait(value time.Duration) *WsConnOptionConfig {
	w.WriteWait = value
	return w
}

func (w *WsConnOptionConfig) SetPongWait(value time.Duration) *WsConnOptionConfig {
	w.PongWait = value
	return w
}

func (w *WsConnOptionConfig) SetPingPeriod(value time.Duration) *WsConnOptionConfig {
	w.PingPeriod = value
	return w
}

func (w *WsConnOptionConfig) Json() string {
	return utils.ToJson(w)
}

func (w *WsConnMessagePayload) SetTopic(value string) *WsConnMessagePayload {
	if utils.IsEmpty(value) {
		log.Panicf("Topic is required")
	}
	w.Topic = value
	return w
}

func (w *WsConnMessagePayload) SetContent(value interface{}) *WsConnMessagePayload {
	w.Content = value
	return w
}

func (w *WsConnMessagePayload) SetGenesisTimestamp(value time.Time) *WsConnMessagePayload {
	if value.IsZero() {
		log.Panicf("Genesis timestamp is required")
	}
	w.GenesisTimestamp = value
	return w
}

func (w *WsConnMessagePayload) Json() string {
	return utils.ToJson(w)
}

func (w *WsConnSubscription) SetTopic(value string) *WsConnSubscription {
	if utils.IsEmpty(value) {
		log.Panicf("Topic is required")
	}
	w.Topic = value
	return w
}

func (w *WsConnSubscription) SetUserId(value string) *WsConnSubscription {
	w.UserId = value
	return w
}

func (w *WsConnSubscription) SetSubscribedAt(value time.Time) *WsConnSubscription {
	w.SubscribedAt = value
	return w
}

func (w *WsConnSubscription) SetExpiredAt(value time.Time) *WsConnSubscription {
	w.ExpiredAt = value
	return w
}

func (w *WsConnSubscription) SetPersistent(value bool) *WsConnSubscription {
	w.IsPersistent = value
	return w
}

func (w *WsConnSubscription) SetMetaData(value map[string]interface{}) *WsConnSubscription {
	w.Metadata = value
	return w
}

func (w *WsConnSubscription) SetScope(value string) *WsConnSubscription {
	if utils.IsEmpty(value) {
		log.Panicf("Scope is required")
	}
	if !w.IsScope(value) {
		log.Panicf("Invalid scope: %s", value)
	}
	w.Scope = value
	return w
}

func (w *WsConnSubscription) SetStatus(value string) *WsConnSubscription {
	if utils.IsEmpty(value) {
		log.Panicf("Status is required")
	}
	if !w.IsStatus(value) {
		log.Panicf("Invalid status: %s", value)
	}
	w.Status = value
	return w
}

func (w *WsConnSubscription) SetContentType(value string) *WsConnSubscription {
	if utils.IsNotEmpty(value) {
		w.ContentType = utils.TrimSpaces(value)
	}
	return w
}

func (w *WsConnSubscription) SetFrequency(value int) *WsConnSubscription {
	if value <= 0 {
		log.Panicf("Invalid frequency: %v", value)
	}
	w.Frequency = value
	return w
}

func (w *WsConnSubscription) SetPriority(value float64) *WsConnSubscription {
	if value < 0 {
		log.Panicf("Invalid priority: %v", value)
	}
	w.Priority = value
	return w
}

func (w *WsConnSubscription) SetEvents(values []string) *WsConnSubscription {
	w.Events = values
	return w
}

func (w *WsConnSubscription) AppendEvents(values ...string) *WsConnSubscription {
	w.SetEvents(values)
	return w
}

func (w *WsConnSubscription) SetData(value []map[string]interface{}) *WsConnSubscription {
	w.Data = value
	return w
}

func (w *WsConnSubscription) AppendData(values ...map[string]interface{}) *WsConnSubscription {
	w.SetData(values)
	return w
}

// Checks if the provided WebSocket connection status is valid.
func (w *WsConnSubscription) IsStatus(status string) bool {
	switch status {
	case StatusConnected, StatusDisconnected, StatusReconnecting,
		StatusClosed, StatusConnecting, StatusAuthenticated,
		StatusFailed, StatusTerminated, StatusIdle,
		StatusBusy, StatusError, StatusReconnected,
		StatusStale:
		return true
	default:
		return false
	}
}

// Checks if the provided WebSocket connection scope is valid.
func (w *WsConnSubscription) IsScope(scope string) bool {
	switch scope {
	case ScopePublic, ScopePrivate, ScopeGroup,
		ScopeRead, ScopeWrite, ScopeReadWrite,
		ScopeAdmin:
		return true
	default:
		return false
	}
}

func (w *WsConnSubscription) Json() string {
	return utils.ToJson(w)
}

func WsConnOptionConfigValidator(w *WsConnOptionConfig) {
	w.SetMaxMessageSize(w.MaxMessageSize)
}

func WsConnMessagePayloadValidator(w *WsConnMessagePayload) {
	w.SetTopic(w.Topic).SetGenesisTimestamp(w.GenesisTimestamp)
}

func WsConnSubscriptionValidator(w *WsConnSubscription) {
	w.SetTopic(w.Topic)
}

func GetWsConnOptionConfigSample() *WsConnOptionConfig {
	w := NewWsConnOptionConfig().
		SetEnabled(true).
		SetMaxMessageSize(512).
		SetWriteWait(10 * time.Second).
		SetPongWait(60 * time.Second)
	w.SetPingPeriod((w.PongWait * 9) / 10)
	return w
}

func GetWsConnMessagePayloadSample() *WsConnMessagePayload {
	w := NewWsConnMessagePayload().
		SetTopic("topic_wsconnx").
		SetContent("websocket_payload").
		SetGenesisTimestamp(time.Now())
	return w
}

func GetWsConnSubscriptionSample() *WsConnSubscription {
	w := NewWsConnSubscription().
		SetTopic("topic_wsconnx").
		SetUserId("user_wsconnx").
		SetSubscribedAt(time.Now()).
		SetExpiredAt(time.Now().AddDate(0, 0, 10)).
		SetPersistent(true).
		SetScope(ScopePublic).
		SetStatus(StatusConnected).
		SetContentType("application/json").
		SetFrequency(10).
		SetPriority(1).
		AppendEvents("topic_created", "topic_updated", "coupons_created")
	return w
}
