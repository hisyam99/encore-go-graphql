package utils

import (
	"context"
	"errors"
	"log"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

// Common error types
var (
	ErrNotFound       = errors.New("resource not found")
	ErrAlreadyExists  = errors.New("resource already exists")
	ErrValidation     = errors.New("validation error")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrInternalServer = errors.New("internal server error")
)

// GraphQLError represents a GraphQL error with extensions
type GraphQLError struct {
	Message    string                 `json:"message"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

// Error implements the error interface
func (e *GraphQLError) Error() string {
	return e.Message
}

// NewValidationError creates a new validation error
func NewValidationError(message string, field string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: message,
		Extensions: map[string]interface{}{
			"code":  "VALIDATION_ERROR",
			"field": field,
		},
	}
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(resource string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: resource + " not found",
		Extensions: map[string]interface{}{
			"code":     "NOT_FOUND",
			"resource": resource,
		},
	}
}

// NewAlreadyExistsError creates a new already exists error
func NewAlreadyExistsError(resource string, field string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: resource + " already exists",
		Extensions: map[string]interface{}{
			"code":     "ALREADY_EXISTS",
			"resource": resource,
			"field":    field,
		},
	}
}

// NewInternalError creates a new internal server error
func NewInternalError(message string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: "Internal server error",
		Extensions: map[string]interface{}{
			"code":    "INTERNAL_ERROR",
			"details": message,
		},
	}
}

// HandleRepositoryError converts repository errors to GraphQL errors
func HandleRepositoryError(err error, resource string) *gqlerror.Error {
	if err == nil {
		return nil
	}

	// Handle GORM specific errors
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NewNotFoundError(resource)
	}

	// Handle duplicate key errors (PostgreSQL)
	if isDuplicateKeyError(err) {
		return NewAlreadyExistsError(resource, "unique field")
	}

	// Log the actual error for debugging
	log.Printf("Repository error for %s: %v", resource, err)

	// Return generic internal error to client
	return NewInternalError("Failed to process request")
}

// isDuplicateKeyError checks if the error is a duplicate key constraint violation
func isDuplicateKeyError(err error) bool {
	errStr := err.Error()
	return contains(errStr, "duplicate key") ||
		contains(errStr, "UNIQUE constraint") ||
		contains(errStr, "uniqueIndex")
}

// contains checks if a string contains a substring (case-insensitive)
func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr ||
			(len(s) > len(substr) &&
				(s[:len(substr)] == substr ||
					s[len(s)-len(substr):] == substr ||
					indexIgnoreCase(s, substr) >= 0)))
}

// indexIgnoreCase finds the index of substr in s (case-insensitive)
func indexIgnoreCase(s, substr string) int {
	sLower := toLower(s)
	substrLower := toLower(substr)
	for i := 0; i <= len(sLower)-len(substrLower); i++ {
		if sLower[i:i+len(substrLower)] == substrLower {
			return i
		}
	}
	return -1
}

// toLower converts string to lowercase
func toLower(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			result[i] = c + 32
		} else {
			result[i] = c
		}
	}
	return string(result)
}

// LogError logs an error with context
func LogError(ctx context.Context, operation string, err error) {
	log.Printf("[ERROR] %s: %v", operation, err)
}

// LogInfo logs an info message with context
func LogInfo(ctx context.Context, operation string, message string) {
	log.Printf("[INFO] %s: %s", operation, message)
}

// RecoverFromPanic recovers from panic and returns a GraphQL error
func RecoverFromPanic() *gqlerror.Error {
	if r := recover(); r != nil {
		log.Printf("[PANIC] Recovered: %v", r)
		return NewInternalError("Unexpected error occurred")
	}
	return nil
}
