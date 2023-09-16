package telegram

const (
	Host string = "https://api.telegram.org"
)

const (
	ModeMarkdown TelegramFormatType = "Markdown"
	ModeHTML     TelegramFormatType = "HTML"
	ModeNone     TelegramFormatType = "None"
)

var (
	ModeText map[TelegramFormatType]bool = map[TelegramFormatType]bool{
		ModeHTML:     true,
		ModeMarkdown: true,
		ModeNone:     true,
	}
)
