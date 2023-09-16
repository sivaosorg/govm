package bot

const (
	DiscordBot  TypeBot = "discord_bot"
	TelegramBot TypeBot = "telegram_bot"
)

var (
	Bot map[TypeBot]bool = map[TypeBot]bool{
		DiscordBot:  true,
		TelegramBot: true,
	}
)
