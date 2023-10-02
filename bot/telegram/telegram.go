package telegram

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/sivaosorg/govm/utils"
)

func NewTelegramConfig() *TelegramConfig {
	t := &TelegramConfig{}
	t.SetTimeout(10 * time.Second)
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

func (t *TelegramConfig) SetTimeout(value time.Duration) *TelegramConfig {
	t.Timeout = value
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

func (m *MultiTenantTelegramConfig) SetOption(value telegramOptionConfig) *MultiTenantTelegramConfig {
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

func NewTelegramOptionConfig() *telegramOptionConfig {
	t := &telegramOptionConfig{}
	t.SetType(ModeMarkdown)
	t.SetMaxRetries(2)
	return t
}

func (t *telegramOptionConfig) SetType(value TelegramFormatType) *telegramOptionConfig {
	t.Type = value
	return t
}

func (t *telegramOptionConfig) SetTypeWith(value string) *telegramOptionConfig {
	if ok, _ := ModeText[TelegramFormatType(value)]; !ok || utils.IsEmpty(value) {
		log.Panicf("Invalid type: %s", value)
	}
	t.Type = TelegramFormatType(value)
	return t
}

func (t *telegramOptionConfig) SetMaxRetries(value int) *telegramOptionConfig {
	if value <= 0 {
		log.Panicf("Invalid max-retries: %v", value)
	}
	t.MaxRetries = value
	return t
}

func (t *telegramOptionConfig) Json() string {
	return utils.ToJson(t)
}

func (c *ClusterMultiTenantTelegramConfig) FindClusterBy(key string) (MultiTenantTelegramConfig, error) {
	if utils.IsEmpty(key) {
		return *NewMultiTenantTelegramConfig(), fmt.Errorf("Key is required")
	}
	if len(c.Clusters) == 0 {
		return *NewMultiTenantTelegramConfig(), fmt.Errorf("No telegram cluster")
	}
	for _, v := range c.Clusters {
		if v.Key == key {
			return v, nil
		}
	}
	return *NewMultiTenantTelegramConfig(), fmt.Errorf("The telegram cluster not found")
}

func NewButton() *button {
	b := &button{}
	return b
}

func (b *button) SetText(value string) *button {
	b.Text = value
	return b
}

func (b *button) SetUrl(value string) *button {
	_, err := url.Parse(value)
	if err != nil {
		log.Fatalf("Parse Url %v got an error: %v", value, err.Error())
	}
	b.Url = value
	return b
}

func (b *button) Json() string {
	return utils.ToJson(b)
}

func GetButtonSample() *button {
	b := NewButton().
		SetText("Click").
		SetUrl("https://www.google.com")
	return b
}

func NewInlineKeyboard() *inlineKeyboard {
	i := &inlineKeyboard{}
	return i
}

func (i *inlineKeyboard) SetText(value string) *inlineKeyboard {
	i.Text = value
	return i
}

func (i *inlineKeyboard) SetUrl(value string) *inlineKeyboard {
	_, err := url.Parse(value)
	if err != nil {
		log.Fatalf("Parse Url %v got an error: %v", value, err.Error())
	}
	i.Url = value
	return i
}

func (i *inlineKeyboard) SetPayload(value string) *inlineKeyboard {
	i.Payload = value
	return i
}

func (i *inlineKeyboard) Json() string {
	return utils.ToJson(i)
}

func GetInlineKeyboardButtonSample() *inlineKeyboard {
	i := NewInlineKeyboard().
		SetText("Click").
		SetUrl("https://www.google.com").
		SetPayload("Hook Trial")
	return i
}
