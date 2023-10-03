package blueprint

const (
	TypeNotification IconText = "Notification"
	TypeInfo         IconText = "Info"
	TypeWarning      IconText = "Warning"
	TypeError        IconText = "Error"
	TypeDebug        IconText = "Debug"
	TypeSuccess      IconText = "Success"
	TypeBug          IconText = "Bug"
	TypeTrace        IconText = "Trace"
)

var TypeIcons = map[IconText]string{
	TypeNotification: "🔔",
	TypeInfo:         "ℹ️",
	TypeWarning:      "⚠️",
	TypeError:        "❌",
	TypeDebug:        "🐞",
	TypeSuccess:      "✅",
	TypeBug:          "🐛",
	TypeTrace:        "🔍",
}

// CardDefault is the HTML template for the card
var CardDefault = `
<b> {{.Icon}} {{.Title}} </b>
<pre>{{.Description}}</pre>
{{if .ImageUrl}}
<img src="{{.ImageUrl}}">
{{end}}
{{if .ButtonUrl}}
<a href="{{.ButtonUrl}}">{{.ButtonText}}</a>
{{end}}
`
