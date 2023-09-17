package tags

import (
	"reflect"
)

type TagConverter func(in reflect.Value) (reflect.Value, error)

type TagConfig struct {
	Name    string `json:"name"`
	Options string `json:"options"`
}
