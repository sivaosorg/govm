package bjson

import "unsafe"

type Type int

// ResultContext represents a json value that is returned from Get().
type ResultContext struct {
	Type    Type    `json:"-"`
	Raw     string  `json:"-"`
	Strings string  `json:"-"`
	Numeric float64 `json:"-"`
	Index   int     `json:"-"`
	Indexes []int   `json:"-"`
}

type AomContext struct {
	ArrayResult       []ResultContext          `json:"-"`
	ArrayInterface    []interface{}            `json:"-"`
	OptionalMap       map[string]ResultContext `json:"-"`
	OptionalInterface map[string]interface{}   `json:"-"`
	valueX            byte                     `json:"-"`
}

type PathContext struct {
	Part  string `json:"-"`
	Path  string `json:"-"`
	Pipe  string `json:"-"`
	Piped bool   `json:"-"`
	Wild  bool   `json:"-"`
	More  bool   `json:"-"`
}

type DeepContext struct {
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

type ParseContext struct {
	json  string        `json:"-"`
	value ResultContext `json:"-"`
	pipe  string        `json:"-"`
	piped bool          `json:"-"`
	calc  bool          `json:"-"`
	lines bool          `json:"-"`
}

// StringHeader instead of reflect.StringHeader
type StringHeader struct {
	data   unsafe.Pointer `json:"-"`
	length int            `json:"-"`
}

// SliceHeader instead of reflect.SliceHeader
type SliceHeader struct {
	data     unsafe.Pointer `json:"-"`
	length   int            `json:"-"`
	capacity int            `json:"-"`
}
