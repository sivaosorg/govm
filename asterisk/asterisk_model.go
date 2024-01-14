package asterisk

import "time"

type TelephonyConfig struct {
	Region               string        `json:"region" yaml:"region"`
	Timezone             string        `json:"timezone" yaml:"timezone"`
	TimeFormat           string        `json:"time_format" yaml:"time_format"`
	PhonePrefixes        []string      `json:"phone_prefixes" yaml:"phone_prefixes"` // phone prefixes which removed from phone. number
	ApplyMaxExtension    []interface{} `json:"apply_max_extension" yaml:"apply_max_extension"`
	ExceptionalExtension []string      `json:"exceptional_extension" yaml:"exceptional_extension"`
}

type SettingConfig struct {
}

type CacheConfig struct {
}

type AsteriskConfig struct {
	IsEnabled bool            `json:"enabled" yaml:"enabled"`
	DebugMode bool            `json:"debug_mode" yaml:"debug_mode"`
	Port      int             `json:"port" binding:"required" yaml:"port"`
	Host      string          `json:"host" binding:"required" yaml:"host"`
	Username  string          `json:"username" binding:"required" yaml:"username"`
	Password  string          `json:"-" yaml:"password"`
	Telephony TelephonyConfig `json:"telephony" yaml:"telephony"`
	Setting   SettingConfig   `json:"setting" yaml:"setting"`
	Caches    CacheConfig     `json:"cache" yaml:"cache"`
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
