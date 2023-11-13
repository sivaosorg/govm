package document

import (
	"fmt"
	"time"

	"github.com/sivaosorg/govm/utils"
)

func NewGoogleSheetConfig() *GoogleSheetConfig {
	return &GoogleSheetConfig{
		Timeout: 10 * time.Second,
	}
}

func (g *GoogleSheetConfig) SetEnabled(value bool) *GoogleSheetConfig {
	g.IsEnabled = value
	return g
}

func (g *GoogleSheetConfig) SetKey(value string) *GoogleSheetConfig {
	g.Key = value
	return g
}

func (g *GoogleSheetConfig) SetSpreadSheetCredential(value string) *GoogleSheetConfig {
	g.SpreadSheetCredential = value
	return g
}

func (g *GoogleSheetConfig) SetSpreadSheetId(value string) *GoogleSheetConfig {
	g.SpreadSheetId = value
	return g
}

func (g *GoogleSheetConfig) SetHeaderRange(value string) *GoogleSheetConfig {
	g.HeaderRange = value
	return g
}

func (g *GoogleSheetConfig) SetTimeout(value time.Duration) *GoogleSheetConfig {
	g.Timeout = value
	return g
}

func (g *GoogleSheetConfig) Json() string {
	return utils.ToJson(g)
}

func GetGoogleSheetSample() *GoogleSheetConfig {
	g := NewGoogleSheetConfig().
		SetEnabled(false).
		SetTimeout(10 * time.Second).
		SetKey("google_spread_sheet").
		SetSpreadSheetCredential("./keys/google_sheet_credentials.json").
		SetSpreadSheetId("sheet_id").
		SetHeaderRange("Sheet1!A:C")
	return g
}

func NewGoogleSheetOptionConfig() *googleSheetOptionConfig {
	return &googleSheetOptionConfig{}
}

func NewMultiTenantGoogleSheetConfig() *MultiTenantGoogleSheetConfig {
	return &MultiTenantGoogleSheetConfig{}
}

func (m *MultiTenantGoogleSheetConfig) SetKey(value string) *MultiTenantGoogleSheetConfig {
	m.Key = value
	return m
}

func (m *MultiTenantGoogleSheetConfig) SetUsableDefault(value bool) *MultiTenantGoogleSheetConfig {
	m.IsUsableDefault = value
	return m
}

func (m *MultiTenantGoogleSheetConfig) SetConfig(value GoogleSheetConfig) *MultiTenantGoogleSheetConfig {
	m.Config = value
	return m
}

func (m *MultiTenantGoogleSheetConfig) SetConfigCursor(value *GoogleSheetConfig) *MultiTenantGoogleSheetConfig {
	m.Config = *value
	return m
}

func (m *MultiTenantGoogleSheetConfig) SetOption(value googleSheetOptionConfig) *MultiTenantGoogleSheetConfig {
	m.Option = value
	return m
}

func (m *MultiTenantGoogleSheetConfig) Json() string {
	return utils.ToJson(m)
}

func GetMultiTenantGoogleSheetConfigSample() *MultiTenantGoogleSheetConfig {
	m := NewMultiTenantGoogleSheetConfig().
		SetUsableDefault(false).
		SetKey("tenant_1").
		SetOption(*NewGoogleSheetOptionConfig()).
		SetConfigCursor(GetGoogleSheetSample())
	return m
}

func NewClusterMultiTenantGoogleSheetConfig() *ClusterMultiTenantGoogleSheetConfig {
	return &ClusterMultiTenantGoogleSheetConfig{}
}

func (c *ClusterMultiTenantGoogleSheetConfig) SetClusters(values []MultiTenantGoogleSheetConfig) *ClusterMultiTenantGoogleSheetConfig {
	c.Clusters = values
	return c
}

func (c *ClusterMultiTenantGoogleSheetConfig) AppendClusters(values ...MultiTenantGoogleSheetConfig) *ClusterMultiTenantGoogleSheetConfig {
	c.Clusters = append(c.Clusters, values...)
	return c
}

func (c *ClusterMultiTenantGoogleSheetConfig) Json() string {
	return utils.ToJson(c.Clusters)
}

func GetClusterMultiTenantGoogleSheetConfigSample() *ClusterMultiTenantGoogleSheetConfig {
	c := NewClusterMultiTenantGoogleSheetConfig().
		AppendClusters(*GetMultiTenantGoogleSheetConfigSample(), *GetMultiTenantGoogleSheetConfigSample().SetKey("tenant_2"))
	return c
}

func (c *ClusterMultiTenantGoogleSheetConfig) FindClusterBy(key string) (MultiTenantGoogleSheetConfig, error) {
	if utils.IsEmpty(key) {
		return *NewMultiTenantGoogleSheetConfig(), fmt.Errorf("Key is required")
	}
	if len(c.Clusters) == 0 {
		return *NewMultiTenantGoogleSheetConfig(), fmt.Errorf("No google sheet cluster")
	}
	for _, v := range c.Clusters {
		if v.Key == key {
			return v, nil
		}
	}
	return *NewMultiTenantGoogleSheetConfig(), fmt.Errorf("The google sheet cluster not found")
}
