package curlx

type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	OPTIONS Method = "OPTIONS"
)

const (
	FieldNameFileForm = "file"
)

var (
	MethodWithRequestBody map[Method]bool = map[Method]bool{
		POST:    true,
		PUT:     true,
		PATCH:   true,
		DELETE:  true,
		OPTIONS: true,
	}
)
