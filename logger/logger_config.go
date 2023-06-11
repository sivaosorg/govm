package logger

const (
	LoggerMessageField = "message"
	LoggerErrorField   = "error"
	LoggerSuccessField = "success"
)

const (
	LoggerJsonFormatter = "json"
	LoggerTextFormatter = "text"
)

var (
	LoggerFormatters map[string]bool = map[string]bool{
		LoggerJsonFormatter: true,
		LoggerTextFormatter: true,
	}
)
