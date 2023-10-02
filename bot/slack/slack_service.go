package slack

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/sivaosorg/govm/builder"
	"github.com/sivaosorg/govm/restify"
)

type SlackService interface {
	SendMessage(message builder.MapBuilder) (builder.MapBuilder, error)
}

type slackServiceImpl struct {
	config SlackConfig       `json:"-"`
	option slackOptionConfig `json:"-"`
}

func NewSlackService(config SlackConfig) SlackService {
	s := &slackServiceImpl{
		config: config,
	}
	s.option = *NewSlackOptionConfig()
	return s
}

func NewSlackServiceWith(config SlackConfig, option slackOptionConfig) SlackService {
	s := &slackServiceImpl{
		config: config,
		option: option,
	}
	return s
}

func (s *slackServiceImpl) SendMessage(message builder.MapBuilder) (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Slack Bot service unavailable")
	}
	SlackConfigValidator(&s.config)
	clusters := []builder.MapBuilder{}
	var _err error
	var _response builder.MapBuilder
	var wg sync.WaitGroup
	wg.Add(len(s.config.ChannelId))
	for _, channelId := range s.config.ChannelId {
		if message.Contains(SecretKeyField) {
			message.Update(SecretKeyField, channelId)
		} else {
			message.Add(SecretKeyField, channelId)
		}
		clone, _ := builder.NewMapBuilder().DeserializeJsonI(message.Build())
		clusters = append(clusters, *clone)
	}
	for _, v := range clusters {
		go func(msg builder.MapBuilder) {
			defer wg.Done()
			_response, _err = s.sendText(msg)
		}(v)
	}
	wg.Wait()
	return _response, _err
}

func (s *slackServiceImpl) sendText(message builder.MapBuilder) (builder.MapBuilder, error) {
	if !s.config.IsEnabled {
		return *builder.NewMapBuilder(), fmt.Errorf("Slack Bot service unavailable")
	}
	url := fmt.Sprintf("%s/chat.postMessage", Host)
	client := restify.New()
	result := &map[string]interface{}{}
	client.SetRetryCount(s.option.MaxRetries).
		// Default is 100 milliseconds.
		SetRetryWaitTime(10 * time.Second).
		// Default is 2 seconds.
		SetRetryMaxWaitTime(20 * time.Second).
		AddRetryCondition(
			// RetryConditionFunc type is for retry condition function
			// input: non-nil Response OR request execution error
			func(r *restify.Response, err error) bool {
				response, _ := builder.NewMapBuilder().DeserializeJsonI(string(r.Body()))
				_, ok := response.Get("error")
				return (r.StatusCode() >= http.StatusBadRequest && r.StatusCode() <= http.StatusNetworkAuthenticationRequired) || ok
			},
		).
		SetDebug(s.config.DebugMode).
		SetHeaders(map[string]string{
			"Content-Type": "application/json;charset=utf-8",
		}).SetAuthToken(s.config.Token)

	_, err := client.R().
		SetResult(&result).
		SetBody(message.Build()).
		ForceContentType("application/json").
		Post(url)
	response, _ := builder.NewMapBuilder().DeserializeJsonI(result)
	return *response, err
}
