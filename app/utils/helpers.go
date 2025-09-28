package utils

import (
	"context"
	"strconv"
	"time"
)

// GetAuthHeaderFromContext extracts authorization header from GraphQL context
func GetAuthHeaderFromContext(ctx context.Context) string {
	// Try to get from direct context value
	if headers, ok := ctx.Value("headers").(map[string]string); ok {
		if auth := headers["authorization"]; auth != "" {
			return auth
		}
		if auth := headers["Authorization"]; auth != "" {
			return auth
		}
	}

	// Try to get from HTTP headers if available
	if headers, ok := ctx.Value("http-headers").(map[string][]string); ok {
		if authHeaders := headers["Authorization"]; len(authHeaders) > 0 {
			return authHeaders[0]
		}
		if authHeaders := headers["authorization"]; len(authHeaders) > 0 {
			return authHeaders[0]
		}
	}

	return ""
}

// TimeToStringPtr converts time pointer to string pointer
func TimeToStringPtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	str := t.Format(time.RFC3339)
	return &str
}

// UintToStringPtr converts uint pointer to string pointer
func UintToStringPtr(u *uint) *string {
	if u == nil {
		return nil
	}
	str := strconv.Itoa(int(*u))
	return &str
}

// StringToUintPtr converts string pointer to uint pointer
func StringToUintPtr(s *string) *uint {
	if s == nil {
		return nil
	}
	u, err := strconv.ParseUint(*s, 10, 32)
	if err != nil {
		return nil
	}
	result := uint(u)
	return &result
}

// UintToString converts uint to string
func UintToString(u uint) string {
	return strconv.Itoa(int(u))
}

// StringToUint converts string to uint
func StringToUint(s string) (uint, error) {
	u, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(u), nil
}