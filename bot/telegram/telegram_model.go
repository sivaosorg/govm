package telegram

type TelegramFormatType string

type TelegramConfig struct {
	IsEnabled bool    `json:"enabled" yaml:"enabled"`
	DebugMode bool    `json:"debug_mode" yaml:"debug-mode"`
	ChatID    []int64 `json:"chat_id" binding:"required" yaml:"chat_id"`
	Token     string  `json:"token" binding:"required" yaml:"token"`
}

type TelegramOptionConfig struct {
	Type TelegramFormatType `json:"type" binding:"required" yaml:"type"`
}

type MultiTenantTelegramConfig struct {
	Key             string               `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool                 `json:"usable_default" yaml:"usable_default"`
	Config          TelegramConfig       `json:"config" yaml:"config"`
	Option          TelegramOptionConfig `json:"option" binding:"required" yaml:"option"`
}

type ClusterMultiTenantTelegramConfig struct {
	Clusters []MultiTenantTelegramConfig `json:"clusters,omitempty" yaml:"clusters"`
}
