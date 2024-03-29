package telegram

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/sivaosorg/govm/blueprint"
	"github.com/sivaosorg/govm/builder"
	"github.com/sivaosorg/govm/restify"
)

type TelegramService interface {
	SendMessage(message interface{}) (builder.MapBuilder, error)
	SendMessageWith(message builder.MapBuilder) (builder.MapBuilder, error)
	SendFile(filename string, message ...string) (builder.MapBuilder, error)
	SendFiles(attachment map[string]string) (builder.MapBuilder, error)
	SendInlineKeyboard(message interface{}, buttons ...button) (builder.MapBuilder, error)

	// Get user information.
	GetMe() (builder.MapBuilder, error)
	// Receive incoming updates using long polling.
	GetUpdates(request builder.MapBuilder) (builder.MapBuilder, error)
	// Send text message by request body
	// Set customize request body
	SendMessageHandshake(request builder.MapBuilder) (builder.MapBuilder, error)
	// Send message as notification
	SendNotification(topic, message string) (builder.MapBuilder, error)
	// Send message as information
	SendInfo(topic, message string) (builder.MapBuilder, error)
	// Send message as warning
	SendWarning(topic, message string) (builder.MapBuilder, error)
	// Send message as error
	SendError(topic, message string) (builder.MapBuilder, error)
	// Send message as debug
	SendDebug(topic, message string) (builder.MapBuilder, error)
	// Send message as success
	SendSuccess(topic, message string) (builder.MapBuilder, error)
	// Send message as bug
	SendBug(topic, message string) (builder.MapBuilder, error)
	// Send message as trace
	SendTrace(topic, message string) (builder.MapBuilder, error)
}

type telegramServiceImpl struct {
	config TelegramConfig
	option telegramOptionConfig
}

func NewTelegramService(config TelegramConfig, option telegramOptionConfig) TelegramService {
	s := &telegramServiceImpl{
		config: config,
		option: option,
	}
	return s
}

func (s *telegramServiceImpl) SendMessage(message interface{}) (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Telegram Bot service unavailable")
	}
	TelegramConfigValidator(&s.config)
	var _err error
	var _response builder.MapBuilder
	var wg sync.WaitGroup
	for _, chatId := range s.config.ChatID {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()
			_response, _err = s.sendText(id, message)
		}(chatId)
	}
	wg.Wait()
	return _response, _err
}

func (s *telegramServiceImpl) SendMessageWith(message builder.MapBuilder) (builder.MapBuilder, error) {
	return s.SendMessage(message.Build())
}

func (s *telegramServiceImpl) SendFile(filename string, message ...string) (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Telegram Bot service unavailable")
	}
	TelegramConfigValidator(&s.config)
	var _err error
	var _response builder.MapBuilder
	var wg sync.WaitGroup
	for _, chatId := range s.config.ChatID {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()
			_response, _err = s.sendFile(id, filename, message...)
		}(chatId)
	}
	wg.Wait()
	return _response, _err
}

func (s *telegramServiceImpl) SendFiles(attachment map[string]string) (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Telegram Bot service unavailable")
	}
	if len(attachment) == 0 {
		return *builder.NewMapBuilder(), fmt.Errorf("No attachment attached")
	}
	TelegramConfigValidator(&s.config)
	var _err error
	var _response builder.MapBuilder
	var wg sync.WaitGroup
	for k, v := range attachment {
		wg.Add(1)
		go func(filename, message string) {
			defer wg.Done()
			_response, _err = s.SendFile(filename, message)
		}(k, v)
	}
	wg.Wait()
	return _response, _err
}

func (s *telegramServiceImpl) SendInlineKeyboard(message interface{}, buttons ...button) (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Telegram Bot service unavailable")
	}
	TelegramConfigValidator(&s.config)
	var _err error
	var _response builder.MapBuilder
	var wg sync.WaitGroup
	for _, chatId := range s.config.ChatID {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()
			_response, _err = s.sendInlineKeyboard(id, message, buttons)
		}(chatId)
	}
	wg.Wait()
	return _response, _err
}

