package tags

import (
	"errors"
	"reflect"
)

const (
	// TagName is used to mention field options.
	//
	// Example:
	// --------
	// Age	int		`defined:"age"`
	// Info	StoreInfo	`defined:"info,no_traverse"`
	TagName    = "defined"
	OmitField  = "-"
	OmitEmpty  = "omitempty"
	NoTraverse = "no_traverse"
)

var (
	NoTraverseTypes     map[reflect.Type]bool
	ReflectConverters   map[reflect.Type]map[reflect.Type]TagConverter
	TypeOfBytes         = reflect.TypeOf([]byte(nil))
	TypeOfInterface     = reflect.TypeOf((*interface{})(nil)).Elem()
	ErrorFieldNotExists = errors.New("Field does not exists")
)
