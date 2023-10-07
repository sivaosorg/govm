package builder

type MapBuilder struct {
	result map[string]interface{} `json:"-"`
}

type KeyValuePair struct {
	Key   string      `json:"key,omitempty"`
	Value interface{} `json:"value,omitempty"`
}
