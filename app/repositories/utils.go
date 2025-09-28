package repositories

import "strings"

// convertToSnakeCase converts camelCase to snake_case for database columns
func convertToSnakeCase(str string) string {
	// Common field mappings for GraphQL to database columns
	switch str {
	case "createdAt":
		return "created_at"
	case "updatedAt":
		return "updated_at"
	case "deletedAt":
		return "deleted_at"
	case "categoryId":
		return "category_id"
	case "userId":
		return "user_id"
	case "publishedAt":
		return "published_at"
	case "metaDescription":
		return "meta_description"
	default:
		// For other fields, convert camelCase to snake_case
		var result strings.Builder
		for i, r := range str {
			if i > 0 && r >= 'A' && r <= 'Z' {
				result.WriteByte('_')
			}
			result.WriteRune(r | 0x20) // Convert to lowercase
		}
		return result.String()
	}
}
