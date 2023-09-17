package slack

import (
	"fmt"

	"github.com/sivaosorg/govm/builder"
	"github.com/sivaosorg/govm/curlx"
)

type SlackService interface {
	SendMessage(message builder.MapBuilder) (interface{}, error)
}

type slackServiceImpl struct {
	config SlackConfig       `json:"-"`
	option SlackOptionConfig `json:"-"`
}

func NewSlackService(config SlackConfig) SlackService {
	s := &slackServiceImpl{
		config: config,
	}
	return s
}

func NewSlackServiceWith(config SlackConfig, option SlackOptionConfig) SlackService {
	s := &slackServiceImpl{
		config: config,
		option: option,
	}
	return s
}

func (s *slackServiceImpl) SendMessage(message builder.MapBuilder) (interface{}, error) {
	if !s.config.IsEnabled {
		return nil, fmt.Errorf("Slack Bot service unavailable")
	}
	SlackConfigValidator(&s.config)
	var _err error
	var _response interface{}
	for _, channelId := range s.config.ChannelId {
		_response, _err = s.sendText(channelId, message)
	}
	return _response, _err
}

func (s *slackServiceImpl) sendText(channelID string, message builder.MapBuilder) (interface{}, error) {
	if !s.config.IsEnabled {
		return nil, fmt.Errorf("Slack Bot service unavailable")
	}
	url := fmt.Sprintf("%s", Host)
	context := curlx.NewCurlxContext().
		SetBaseURL(url).
		SetMaxRetries(2)
	request := curlx.NewCurlxRequest().
		SetMethod(curlx.POST).
		SetEndpoint("/chat.postMessage").
		SetDebugMode(s.config.DebugMode).
		AppendHeader("Content-Type", "application/json;charset=utf-8").
		AppendHeader("Authorization", fmt.Sprintf("Bearer %s", s.config.Token))
	message.AddKeyValue("channel", channelID)
	request.SetRequestBody(message.Build())
	curl := curlx.NewCurlxService(*context, *request)
	err := curl.FetchWithProgress(nil)
	return request.ResponseBody, err
}
