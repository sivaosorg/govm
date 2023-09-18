package telegram

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/sivaosorg/govm/builder"
	"github.com/sivaosorg/govm/restify"
)

type TelegramService interface {
	SendMessage(message interface{}) (builder.MapBuilder, error)
	SendMessageWith(message builder.MapBuilder) (builder.MapBuilder, error)
	SendFile(filename string, message ...string) (builder.MapBuilder, error)
	SendFiles(attachment map[string]string) (builder.MapBuilder, error)
}

type telegramServiceImpl struct {
	config TelegramConfig       `json:"-"`
	option TelegramOptionConfig `json:"-"`
}

func NewTelegramService(config TelegramConfig, option TelegramOptionConfig) TelegramService {
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

func (s *telegramServiceImpl) sendText(chatId int64, message interface{}) (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Telegram Bot service unavailable")
	}
	url := fmt.Sprintf("%s/bot%s/sendMessage", Host, s.config.Token)
	client := restify.New()
	result := &map[string]interface{}{}
	client.
		SetRetryCount(2).
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
	_, err := client.R().
		SetResult(&result).
		SetBody(payload).
		ForceContentType("application/json").
		Post(url)
	response, _ := builder.NewMapBuilder().DeserializeJsonI(result)
	return *response, err
}

func (s *telegramServiceImpl) sendFile(chatId int64, filename string, message ...string) (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Telegram Bot service unavailable")
	}
	url := fmt.Sprintf("%s/bot%s/sendDocument", Host, s.config.Token)
	client := restify.New()
	result := &map[string]interface{}{}
	client.
		SetRetryCount(2).
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
	_, err := client.R().
		SetResult(&result).
		SetQueryParams(payload).
		SetFile("document", filename).
		ForceContentType("application/json").
		Post(url)
	response, _ := builder.NewMapBuilder().DeserializeJsonI(result)
	return *response, err
}
