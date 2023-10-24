package corsx

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/sivaosorg/govm/common"
	"github.com/sivaosorg/govm/restify"
	"github.com/sivaosorg/govm/utils"
)

func NewCorsConfig() *CorsConfig {
	c := &CorsConfig{}
	c.SetMaxAge(0)
	return c
}

func (c *CorsConfig) SetEnabled(value bool) *CorsConfig {
	c.IsEnabled = value
	return c
}

func (c *CorsConfig) SetAllowCredentials(value bool) *CorsConfig {
	c.AllowCredentials = value
	return c
}

func (c *CorsConfig) SetMaxAge(value int) *CorsConfig {
	if value < 0 {
		log.Panicf("Invalid max-age: %v", value)
	}
	c.MaxAge = value
	return c
}

func (c *CorsConfig) SetAllowedOrigins(values []string) *CorsConfig {
	c.AllowedOrigins = values
	return c
}

func (c *CorsConfig) AppendAllowedOrigins(values ...string) *CorsConfig {
	c.AllowedOrigins = append(c.AllowedOrigins, values...)
	return c
}

func (c *CorsConfig) SetAllowedHeaders(values []string) *CorsConfig {
	c.AllowedHeaders = values
	return c
}

func (c *CorsConfig) AppendAllowedHeaders(values ...string) *CorsConfig {
	c.AllowedHeaders = append(c.AllowedHeaders, values...)
	return c
}

func (c *CorsConfig) SetAllowedMethods(values []string) *CorsConfig {
	c.AllowedMethods = values
	return c
}

func (c *CorsConfig) AppendAllowedMethods(values ...string) *CorsConfig {
	c.AllowedMethods = append(c.AllowedMethods, values...)
	return c
}

func (c *CorsConfig) SetExposedHeaders(values []string) *CorsConfig {
	c.ExposedHeaders = values
	return c
}

func (c *CorsConfig) AppendExposedHeaders(values ...string) *CorsConfig {
	c.ExposedHeaders = append(c.ExposedHeaders, values...)
	return c
}

func (c *CorsConfig) Json() string {
	return utils.ToJson(c)
}

func GetCorsConfigSample() *CorsConfig {
	c := NewCorsConfig().
		SetEnabled(true).
		SetMaxAge(3600).
		SetAllowCredentials(true).
		AppendAllowedOrigins("*").
		AppendAllowedMethods(restify.MethodGet, restify.MethodPost, restify.MethodPut, restify.MethodDelete, restify.MethodOptions).
		SetExposedHeaders(make([]string, 0)).
		AppendAllowedHeaders(common.HeaderOrigin, common.HeaderAccept, common.HeaderContentType, common.HeaderAuthorization)
	return c
}

// ApplyCorsHeaders applies CORS headers to an HTTP response based on the configuration.
func (c *CorsConfig) ApplyCorsHeaders(w http.ResponseWriter, r *http.Request) {
	if !c.IsEnabled {
		return
	}
	origin := r.Header.Get(common.HeaderOrigin)
	if c.isOriginAllowed(origin) {
		w.Header().Set(common.HeaderAccessControlAllowOrigin, origin)
		w.Header().Set(common.HeaderAccessControlAllowMethods, strings.Join(c.AllowedMethods, ","))
		w.Header().Set(common.HeaderAccessControlAllowHeaders, strings.Join(c.AllowedHeaders, ","))
		w.Header().Set(common.HeaderAccessControlExposeHeaders, strings.Join(c.ExposedHeaders, ","))
		if r.Method == http.MethodOptions { // handle preflight request
			w.WriteHeader(http.StatusOK)
			return
		}
		if c.AllowCredentials {
			w.Header().Set(common.HeaderAccessControlAllowCredentials, strconv.FormatBool(c.AllowCredentials))
		}
		if c.MaxAge > 0 {
			w.Header().Set(common.HeaderAccessControlMaxAge, strconv.Itoa(c.MaxAge))
		}
	}
}

func (c *CorsConfig) CoreMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.ApplyCorsHeaders(w, r)
		next.ServeHTTP(w, r)
	})
}

// isOriginAllowed checks if the provided origin is allowed based on the configuration.
func (c *CorsConfig) isOriginAllowed(origin string) bool {
	for _, allowedOrigin := range c.AllowedOrigins {
		if allowedOrigin == "*" || allowedOrigin == origin {
			return true
		}
	}
	return false
}
