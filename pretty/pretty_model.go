package pretty

type OptionsConfig struct {
	// Width is an max column width for single line arrays
	// Default is 80
	Width int `json:"width"`
	// Prefix is a prefix for all lines
	// Default is an empty string
	Prefix string `json:"prefix"`
	// Indent is the nested indentation
	// Default is two spaces
	Indent string `json:"indent"`
	// SortKeys will sort the keys alphabetically
	// Default is false
	SortKeys bool `json:"sort_keys"`
}

// DefaultOptionsConfig is the default options for pretty formats.
var DefaultOptionsConfig = &OptionsConfig{Width: 80, Prefix: "", Indent: "  ", SortKeys: false}
