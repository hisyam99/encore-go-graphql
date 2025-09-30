package utils

import (
	"net/http"
	"strings"

	"encore.dev"
)

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
	MaxAge           int
}

// GetCORSConfig returns CORS configuration based on environment
func GetCORSConfig() CORSConfig {
	if encore.Meta().Environment.Type == encore.EnvProduction {
		return CORSConfig{
			AllowedOrigins: []string{
				"https://hisyam.tar.my.id",
				"https://www.hisyam.tar.my.id",
			},
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
				http.MethodOptions,
			},
			AllowedHeaders: []string{
				"Accept",
				"Authorization",
				"Content-Type",
				"X-CSRF-Token",
				"X-Requested-With",
			},
			AllowCredentials: true,
			MaxAge:           3600,
		}
	}

	// Development configuration - more permissive
	return CORSConfig{
		AllowedOrigins: []string{
			"http://localhost:4003",
			"http://localhost:4003",
			"https://localhost:4003",
			"https://localhost:4003",
			"https://hisyam.tar.my.id",
			"http://hisyam.tar.my.id",
			"http://127.0.0.1:4000",
			"http://127.0.0.1:4003",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"X-Requested-With",
		},
		AllowCredentials: true,
		MaxAge:           3600,
	}
}

// isOriginAllowed checks if the origin is in the allowed list
func (c CORSConfig) isOriginAllowed(origin string) bool {
	for _, allowedOrigin := range c.AllowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}
	return false
}

// CORSMiddleware creates CORS middleware with the given configuration
func CORSMiddleware(config CORSConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// Set CORS headers
			if config.isOriginAllowed(origin) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

			w.Header().Set("Access-Control-Allow-Methods", strings.Join(config.AllowedMethods, ", "))
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(config.AllowedHeaders, ", "))

			if config.AllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			if config.MaxAge > 0 {
				w.Header().Set("Access-Control-Max-Age", "3600")
			}

			// Handle preflight requests
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
