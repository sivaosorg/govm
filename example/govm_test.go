package example

import (
	"testing"

	"github.com/sivaosorg/govm/bot/telegram"
	"github.com/sivaosorg/govm/logger"
	"github.com/sivaosorg/govm/query"
	"github.com/sivaosorg/govm/timex"
)

// bot: t.me/javis_notify_forum_bot
// group: https://t.me/javis_forum_bot
// chat_id: -1002042977093
// token: 6806983892:AAGcPZiuNktLFnyVWrRyOyYssECcVmNJSRo
func createTelegramService() telegram.TelegramService {
	options := telegram.NewTelegramOptionConfig().
		SetType(telegram.ModeHTML).
		SetTimezone(timex.DefaultTimezoneVietnam)
	svc := telegram.NewTelegramService(*telegram.GetTelegramConfigSample().
		SetChatId([]int64{-1002042977093}).
		SetToken("6806983892:AAGcPZiuNktLFnyVWrRyOyYssECcVmNJSRo").
		SetDebugMode(false),
		*options)

	return svc
}

func TestCardNotification(t *testing.T) {
	svc := createTelegramService()
	svc.SendWarning("Kafka Stream", "Kafka Streams is a part of the Apache Kafka project that enables developers to build real-time processing applications, where data can be ingested, processed, and transformed in real-time as it flows through the Kafka cluster. Kafka Streams is a library for building scalable and fault-tolerant stream processing applications without the need for a separate processing cluster.")
}

func TestDecisionModify(t *testing.T) {
	m := query.NewModify().
		Add("username_column", *query.WithDecision(true).SetValue(1)).
		Add("password_column", *query.NewDecision().SetEnabled(true).SetValue("123"))
	logger.Infof("modifies: %v", m.Transform())
}
