package builder

type MapBuilder struct {
	Result map[string]interface{} `json:"_map,omitempty"`
}

type KeyValuePair struct {
	Key   string      `json:"key,omitempty"`
	Value interface{} `json:"value,omitempty"`
}
