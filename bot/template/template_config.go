package template

const (
	TypeNotification = "Notification"
	TypeInformation  = "Info"
	TypeWarning      = "Warning"
	TypeError        = "Error"
	TypeDebug        = "Debug"
	TypeSuccess      = "Success"
	TypeBug          = "Bug"
	TypeTrace        = "Trace"
)

var TypeIcons = map[string]string{
	TypeNotification: "🔔",
	TypeInformation:  "ℹ️",
	TypeWarning:      "⚠️",
	TypeError:        "❌",
	TypeDebug:        "🐞",
	TypeSuccess:      "✅",
	TypeBug:          "🐛",
	TypeTrace:        "🔍",
}
