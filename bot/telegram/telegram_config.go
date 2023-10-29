package telegram

const (
	Host string = "https://api.telegram.org"
)

const (
	ModeMarkdown   TelegramFormatType = "Markdown"
	ModeMarkdownV2 TelegramFormatType = "MarkdownV2"
	ModeHTML       TelegramFormatType = "HTML"
	ModeNone       TelegramFormatType = "None"
)

var (
	ModeText map[TelegramFormatType]bool = map[TelegramFormatType]bool{
		ModeHTML:     true,
		ModeMarkdown: true,
		ModeNone:     true,
	}
)
