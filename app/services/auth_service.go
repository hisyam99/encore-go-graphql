package services

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"encore.app/app"
	"encore.app/app/auth"
	"encore.app/app/repositories"
)

// AuthService handles authentication operations
type AuthService struct {
	userRepo repositories.UserRepository
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

// LoginRequest represents login request
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// RegisterRequest represents registration request
type RegisterRequest struct {
	Name     string       `json:"name" validate:"required,min=2,max=100"`
	Email    string       `json:"email" validate:"required,email"`
	Password string       `json:"password" validate:"required,min=6"`
	Role     app.UserRole `json:"role,omitempty"`
}

// LoginResponse represents login response
type LoginResponse struct {
	User   *app.User       `json:"user"`
	Tokens *auth.TokenPair `json:"tokens"`
}

// validateEmail validates email format
func validateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return errors.New("email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

// Login authenticates user and returns tokens
func (s *AuthService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	// Validate input
	if err := validateEmail(req.Email); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	if len(strings.TrimSpace(req.Password)) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	// Find user by email
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Verify password
	if !user.CheckPassword(req.Password) {
		return nil, errors.New("invalid credentials")
	}

	// Update last login
	now := time.Now()
	user.LastLogin = &now
	if err := s.userRepo.Update(ctx, user); err != nil {
		// Log error but don't fail login
		fmt.Printf("Failed to update last login for user %d: %v\n", user.ID, err)
	}

	// Generate tokens
	tokens, err := auth.GenerateTokenPair(user)
	if err != nil {
		return nil, errors.New("failed to generate tokens")
	}

	// Remove password from response
	user.Password = ""

	return &LoginResponse{
		User:   user,
		Tokens: tokens,
	}, nil
}

// Register creates new user account
func (s *AuthService) Register(ctx context.Context, req *RegisterRequest) (*LoginResponse, error) {
	// Validate input
	if err := validateEmail(req.Email); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	if len(strings.TrimSpace(req.Password)) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	req.Name = strings.TrimSpace(req.Name)
	if len(req.Name) < 2 {
		return nil, errors.New("name must be at least 2 characters")
	}

	// Check if user already exists
	existingUser, _ := s.userRepo.GetByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Set default role if not provided
	role := req.Role
	if role == "" {
		role = app.UserRoleViewer
	}

	// Create user
	user := &app.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, // Will be hashed in BeforeCreate hook
		Role:     role,
		IsActive: true,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate tokens
	tokens, err := auth.GenerateTokenPair(user)
	if err != nil {
		return nil, errors.New("failed to generate tokens")
	}

	// Remove password from response
	user.Password = ""

	return &LoginResponse{
		User:   user,
		Tokens: tokens,
	}, nil
}

// RefreshToken generates new access token
func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*auth.TokenPair, error) {
	if strings.TrimSpace(refreshToken) == "" {
		return nil, errors.New("refresh token is required")
	}

	return auth.RefreshAccessToken(refreshToken)
}

// GetCurrentUser returns current user from token
func (s *AuthService) GetCurrentUser(ctx context.Context, userID uint) (*app.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Remove password from response
	user.Password = ""
	return user, nil
}

// ChangePassword changes user password
func (s *AuthService) ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error {
	if len(strings.TrimSpace(newPassword)) < 6 {
		return errors.New("new password must be at least 6 characters")
	}

	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	// Verify old password
	if !user.CheckPassword(oldPassword) {
		return errors.New("invalid old password")
	}

	// Update password
	user.Password = newPassword // Will be hashed in BeforeUpdate hook
	return s.userRepo.Update(ctx, user)
}

// DeactivateUser deactivates user account
func (s *AuthService) DeactivateUser(ctx context.Context, userID uint) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	user.IsActive = false
	return s.userRepo.Update(ctx, user)
}

// UpdateUserRole updates user role (admin only)
func (s *AuthService) UpdateUserRole(ctx context.Context, userID uint, newRole app.UserRole) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	user.Role = newRole
	return s.userRepo.Update(ctx, user)
}

// GetUserList returns paginated list of users (admin only)
func (s *AuthService) GetUserList(ctx context.Context, params repositories.PaginationParams) (*repositories.PaginatedResult[app.User], error) {
	result, err := s.userRepo.List(ctx, params)
	if err != nil {
		return nil, err
	}

	// Remove passwords from response
	for i := range result.Data {
		result.Data[i].Password = ""
	}

	return result, nil
}

// GetUserByID returns user by ID (for admin or self)
func (s *AuthService) GetUserByID(ctx context.Context, userID uint) (*app.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Remove password from response
	user.Password = ""
	return user, nil
}