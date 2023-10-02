package asterisk

import (
	"fmt"
	"log"
	"time"

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

func (t *TelephonyConfig) SetPhonePrefix(values []string) *TelephonyConfig {
	t.PhonePrefix = values
	return t
}

func (t *TelephonyConfig) AppendPhonePrefix(values ...string) *TelephonyConfig {
	t.PhonePrefix = append(t.PhonePrefix, values...)
	return t
}

func (t *TelephonyConfig) SetDigitsExten(values []interface{}) *TelephonyConfig {
	t.DigitsExten = values
	return t
}

func (t *TelephonyConfig) AppendDigitExten(values ...interface{}) *TelephonyConfig {
	t.DigitsExten = append(t.DigitsExten, values...)
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

func GetAsteriskConfigSample() *AsteriskConfig {
	a := NewAsteriskConfig()
	t := NewTelephonyConfig()
	t.SetRegion("VN")
	t.AppendDigitExten(4, 5, 6)
	t.AppendPhonePrefix("9", "6064")
	a.SetEnabled(true)
	a.SetHost("http://127.0.0.1")
	a.SetPort(5038)
	a.SetUsername("u@root")
	a.SetPassword("pwd")
	a.SetTelephony(*t)
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
