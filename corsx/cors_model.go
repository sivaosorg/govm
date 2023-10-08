package corsx

// CorsConfig represents the CORS (Cross-Origin Resource Sharing) configuration.
type CorsConfig struct {
	IsEnabled        bool     `json:"enabled" yaml:"enabled"`
	AllowedOrigins   []string `json:"allowed_origins" yaml:"allowed-origins"`
	AllowedMethods   []string `json:"allowed_methods" yaml:"allowed-methods"`
	AllowedHeaders   []string `json:"allowed_headers" yaml:"allowed-headers"`
	ExposedHeaders   []string `json:"exposed_headers" yaml:"exposed-headers"`
	AllowCredentials bool     `json:"allow_credentials" yaml:"allow-credentials"`
	MaxAge           int      `json:"max_age" binding:"gte=0" yaml:"max-age"`
}
