package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"encore.app/app/repositories"
)

// ValidateEmail validates email format
func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(strings.ToLower(email)) {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidateSlug validates blog slug format
func ValidateSlug(slug string) error {
	if slug == "" {
		return errors.New("slug cannot be empty")
	}
	slugRegex := regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	if !slugRegex.MatchString(slug) {
		return errors.New("slug must contain only lowercase letters, numbers, and hyphens")
	}
	return nil
}

// GenerateSlug generates a URL-friendly slug from a title
func GenerateSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace spaces and special characters with hyphens
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, "-")

	// Remove leading and trailing hyphens
	slug = strings.Trim(slug, "-")

	// Add timestamp suffix to ensure uniqueness
	timestamp := time.Now().Unix()
	slug = fmt.Sprintf("%s-%d", slug, timestamp)

	return slug
}

// ValidatePaginationParams validates and sets defaults for pagination parameters
func ValidatePaginationParams(page, pageSize int, sortBy string) repositories.PaginationParams {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// Validate sortBy field (basic validation)
	allowedSortFields := map[string]bool{
		"id":          true,
		"name":        true,
		"title":       true,
		"email":       true,
		"slug":        true,
		"author":      true,
		"status":      true,
		"createdAt":   true,
		"updatedAt":   true,
		"publishedAt": true,
	}

	if sortBy == "" || !allowedSortFields[sortBy] {
		sortBy = "createdAt"
	}

	return repositories.PaginationParams{
		Page:     page,
		PageSize: pageSize,
		SortBy:   sortBy,
		SortDesc: true, // Default to descending
	}
}

// ParseID parses string ID to uint
func ParseID(id string) (uint, error) {
	parsed, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format: %w", err)
	}
	return uint(parsed), nil
}

// FormatID formats uint ID to string
func FormatID(id uint) string {
	return strconv.FormatUint(uint64(id), 10)
}

// FormatTime formats time to RFC3339 string
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// FormatTimePtr formats time pointer to RFC3339 string
func FormatTimePtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formatted := t.Format(time.RFC3339)
	return &formatted
}

// ValidateRequired validates that a string field is not empty
func ValidateRequired(value, fieldName string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

// ValidateMaxLength validates that a string field doesn't exceed max length
func ValidateMaxLength(value, fieldName string, maxLen int) error {
	if len(value) > maxLen {
		return fmt.Errorf("%s cannot exceed %d characters", fieldName, maxLen)
	}
	return nil
}

// ValidateMinLength validates that a string field meets minimum length
func ValidateMinLength(value, fieldName string, minLen int) error {
	if len(strings.TrimSpace(value)) < minLen {
		return fmt.Errorf("%s must be at least %d characters", fieldName, minLen)
	}
	return nil
}

// SanitizeString removes leading/trailing whitespace and normalizes spaces
func SanitizeString(value string) string {
	// Remove leading and trailing whitespace
	value = strings.TrimSpace(value)

	// Replace multiple spaces with single space
	spaceRegex := regexp.MustCompile(`\s+`)
	value = spaceRegex.ReplaceAllString(value, " ")

	return value
}

// ToPointerSlice converts []T to []*T
func ToPointerSlice[T any](slice []T) []*T {
	result := make([]*T, len(slice))
	for i := range slice {
		result[i] = &slice[i]
	}
	return result
}
