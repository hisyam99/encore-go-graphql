package graphql

import (
	"context"

	"encore.app/app"
	"encore.app/app/auth"
	"encore.app/app/middleware"
	"encore.app/app/utils"
)

// AuthResult represents the result of authentication
type AuthResult struct {
	Claims *auth.Claims
	UserID uint
}

// RequireAuthWithRole authenticates user and checks for required role
func RequireAuthWithRole(ctx context.Context, requiredRole app.UserRole) (*AuthResult, error) {
	// Get auth token from headers
	authHeader := utils.GetAuthHeaderFromContext(ctx)
	claims, err := middleware.RequireAuth(ctx, authHeader)
	if err != nil {
		return nil, err
	}

	// Check role requirement
	if err := middleware.RequireRole(claims, requiredRole); err != nil {
		return nil, err
	}

	return &AuthResult{
		Claims: claims,
		UserID: claims.UserID,
	}, nil
}

// RequireAuth authenticates user without role check
func RequireAuth(ctx context.Context) (*AuthResult, error) {
	// Get auth token from headers
	authHeader := utils.GetAuthHeaderFromContext(ctx)
	claims, err := middleware.RequireAuth(ctx, authHeader)
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		Claims: claims,
		UserID: claims.UserID,
	}, nil
}

// ExtractOptionalString safely extracts string pointer value
func ExtractOptionalString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

// ExtractOptionalStringSlice safely extracts string slice from pointer
func ExtractOptionalStringSlice(ptr []string) []string {
	if ptr == nil {
		return []string{}
	}
	return ptr
}