func (s *telegramServiceImpl) GetMe() (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Telegram Bot service unavailable")
	}
	url := fmt.Sprintf("%s/bot%s/getMe", Host, s.config.Token)
	client := restify.New()
	var resultSuccess map[string]interface{}
	var resultFailure interface{}
	client.
		SetRetryCount(s.option.MaxRetries).
		// Default is 100 milliseconds.
		SetRetryWaitTime(10 * time.Second).
		// Default is 2 seconds.
		SetRetryMaxWaitTime(20 * time.Second).
		AddRetryCondition(
			// RetryConditionFunc type is for retry condition function
			// input: non-nil Response OR request execution error
			func(r *restify.Response, err error) bool {
				return (r.StatusCode() >= http.StatusBadRequest && r.StatusCode() <= http.StatusNetworkAuthenticationRequired)
			},
		).
		SetDebug(s.config.DebugMode).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		})

	payload := map[string]interface{}{}
	r, err := client.R().
		SetBody(payload).
		SetResult(&resultSuccess).
		SetError(&resultFailure).
		ForceContentType("application/json").
		Post(url)

	if r.IsError() {
		response, _ := builder.NewMapBuilder().DeserializeJsonI(resultFailure)
		return *response, err
	}
	if r.IsSuccess() {
		response, _ := builder.NewMapBuilder().DeserializeJsonI(resultSuccess)
		return *response, err
	}
	return *builder.NewMapBuilder(), err
}

func (s *telegramServiceImpl) GetUpdates(request builder.MapBuilder) (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Telegram Bot service unavailable")
	}
	url := fmt.Sprintf("%s/bot%s/getUpdates", Host, s.config.Token)
	client := restify.New()
	var resultSuccess map[string]interface{}
	var resultFailure interface{}
	client.
		SetRetryCount(s.option.MaxRetries).
		// Default is 100 milliseconds.
		SetRetryWaitTime(10 * time.Second).
		// Default is 2 seconds.
		SetRetryMaxWaitTime(20 * time.Second).
		AddRetryCondition(
			// RetryConditionFunc type is for retry condition function
			// input: non-nil Response OR request execution error
			func(r *restify.Response, err error) bool {
				return (r.StatusCode() >= http.StatusBadRequest && r.StatusCode() <= http.StatusNetworkAuthenticationRequired)
			},
		).
		SetDebug(s.config.DebugMode).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		})

	r, err := client.R().
		SetBody(request.Build()).
		SetResult(&resultSuccess).
		SetError(&resultFailure).
		ForceContentType("application/json").
		Post(url)

	if r.IsError() {
		response, _ := builder.NewMapBuilder().DeserializeJsonI(resultFailure)
		return *response, err
	}
	if r.IsSuccess() {
		response, _ := builder.NewMapBuilder().DeserializeJsonI(resultSuccess)
		return *response, err
	}
	return *builder.NewMapBuilder(), err
}

func (s *telegramServiceImpl) SendMessageHandshake(request builder.MapBuilder) (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Telegram Bot service unavailable")
	}
	if !request.Contains("chat_id") {
		return *builder.NewMapBuilder(), fmt.Errorf("chat_id is required")
	}
	if !request.Contains("text") {
		return *builder.NewMapBuilder(), fmt.Errorf("text is required")
	}
	url := fmt.Sprintf("%s/bot%s/sendMessage", Host, s.config.Token)
	client := restify.New()
	var resultSuccess map[string]interface{}
	var resultFailure interface{}
	client.
		SetRetryCount(s.option.MaxRetries).
		// Default is 100 milliseconds.
		SetRetryWaitTime(10 * time.Second).
		// Default is 2 seconds.
		SetRetryMaxWaitTime(20 * time.Second).
		AddRetryCondition(
			// RetryConditionFunc type is for retry condition function
			// input: non-nil Response OR request execution error
			func(r *restify.Response, err error) bool {
				return (r.StatusCode() >= http.StatusBadRequest && r.StatusCode() <= http.StatusNetworkAuthenticationRequired)
			},
		).
		SetDebug(s.config.DebugMode).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		})

	r, err := client.R().
		SetBody(request.Build()).
		SetResult(&resultSuccess).
		SetError(&resultFailure).
		ForceContentType("application/json").
		Post(url)

	if r.IsError() {
		response, _ := builder.NewMapBuilder().DeserializeJsonI(resultFailure)
		return *response, err
	}
	if r.IsSuccess() {
		response, _ := builder.NewMapBuilder().DeserializeJsonI(resultSuccess)
		return *response, err
	}
	return *builder.NewMapBuilder(), err
}

