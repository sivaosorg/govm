package bot

import "github.com/sivaosorg/govm/bot/telegram"

type TypeBot string

type BotConfig struct {
	IsEnabled bool                                      `json:"enabled" yaml:"enabled"`
	DebugMode bool                                      `json:"debug_mode" yaml:"debug_mode"`
	Telegram  telegram.ClusterMultiTenantTelegramConfig `json:"telegram_config,omitempty" yaml:"telegram_config"`
}
