package slack

import (
	"log"

	"github.com/sivaosorg/govm/utils"
)

func NewSlackConfig() *SlackConfig {
	s := &SlackConfig{}
	return s
}

func (s *SlackConfig) SetEnabled(value bool) *SlackConfig {
	s.IsEnabled = value
	return s
}

func (s *SlackConfig) SetDebugMode(value bool) *SlackConfig {
	s.DebugMode = value
	return s
}

func (s *SlackConfig) SetChannelId(values []string) *SlackConfig {
	if len(values) == 0 {
		log.Panicf("Invalid chat_id")
	}
	s.ChannelId = values
	return s
}

func (s *SlackConfig) AppendChannelId(values ...string) *SlackConfig {
	s.ChannelId = append(s.ChannelId, values...)
	return s
}

func (s *SlackConfig) SetToken(value string) *SlackConfig {
	if utils.IsEmpty(value) {
		log.Panicf("Invalid token: %v", value)
	}
	s.Token = value
	return s
}

func (s *SlackConfig) Json() string {
	return utils.ToJson(s)
}

func SlackConfigValidator(s *SlackConfig) {
	s.SetChannelId(s.ChannelId).SetToken(s.Token)
}

func GetSlackConfigSample() *SlackConfig {
	s := NewSlackConfig().
		SetEnabled(true).
		SetDebugMode(true).
		SetToken("<token-here>").
		AppendChannelId("123456789")
	return s
}

func NewSlackOptionConfig() *SlackOptionConfig {
	s := &SlackOptionConfig{}
	return s
}

func NewMultiTenantSlackConfig() *MultiTenantSlackConfig {
	m := &MultiTenantSlackConfig{}
	return m
}

func (m *MultiTenantSlackConfig) SetKey(value string) *MultiTenantSlackConfig {
	if utils.IsEmpty(value) {
		log.Panicf("Invalid key: %v", value)
	}
	m.Key = value
	return m
}

func (m *MultiTenantSlackConfig) SetUsableDefault(value bool) *MultiTenantSlackConfig {
	m.IsUsableDefault = value
	return m
}

func (m *MultiTenantSlackConfig) SetConfig(value SlackConfig) *MultiTenantSlackConfig {
	m.Config = value
	return m
}

func (m *MultiTenantSlackConfig) SetConfigCursor(value *SlackConfig) *MultiTenantSlackConfig {
	m.Config = *value
	return m
}

func (m *MultiTenantSlackConfig) SetOption(value SlackOptionConfig) *MultiTenantSlackConfig {
	m.Option = value
	return m
}

func (m *MultiTenantSlackConfig) Json() string {
	return utils.ToJson(m)
}

func MultiTenantSlackConfigValidator(m *MultiTenantSlackConfig) {
	m.SetKey(m.Key)
}

func GetMultiTenantSlackConfigSample() *MultiTenantSlackConfig {
	m := NewMultiTenantSlackConfig().
		SetKey("tenant_1").
		SetUsableDefault(false).
		SetConfigCursor(GetSlackConfigSample())
	return m
}

func NewClusterMultiTenantSlackConfig() *ClusterMultiTenantSlackConfig {
	c := &ClusterMultiTenantSlackConfig{}
	return c
}

func (c *ClusterMultiTenantSlackConfig) SetClusters(values []MultiTenantSlackConfig) *ClusterMultiTenantSlackConfig {
	c.Clusters = values
	return c
}

func (c *ClusterMultiTenantSlackConfig) AppendClusters(values ...MultiTenantSlackConfig) *ClusterMultiTenantSlackConfig {
	c.Clusters = append(c.Clusters, values...)
	return c
}

func (c *ClusterMultiTenantSlackConfig) Json() string {
	return utils.ToJson(c)
}

func GetClusterMultiTenantSlackConfigSample() *ClusterMultiTenantSlackConfig {
	c := NewClusterMultiTenantSlackConfig()
	c.AppendClusters(*GetMultiTenantSlackConfigSample(), *GetMultiTenantSlackConfigSample().SetKey("tenant_2"))
	return c
}