func (s *telegramServiceImpl) SendNotification(topic, message string) (builder.MapBuilder, error) {
	b := blueprint.WithCard(s.option.Timezone).SetIconText(blueprint.TypeNotification).SetDescription(message).SetTitle(topic)
	return s.SendMessage(b.GenCardDefault())
}

func (s *telegramServiceImpl) SendInfo(topic, message string) (builder.MapBuilder, error) {
	b := blueprint.WithCard(s.option.Timezone).SetIconText(blueprint.TypeInfo).SetDescription(message).SetTitle(topic)
	return s.SendMessage(b.GenCardDefault())
}

func (s *telegramServiceImpl) SendWarning(topic, message string) (builder.MapBuilder, error) {
	b := blueprint.WithCard(s.option.Timezone).SetIconText(blueprint.TypeWarning).SetDescription(message).SetTitle(topic)
	return s.SendMessage(b.GenCardDefault())
}

func (s *telegramServiceImpl) SendError(topic, message string) (builder.MapBuilder, error) {
	b := blueprint.WithCard(s.option.Timezone).SetIconText(blueprint.TypeError).SetDescription(message).SetTitle(topic)
	return s.SendMessage(b.GenCardDefault())
}

func (s *telegramServiceImpl) SendDebug(topic, message string) (builder.MapBuilder, error) {
	b := blueprint.WithCard(s.option.Timezone).SetIconText(blueprint.TypeDebug).SetDescription(message).SetTitle(topic)
	return s.SendMessage(b.GenCardDefault())
}

func (s *telegramServiceImpl) SendSuccess(topic, message string) (builder.MapBuilder, error) {
	b := blueprint.WithCard(s.option.Timezone).SetIconText(blueprint.TypeSuccess).SetDescription(message).SetTitle(topic)
	return s.SendMessage(b.GenCardDefault())
}

func (s *telegramServiceImpl) SendBug(topic, message string) (builder.MapBuilder, error) {
	b := blueprint.WithCard(s.option.Timezone).SetIconText(blueprint.TypeBug).SetDescription(message).SetTitle(topic)
	return s.SendMessage(b.GenCardDefault())
}

func (s *telegramServiceImpl) SendTrace(topic, message string) (builder.MapBuilder, error) {
	b := blueprint.WithCard(s.option.Timezone).SetIconText(blueprint.TypeTrace).SetDescription(message).SetTitle(topic)
	return s.SendMessage(b.GenCardDefault())
}

func (s *telegramServiceImpl) sendText(chatId int64, message interface{}) (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Telegram Bot service unavailable")
	}
	url := fmt.Sprintf("%s/bot%s/sendMessage", Host, s.config.Token)
	client := restify.New()
	var resultSuccess map[string]interface{}
	var resultFailure interface{}
	client.
		SetRetryCount(s.option.MaxRetries).
		// Default is 100 milliseconds.
		SetRetryWaitTime(10 * time.Second).
		// Default is 2 seconds.
		SetRetryMaxWaitTime(20 * time.Second).
		AddRetryCondition(
			// RetryConditionFunc type is for retry condition function
			// input: non-nil Response OR request execution error
			func(r *restify.Response, err error) bool {
				return (r.StatusCode() >= http.StatusBadRequest && r.StatusCode() <= http.StatusNetworkAuthenticationRequired)
			},
		).
		SetDebug(s.config.DebugMode).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		})

	payload := map[string]interface{}{
		"chat_id":    chatId,
		"text":       message,
		"parse_mode": s.option.Type,
	}
	r, err := client.R().
		SetBody(payload).
		SetResult(&resultSuccess).
		SetError(&resultFailure).
		ForceContentType("application/json").
		Post(url)

	if r.IsError() {
		response, _ := builder.NewMapBuilder().DeserializeJsonI(resultFailure)
		return *response, err
	}
	if r.IsSuccess() {
		response, _ := builder.NewMapBuilder().DeserializeJsonI(resultSuccess)
		return *response, err
	}
	return *builder.NewMapBuilder(), err
}

