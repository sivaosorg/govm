package asterisk

import (
	"fmt"
	"log"
	"time"

	"github.com/sivaosorg/govm/timex"
	"github.com/sivaosorg/govm/utils"
)

func NewAsteriskConfig() *AsteriskConfig {
	a := &AsteriskConfig{}
	a.SetTimeout(10 * time.Second) // default timeout 10s
	return a
}

func NewTelephonyConfig() *TelephonyConfig {
	t := &TelephonyConfig{}
	return t
}

func (a *AsteriskConfig) SetEnabled(value bool) *AsteriskConfig {
	a.IsEnabled = value
	return a
}

func (a *AsteriskConfig) SetPort(value int) *AsteriskConfig {
	if value <= 0 {
		log.Panic("Invalid port")
	}
	a.Port = value
	return a
}

func (a *AsteriskConfig) SetHost(value string) *AsteriskConfig {
	a.Host = utils.TrimSpaces(value)
	return a
}

func (a *AsteriskConfig) SetUsername(value string) *AsteriskConfig {
	a.Username = value
	return a
}

func (a *AsteriskConfig) SetPassword(value string) *AsteriskConfig {
	a.Password = value
	return a
}

func (t *TelephonyConfig) SetRegion(value string) *TelephonyConfig {
	t.Region = utils.TrimSpaces(value)
	return t
}

func (t *TelephonyConfig) SetPhonePrefixes(values []string) *TelephonyConfig {
	t.PhonePrefixes = values
	return t
}

func (t *TelephonyConfig) AppendPhonePrefixes(values ...string) *TelephonyConfig {
	t.PhonePrefixes = append(t.PhonePrefixes, values...)
	return t
}

func (t *TelephonyConfig) SetApplyMaxExtension(values []interface{}) *TelephonyConfig {
	t.ApplyMaxExtension = values
	return t
}

func (t *TelephonyConfig) AppendApplyMaxExtension(values ...interface{}) *TelephonyConfig {
	t.ApplyMaxExtension = append(t.ApplyMaxExtension, values...)
	return t
}

func (t *TelephonyConfig) SetExceptionalExtension(values []string) *TelephonyConfig {
	t.ExceptionalExtension = values
	return t
}

func (t *TelephonyConfig) AppendExceptionalExtension(values ...string) *TelephonyConfig {
	t.ExceptionalExtension = append(t.ExceptionalExtension, values...)
	return t
}

func (t *TelephonyConfig) SetTimezone(value string) *TelephonyConfig {
	t.Timezone = value
	return t
}

func (t *TelephonyConfig) SetTimeFormat(value string) *TelephonyConfig {
	t.TimeFormat = value
	return t
}

func (a *AsteriskConfig) SetTelephony(value TelephonyConfig) *AsteriskConfig {
	a.Telephony = value
	return a
}

func (a *AsteriskConfig) SetDebugMode(value bool) *AsteriskConfig {
	a.DebugMode = value
	return a
}

func (a *AsteriskConfig) SetTimeout(value time.Duration) *AsteriskConfig {
	a.Timeout = value
	return a
}

func (a *AsteriskConfig) Json() string {
	return utils.ToJson(a)
}

func (t *TelephonyConfig) Json() string {
	return utils.ToJson(t)
}

func AsteriskConfigValidator(a *AsteriskConfig) {
	a.SetPort(a.Port)
}

func GetTelephonyConfigSample() *TelephonyConfig {
	t := NewTelephonyConfig()
	t.SetRegion("VN")
	t.AppendApplyMaxExtension(4, 5)
	t.AppendPhonePrefixes("9")
	t.SetTimezone(timex.DefaultTimezoneVietnam)
	t.SetTimeFormat(timex.TimeFormat20060102150405)
	t.AppendExceptionalExtension("022123456XX..XX")
	return t
}

func GetAsteriskConfigSample() *AsteriskConfig {
	a := NewAsteriskConfig()
	a.SetEnabled(true)
	a.SetHost("127.0.0.1")
	a.SetPort(5038)
	a.SetUsername("u@root")
	a.SetPassword("pwd")
	a.SetTelephony(*GetTelephonyConfigSample())
	return a
}

func NewMultiTenantAsteriskConfig() *MultiTenantAsteriskConfig {
	m := &MultiTenantAsteriskConfig{}
	return m
}

func (m *MultiTenantAsteriskConfig) SetKey(value string) *MultiTenantAsteriskConfig {
	if utils.IsEmpty(value) {
		log.Panicf("Key is required")
	}
	m.Key = value
	return m
}

func (m *MultiTenantAsteriskConfig) SetUsableDefault(value bool) *MultiTenantAsteriskConfig {
	m.IsUsableDefault = value
	return m
}

func (m *MultiTenantAsteriskConfig) SetConfig(value AsteriskConfig) *MultiTenantAsteriskConfig {
	m.Config = value
	return m
}

func (m *MultiTenantAsteriskConfig) SetConfigCursor(value *AsteriskConfig) *MultiTenantAsteriskConfig {
	m.Config = *value
	return m
}

func (m *MultiTenantAsteriskConfig) SetOption(value asteriskOptionConfig) *MultiTenantAsteriskConfig {
	m.Option = value
	return m
}

func (m *MultiTenantAsteriskConfig) Json() string {
	return utils.ToJson(m)
}

func MultiTenantAsteriskConfigValidator(m *MultiTenantAsteriskConfig) {
	m.SetKey(m.Key)
}

func GetMultiTenantAsteriskConfigSample() *MultiTenantAsteriskConfig {
	m := NewMultiTenantAsteriskConfig().
		SetKey("tenant_1").
		SetUsableDefault(false).
		SetConfigCursor(GetAsteriskConfigSample())
	return m
}

func NewClusterMultiTenantAsteriskConfig() *ClusterMultiTenantAsteriskConfig {
	c := &ClusterMultiTenantAsteriskConfig{}
	return c
}

func (c *ClusterMultiTenantAsteriskConfig) SetClusters(values []MultiTenantAsteriskConfig) *ClusterMultiTenantAsteriskConfig {
	c.Clusters = values
	return c
}

func (c *ClusterMultiTenantAsteriskConfig) AppendClusters(values ...MultiTenantAsteriskConfig) *ClusterMultiTenantAsteriskConfig {
	c.Clusters = append(c.Clusters, values...)
	return c
}

func (c *ClusterMultiTenantAsteriskConfig) Json() string {
	return utils.ToJson(c.Clusters)
}

func GetClusterMultiTenantAsteriskConfigSample() *ClusterMultiTenantAsteriskConfig {
	c := NewClusterMultiTenantAsteriskConfig()
	c.AppendClusters(*GetMultiTenantAsteriskConfigSample(), *GetMultiTenantAsteriskConfigSample().SetKey("tenant_2"))
	return c
}

func (c *ClusterMultiTenantAsteriskConfig) FindClusterBy(key string) (MultiTenantAsteriskConfig, error) {
	if utils.IsEmpty(key) {
		return *NewMultiTenantAsteriskConfig(), fmt.Errorf("Key is required")
	}
	if len(c.Clusters) == 0 {
		return *NewMultiTenantAsteriskConfig(), fmt.Errorf("No asterisk cluster")
	}
	for _, v := range c.Clusters {
		if v.Key == key {
			return v, nil
		}
	}
	return *NewMultiTenantAsteriskConfig(), fmt.Errorf("The asterisk cluster not found")
}

func NewSettingConfig() *SettingConfig {
	s := &SettingConfig{}
	return s
}

func NewCacheConfig() *CacheConfig {
	c := &CacheConfig{}
	return c
}
