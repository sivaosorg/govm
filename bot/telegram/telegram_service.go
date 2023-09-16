package telegram

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/sivaosorg/govm/curlx"
	"github.com/sivaosorg/govm/entity"
)

type TelegramService interface {
	SendMessage(message interface{}) (interface{}, error)
	SendFile(filename string, message ...string) (interface{}, error)
	SendFiles(attachment map[string]string) (interface{}, error)
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

func (s *telegramServiceImpl) SendMessage(message interface{}) (interface{}, error) {
	if !s.config.IsEnabled {
		return nil, fmt.Errorf("Telegram Bot service unavailable")
	}
	TelegramConfigValidator(&s.config)
	var _err error
	var _response interface{}
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

func (s *telegramServiceImpl) SendFile(filename string, message ...string) (interface{}, error) {
	if !s.config.IsEnabled {
		return nil, fmt.Errorf("Telegram Bot service unavailable")
	}
	TelegramConfigValidator(&s.config)
	var _err error
	var _response interface{}
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

func (s *telegramServiceImpl) SendFiles(attachment map[string]string) (interface{}, error) {
	if !s.config.IsEnabled {
		return nil, fmt.Errorf("Telegram Bot service unavailable")
	}
	if len(attachment) == 0 {
		return nil, fmt.Errorf("No attachment attached")
	}
	TelegramConfigValidator(&s.config)
	var _err error
	var _response interface{}
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

func (s *telegramServiceImpl) sendText(id int64, message interface{}) (interface{}, error) {
	if !s.config.IsEnabled {
		return nil, fmt.Errorf("Telegram Bot service unavailable")
	}
	url := fmt.Sprintf("%s/bot%s/sendMessage", Host, s.config.Token)
	context := curlx.NewCurlxContext().
		SetBaseURL(url).
		AppendRetryContext(
			func(attempt int, response *http.Response, err error) bool {
				return (response != nil && entity.IsStatusCodeFailure(response.StatusCode)) ||
					response == nil || err != nil || response.StatusCode == 200
			},
		).
		SetMaxRetries(2)
	request := curlx.NewCurlxRequest().
		SetMethod(curlx.POST).
		SetEndpoint("").
		SetDebugMode(s.config.DebugMode).
		AppendHeader("Content-Type", "application/json")
	payload := map[string]interface{}{
		"chat_id":    id,
		"text":       message,
		"parse_mode": s.option.Type,
	}
	request.SetQueryParamsWith(payload)
	curl := curlx.NewCurlxService(*context, *request)
	err := curl.FetchWithProgress(nil)
	return request.ResponseError, err
}

func (s *telegramServiceImpl) sendFile(id int64, filename string, message ...string) (interface{}, error) {
	if !s.config.IsEnabled {
		return nil, fmt.Errorf("Telegram Bot service unavailable")
	}
	url := fmt.Sprintf("%s/bot%s/sendDocument", Host, s.config.Token)
	context := curlx.NewCurlxContext().
		SetBaseURL(url).
		AppendRetryContext(
			func(attempt int, response *http.Response, err error) bool {
				return (response != nil && entity.IsStatusCodeFailure(response.StatusCode)) ||
					response == nil || err != nil || response.StatusCode == 200
			},
		).
		SetMaxRetries(2)
	request := curlx.NewCurlxRequest().
		SetMethod(curlx.POST).
		SetEndpoint("").
		SetDebugMode(s.config.DebugMode).
		SetAttachFileField("document").
		SetAttachment(filename).
		AppendHeader("Content-Type", "application/json")
	payload := map[string]interface{}{
		"chat_id":    id,
		"parse_mode": s.option.Type,
	}
	if len(message) > 0 {
		payload["caption"] = strings.Join(message, ",")
	}
	request.SetQueryParamsWith(payload)
	curl := curlx.NewCurlxService(*context, *request)
	err := curl.FetchWithProgress(nil)
	return request.ResponseError, err
}