func (s *telegramServiceImpl) sendInlineKeyboard(chatId int64, message interface{}, buttons []button) (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Telegram Bot service unavailable")
	}
	var keyboard [][]inlineKeyboard
	for _, button := range buttons {
		k := []inlineKeyboard{}
		b := NewInlineKeyboard().SetText(button.Text).SetUrl(button.Url)
		k = append(k, *b)
		keyboard = append(keyboard, k)
	}
	url := fmt.Sprintf("%s/bot%s/sendMessage", Host, s.config.Token)
	client := restify.New()
	var resultSuccess map[string]interface{}
	var resultFailure interface{}
	client.
		SetRetryCount(s.option.MaxRetries).
		// Default is 100 milliseconds.
		SetRetryWaitTime(10 * time.Second).
		// Default is 2 seconds.
		SetRetryMaxWaitTime(20 * time.Second).
		AddRetryCondition(
			// RetryConditionFunc type is for retry condition function
			// input: non-nil Response OR request execution error
			func(r *restify.Response, err error) bool {
				return (r.StatusCode() >= http.StatusBadRequest && r.StatusCode() <= http.StatusNetworkAuthenticationRequired)
			},
		).
		SetDebug(s.config.DebugMode).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		})

	payload := map[string]interface{}{
		"chat_id":    chatId,
		"text":       message,
		"parse_mode": s.option.Type,
		"reply_markup": map[string]interface{}{
			"inline_keyboard": keyboard,
			"resize_keyboard": true,
		},
	}
	r, err := client.R().
		SetBody(payload).
		SetResult(&resultSuccess).
		SetError(&resultFailure).
		ForceContentType("application/json").
		Post(url)

	if r.IsError() {
		response, _ := builder.NewMapBuilder().DeserializeJsonI(resultFailure)
		return *response, err
	}
	if r.IsSuccess() {
		response, _ := builder.NewMapBuilder().DeserializeJsonI(resultSuccess)
		return *response, err
	}
	return *builder.NewMapBuilder(), err
}

func (s *telegramServiceImpl) sendFile(chatId int64, filename string, message ...string) (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Telegram Bot service unavailable")
	}
	url := fmt.Sprintf("%s/bot%s/sendDocument", Host, s.config.Token)
	client := restify.New()
	var resultSuccess map[string]interface{}
	var resultFailure interface{}
	client.
		SetRetryCount(s.option.MaxRetries).
		// Default is 100 milliseconds.
		SetRetryWaitTime(10 * time.Second).
		// Default is 2 seconds.
		SetRetryMaxWaitTime(20 * time.Second).
		AddRetryCondition(
			// RetryConditionFunc type is for retry condition function
			// input: non-nil Response OR request execution error
			func(r *restify.Response, err error) bool {
				return (r.StatusCode() >= http.StatusBadRequest && r.StatusCode() <= http.StatusNetworkAuthenticationRequired)
			},
		).
		SetDebug(s.config.DebugMode).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		})
	payload := map[string]string{
		"chat_id":    fmt.Sprintf("%v", chatId),
		"parse_mode": string(s.option.Type),
	}
	if len(message) > 0 {
		payload["caption"] = strings.Join(message, ",")
	}
	r, err := client.R().
		SetQueryParams(payload).
		SetResult(&resultSuccess).
		SetError(&resultFailure).
		SetFile("document", filename).
		ForceContentType("application/json").
		Post(url)

	if r.IsError() {
		response, _ := builder.NewMapBuilder().DeserializeJsonI(resultFailure)
		return *response, err
	}
	if r.IsSuccess() {
		response, _ := builder.NewMapBuilder().DeserializeJsonI(resultSuccess)
		return *response, err
	}
	return *builder.NewMapBuilder(), err
}
