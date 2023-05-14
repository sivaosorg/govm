package asterisk

import (
	"log"

	"github.com/sivaosorg/govm/utils"
)

func NewAsteriskConfig() *AsteriskConfig {
	a := &AsteriskConfig{}
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

func (a *AsteriskConfig) Json() string {
	return utils.ToJson(a)
}

func (t *TelephonyConfig) Json() string {
	return utils.ToJson(t)
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
