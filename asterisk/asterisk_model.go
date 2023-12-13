package asterisk

import "time"

type TelephonyConfig struct {
	Region          string        `json:"region" yaml:"region"`
	PhonePrefix     []string      `json:"phone_prefix" yaml:"phone_prefix"`
	DigitExtensions []interface{} `json:"digit_extensions" yaml:"digit_extensions"`
}

type AsteriskConfig struct {
	IsEnabled bool            `json:"enabled" yaml:"enabled"`
	DebugMode bool            `json:"debug_mode" yaml:"debug_mode"`
	Port      int             `json:"port" binding:"required" yaml:"port"`
	Host      string          `json:"host" binding:"required" yaml:"host"`
	Username  string          `json:"username" binding:"required" yaml:"username"`
	Password  string          `json:"-" yaml:"password"`
	Telephony TelephonyConfig `json:"telephony" yaml:"telephony"`
	Timeout   time.Duration   `json:"timeout" yaml:"timeout"`
}

type asteriskOptionConfig struct {
}

type MultiTenantAsteriskConfig struct {
	Key             string               `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool                 `json:"usable_default" yaml:"usable_default"`
	Config          AsteriskConfig       `json:"config" yaml:"config"`
	Option          asteriskOptionConfig `json:"option,omitempty" yaml:"option"`
}

type ClusterMultiTenantAsteriskConfig struct {
	Clusters []MultiTenantAsteriskConfig `json:"clusters,omitempty" yaml:"clusters"`
}
