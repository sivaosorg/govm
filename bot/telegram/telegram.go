package telegram

import (
	"log"

	"github.com/sivaosorg/govm/utils"
)

func NewTelegramConfig() *TelegramConfig {
	t := &TelegramConfig{}
	return t
}

func (t *TelegramConfig) SetEnabled(value bool) *TelegramConfig {
	t.IsEnabled = value
	return t
}

func (t *TelegramConfig) SetDebugMode(value bool) *TelegramConfig {
	t.DebugMode = value
	return t
}

func (t *TelegramConfig) SetToken(value string) *TelegramConfig {
	if utils.IsEmpty(value) {
		log.Panicf("Invalid token: %s", value)
	}
	t.Token = value
	return t
}

func (t *TelegramConfig) SetChatId(values []int64) *TelegramConfig {
	if len(values) == 0 {
		log.Panicf("ChatID is required")
	}
	t.ChatID = values
	return t
}

func (t *TelegramConfig) AppendChatId(values ...int64) *TelegramConfig {
	t.ChatID = append(t.ChatID, values...)
	return t
}

func (t *TelegramConfig) Json() string {
	return utils.ToJson(t)
}

func TelegramConfigValidator(t *TelegramConfig) {
	t.SetToken(t.Token).SetChatId(t.ChatID)
}

func GetTelegramConfigSample() *TelegramConfig {
	t := NewTelegramConfig().
		SetEnabled(true).
		SetDebugMode(true).
		AppendChatId(123456789).
		SetToken("<token_here>")
	return t
}

func NewMultiTenantTelegramConfig() *MultiTenantTelegramConfig {
	m := &MultiTenantTelegramConfig{}
	m.SetUsableDefault(false)
	m.SetOption(*NewTelegramOptionConfig())
	return m
}

func (m *MultiTenantTelegramConfig) SetKey(value string) *MultiTenantTelegramConfig {
	if utils.IsEmpty(value) {
		log.Panicf("Invalid key: %v", value)
	}
	m.Key = value
	return m
}

func (m *MultiTenantTelegramConfig) SetUsableDefault(value bool) *MultiTenantTelegramConfig {
	m.IsUsableDefault = value
	return m
}

func (m *MultiTenantTelegramConfig) SetConfig(value TelegramConfig) *MultiTenantTelegramConfig {
	m.Config = value
	return m
}

func (m *MultiTenantTelegramConfig) SetConfigCursor(value *TelegramConfig) *MultiTenantTelegramConfig {
	m.Config = *value
	return m
}

func (m *MultiTenantTelegramConfig) SetOption(value TelegramOptionConfig) *MultiTenantTelegramConfig {
	m.Option = value
	return m
}

func (m *MultiTenantTelegramConfig) Json() string {
	return utils.ToJson(m)
}

func MultiTenantTelegramConfigValidator(m *MultiTenantTelegramConfig) {
	m.SetKey(m.Key)
}

func GetMultiTenantTelegramConfigSample() *MultiTenantTelegramConfig {
	m := NewMultiTenantTelegramConfig().SetKey("tenant_1").SetConfigCursor(GetTelegramConfigSample())
	return m
}

func NewClusterMultiTenantTelegramConfig() *ClusterMultiTenantTelegramConfig {
	c := &ClusterMultiTenantTelegramConfig{}
	return c
}

func (c *ClusterMultiTenantTelegramConfig) SetClusters(values []MultiTenantTelegramConfig) *ClusterMultiTenantTelegramConfig {
	c.Clusters = values
	return c
}

func (c *ClusterMultiTenantTelegramConfig) AppendClusters(values ...MultiTenantTelegramConfig) *ClusterMultiTenantTelegramConfig {
	c.Clusters = append(c.Clusters, values...)
	return c
}

func (c *ClusterMultiTenantTelegramConfig) Json() string {
	return utils.ToJson(c)
}

func GetClusterMultiTenantTelegramConfigSample() *ClusterMultiTenantTelegramConfig {
	c := NewClusterMultiTenantTelegramConfig()
	c.AppendClusters(*GetMultiTenantTelegramConfigSample(), *GetMultiTenantTelegramConfigSample().SetKey("tenant_2"))
	return c
}

func NewTelegramOptionConfig() *TelegramOptionConfig {
	t := &TelegramOptionConfig{}
	t.SetType(ModeMarkdown)
	return t
}

func (t *TelegramOptionConfig) SetType(value TelegramFormatType) *TelegramOptionConfig {
	t.Type = value
	return t
}

func (t *TelegramOptionConfig) SetTypeWith(value string) *TelegramOptionConfig {
	if ok, _ := ModeText[TelegramFormatType(value)]; !ok || utils.IsEmpty(value) {
		log.Panicf("Invalid type: %s", value)
	}
	t.Type = TelegramFormatType(value)
	return t
}

func (t *TelegramOptionConfig) Json() string {
	return utils.ToJson(t)
}
