package auth

import (
	"errors"
	"os"
	"strconv"
	"time"

	"encore.app/app"
	"github.com/golang-jwt/jwt/v5"
)

// jwtSecret retrieves the JWT secret from environment variables.
// In Encore, use secrets for production: encore.secret.set("JWT_SECRET", "your-secret-key")
// This ensures security and avoids hardcoding.
var jwtSecret = []byte(getJWTSecret())

// getJWTSecret returns the JWT secret, defaulting to a development key if not set.
// Encore automatically handles secrets in production environments.
func getJWTSecret() string {
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		return secret
	}
	// Fallback for development; change this in production.
	return "your-super-secret-jwt-key-change-in-production-2024-portfolio-api"
}

// Claims represents JWT claims for user authentication.
// This struct embeds jwt.RegisteredClaims for standard JWT fields.
// Used in GraphQL resolvers for auth middleware (e.g., RequireAuth in middleware/auth.go).
type Claims struct {
	UserID uint         `json:"user_id"` // User ID as uint for GORM models
	Email  string       `json:"email"`   // User email for identification
	Role   app.UserRole `json:"role"`    // User role (e.g., ADMIN) for authorization
	jwt.RegisteredClaims
}

// TokenPair represents access and refresh tokens returned after login.
// Used in GraphQL mutations like Login and Register.
type TokenPair struct {
	AccessToken  string `json:"access_token"`  // Short-lived token for API access
	RefreshToken string `json:"refresh_token"` // Long-lived token to refresh access token
	ExpiresIn    int64  `json:"expires_in"`    // Unix timestamp for access token expiry
	TokenType    string `json:"token_type"`    // Always "Bearer" for standard auth
}

// GenerateTokenPair generates access and refresh tokens for a user.
// Access token expires in 15 minutes; refresh token in 7 days.
// Called from auth service (e.g., Login in services/auth_service.go).
// Uses Go's jwt/v5 for secure token generation.
func GenerateTokenPair(user *app.User) (*TokenPair, error) {
	// Access token: Short expiry for security
	accessExpiry := time.Now().Add(15 * time.Minute)
	accessClaims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.FormatUint(uint64(user.ID), 10), // Convert ID to string (e.g., "2")
			Issuer:    "portfolio-api",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	// Refresh token: Longer expiry for convenience
	refreshExpiry := time.Now().Add(7 * 24 * time.Hour)
	refreshClaims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.FormatUint(uint64(user.ID), 10), // Same conversion for consistency
			Issuer:    "portfolio-api-refresh",
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresIn:    accessExpiry.Unix(), // Unix timestamp for client-side expiry checks
		TokenType:    "Bearer",
	}, nil
}

// ValidateToken validates a JWT token and returns claims.
// Used in middleware (e.g., RequireAuth) and GraphQL helpers.
// Returns error for invalid/expired tokens.
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// RefreshAccessToken generates a new access token from a valid refresh token.
// Verifies the refresh token's issuer before issuing new tokens.
// Used in GraphQL mutation RefreshToken.
func RefreshAccessToken(refreshTokenString string) (*TokenPair, error) {
	claims, err := ValidateToken(refreshTokenString)
	if err != nil {
		return nil, err
	}

	// Ensure this is a refresh token (not access)
	if claims.Issuer != "portfolio-api-refresh" {
		return nil, errors.New("invalid refresh token")
	}

	// Reconstruct user from claims (avoid DB query for performance)
	user := &app.User{
		ID:    claims.UserID,
		Email: claims.Email,
		Role:  claims.Role,
	}

	return GenerateTokenPair(user)
}
