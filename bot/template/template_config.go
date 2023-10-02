package template

const (
	MessageTypeNotification = "Notification"
	MessageTypeInformation  = "Information"
	MessageTypeWarning      = "Warning"
	MessageTypeError        = "Error"
	MessageTypeDebug        = "Debug"
	MessageTypeSuccess      = "Success"
	MessageTypeBug          = "Bug"
	MessageTypeTrace        = "Trace"
)

var MessageTypeIcons = map[string]string{
	MessageTypeNotification: "🔔",
	MessageTypeInformation:  "ℹ️",
	MessageTypeWarning:      "⚠️",
	MessageTypeError:        "❌",
	MessageTypeDebug:        "🐞",
	MessageTypeSuccess:      "✅",
	MessageTypeBug:          "🐛",
	MessageTypeTrace:        "🔍",
}
