package asterisk

type TelephonyConfig struct {
	Region      string        `json:"region" yaml:"region"`
	PhonePrefix []string      `json:"phone_prefix" yaml:"phone-prefix"`
	DigitsExten []interface{} `json:"digits_exten" yaml:"digits-exten"`
}

type AsteriskConfig struct {
	IsEnabled bool            `json:"enabled" yaml:"enabled"`
	Port      int             `json:"port" binding:"required" yaml:"port"`
	Host      string          `json:"host" binding:"required" yaml:"host"`
	Username  string          `json:"username" binding:"required" yaml:"username"`
	Password  string          `json:"password" yaml:"password"`
	Telephony TelephonyConfig `json:"telephony" yaml:"telephony"`
}
