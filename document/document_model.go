package document

import "time"

type GoogleSheetConfig struct {
	IsEnabled             bool          `json:"enabled" yaml:"enabled"`
	Key                   string        `json:"key" yaml:"key"`
	SpreadSheetCredential string        `json:"spread_sheet_credential" yaml:"spread_sheet_credential"`
	SpreadSheetId         string        `json:"spread_sheet_id" yaml:"spread_sheet_id"`
	HeaderRange           string        `json:"header_range" yaml:"header_ranges"`
	Timeout               time.Duration `json:"timeout" yaml:"timeout"`
}

type googleSheetOptionConfig struct {
}

type MultiTenantGoogleSheetConfig struct {
	Key             string                  `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool                    `json:"usable_default" yaml:"usable_default"`
	Config          GoogleSheetConfig       `json:"config" yaml:"config"`
	Option          googleSheetOptionConfig `json:"option,omitempty" yaml:"option"`
}

type ClusterMultiTenantGoogleSheetConfig struct {
	Clusters []MultiTenantGoogleSheetConfig `json:"clusters,omitempty" yaml:"clusters"`
}
