package bjson

import "unsafe"

type Type int

// BJsonContext represents a json value that is returned from Get().
type BJsonContext struct {
	Type    Type    `json:"-"`
	Raw     string  `json:"raw"`
	Strings string  `json:"-"`
	Numeric float64 `json:"-"`
	Index   int     `json:"index"`
	Indexes []int   `json:"-"`
}

type aomContext struct {
	ArrayResult       []BJsonContext          `json:"-"`
	ArrayInterface    []interface{}           `json:"-"`
	OptionalMap       map[string]BJsonContext `json:"-"`
	OptionalInterface map[string]interface{}  `json:"-"`
	valueX            byte                    `json:"-"`
}

type pathContext struct {
	Part  string `json:"-"`
	Path  string `json:"-"`
	Pipe  string `json:"-"`
	Piped bool   `json:"-"`
	Wild  bool   `json:"-"`
	More  bool   `json:"-"`
}

type deepContext struct {
	Part    string `json:"-"`
	Path    string `json:"-"`
	Pipe    string `json:"-"`
	Piped   bool   `json:"-"`
	More    bool   `json:"-"`
	Arch    bool   `json:"-"`
	ALogOk  bool   `json:"-"`
	ALogKey string `json:"-"`
	query   struct {
		On        bool   `json:"-"`
		All       bool   `json:"-"`
		QueryPath string `json:"-"`
		Option    string `json:"-"`
		Value     string `json:"-"`
	} `json:"-"`
}

type parseContext struct {
	json  string       `json:"-"`
	value BJsonContext `json:"-"`
	pipe  string       `json:"-"`
	piped bool         `json:"-"`
	calc  bool         `json:"-"`
	lines bool         `json:"-"`
}

// stringHeader instead of reflect.stringHeader
type stringHeader struct {
	data   unsafe.Pointer `json:"-"`
	length int            `json:"-"`
}

// sliceHeader instead of reflect.sliceHeader
type sliceHeader struct {
	data     unsafe.Pointer `json:"-"`
	length   int            `json:"-"`
	capacity int            `json:"-"`
}
