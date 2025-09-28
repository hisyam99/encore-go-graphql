package middleware

import (
	"context"
	"errors"
	"strings"

	"encore.app/app"
	"encore.app/app/auth"
)

// AuthContext key for storing user in context
type contextKey string

const UserContextKey contextKey = "user"

// GetUserFromContext extracts user from context
func GetUserFromContext(ctx context.Context) (*auth.Claims, bool) {
	user, ok := ctx.Value(UserContextKey).(*auth.Claims)
	return user, ok
}

// RequireAuth middleware that requires valid JWT token
func RequireAuth(ctx context.Context, token string) (*auth.Claims, error) {
	if strings.TrimSpace(token) == "" {
		return nil, errors.New("authorization token required")
	}

	// Remove "Bearer " prefix if present
	if after, ok := strings.CutPrefix(token, "Bearer "); ok {
		token = after
	}

	// Validate token
	claims, err := auth.ValidateToken(token)
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}

// RequireRole middleware that requires specific role
func RequireRole(claims *auth.Claims, requiredRole app.UserRole) error {
	if claims == nil {
		return errors.New("authentication required")
	}

	user := &app.User{Role: claims.Role}
	if !user.HasPermission(requiredRole) {
		return errors.New("insufficient permissions")
	}

	return nil
}

// RequireOwnershipOrAdmin checks if user owns resource or is admin
func RequireOwnershipOrAdmin(claims *auth.Claims, resourceUserID uint) error {
	if claims == nil {
		return errors.New("authentication required")
	}

	// Allow if user is admin
	if claims.Role == app.UserRoleAdmin {
		return nil
	}

	// Allow if user owns the resource
	if claims.UserID == resourceUserID {
		return nil
	}

	return errors.New("access denied: insufficient permissions")
}

// RequireOwnershipOrRole checks if user owns resource or has required role
func RequireOwnershipOrRole(claims *auth.Claims, resourceUserID uint, requiredRole app.UserRole) error {
	if claims == nil {
		return errors.New("authentication required")
	}

	// Allow if user owns the resource
	if claims.UserID == resourceUserID {
		return nil
	}

	// Allow if user has required role or higher
	user := &app.User{Role: claims.Role}
	if user.HasPermission(requiredRole) {
		return nil
	}

	return errors.New("access denied: insufficient permissions")
}

// RequireActiveUser checks if user account is active
func RequireActiveUser(claims *auth.Claims) error {
	if claims == nil {
		return errors.New("authentication required")
	}

	// In production, you might want to check the database for user status
	// For now, we assume if token is valid, user was active when token was created
	return nil
}

// RequireNotGuest ensures user is not a guest
func RequireNotGuest(claims *auth.Claims) error {
	if claims == nil {
		return errors.New("authentication required")
	}

	if claims.Role == app.UserRoleGuest {
		return errors.New("guest users are not allowed for this operation")
	}

	return nil
}
