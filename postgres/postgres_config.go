package postgres

import "fmt"

const (
	PostgresSslDisabledMode = "disable"
	PostgresSslEnabledMode  = "enable"
)

var (
	PostgresSslModes map[string]bool = map[string]bool{
		PostgresSslDisabledMode: true,
		PostgresSslEnabledMode:  true,
	}
)

var (
	PostgresSslModeError = fmt.Sprintf("Invalid ssl-mode, only supported values: %v, %v", PostgresSslDisabledMode, PostgresSslEnabledMode)
)
