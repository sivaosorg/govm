package logger

const (
	LoggerMessageField = "message"
	LoggerErrorField   = "error"
	LoggerSuccessField = "success"
	LoggerFileField    = "@file"
	LoggerLineField    = "@line"
	LoggerCallerField  = "@caller"
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
