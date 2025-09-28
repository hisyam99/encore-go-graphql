package services

import (
	"context"
	"errors"

	"encore.app/app"
	"encore.app/app/repositories"
	"encore.app/app/utils"
	"gorm.io/gorm"
)

// UserService provides business logic for user operations
type UserService struct {
	userRepo repositories.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// CreateUser creates a new user with validation
func (s *UserService) CreateUser(ctx context.Context, name, email string) (*app.User, error) {
	// Validate input
	if err := utils.ValidateRequired(name, "name"); err != nil {
		return nil, err
	}
	if err := utils.ValidateRequired(email, "email"); err != nil {
		return nil, err
	}
	if err := utils.ValidateEmail(email); err != nil {
		return nil, err
	}
	if err := utils.ValidateMaxLength(name, "name", 255); err != nil {
		return nil, err
	}

	// Sanitize input
	name = utils.SanitizeString(name)
	email = utils.SanitizeString(email)

	// Check if user already exists
	if _, err := s.userRepo.GetByEmail(ctx, email); err == nil {
		return nil, errors.New("user with this email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Create user
	user := &app.User{
		Name:  name,
		Email: email,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(ctx context.Context, id uint) (*app.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(ctx context.Context, id uint, name, email *string) (*app.User, error) {
	// Get existing user
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if name != nil {
		if err := utils.ValidateRequired(*name, "name"); err != nil {
			return nil, err
		}
		if err := utils.ValidateMaxLength(*name, "name", 255); err != nil {
			return nil, err
		}
		user.Name = utils.SanitizeString(*name)
	}

	if email != nil {
		if err := utils.ValidateRequired(*email, "email"); err != nil {
			return nil, err
		}
		if err := utils.ValidateEmail(*email); err != nil {
			return nil, err
		}

		// Check if email is already taken by another user
		if existingUser, err := s.userRepo.GetByEmail(ctx, *email); err == nil && existingUser.ID != id {
			return nil, errors.New("email is already taken")
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		user.Email = utils.SanitizeString(*email)
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	// Check if user exists
	if _, err := s.userRepo.GetByID(ctx, id); err != nil {
		return err
	}

	return s.userRepo.Delete(ctx, id)
}

// ListUsers retrieves users with pagination
func (s *UserService) ListUsers(ctx context.Context, params repositories.PaginationParams) (*repositories.PaginatedResult[app.User], error) {
	return s.userRepo.List(ctx, params)
}

// CategoryService provides business logic for category operations
type CategoryService struct {
	categoryRepo repositories.CategoryRepository
}

// NewCategoryService creates a new category service
func NewCategoryService(categoryRepo repositories.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

// CreateCategory creates a new category with validation
func (s *CategoryService) CreateCategory(ctx context.Context, name, description string) (*app.Category, error) {
	// Validate input
	if err := utils.ValidateRequired(name, "name"); err != nil {
		return nil, err
	}
	if err := utils.ValidateMaxLength(name, "name", 100); err != nil {
		return nil, err
	}
	if description != "" {
		if err := utils.ValidateMaxLength(description, "description", 500); err != nil {
			return nil, err
		}
	}

	// Sanitize input
	name = utils.SanitizeString(name)
	description = utils.SanitizeString(description)

	// Check if category already exists
	if _, err := s.categoryRepo.GetByName(ctx, name); err == nil {
		return nil, errors.New("category with this name already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Create category
	category := &app.Category{
		Name:        name,
		Description: description,
	}

	if err := s.categoryRepo.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

// GetCategory retrieves a category by ID
func (s *CategoryService) GetCategory(ctx context.Context, id uint) (*app.Category, error) {
	return s.categoryRepo.GetByID(ctx, id)
}

// UpdateCategory updates an existing category
func (s *CategoryService) UpdateCategory(ctx context.Context, id uint, name, description *string) (*app.Category, error) {
	// Get existing category
	category, err := s.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if name != nil {
		if err := utils.ValidateRequired(*name, "name"); err != nil {
			return nil, err
		}
		if err := utils.ValidateMaxLength(*name, "name", 100); err != nil {
			return nil, err
		}

		// Check if name is already taken by another category
		if existingCategory, err := s.categoryRepo.GetByName(ctx, *name); err == nil && existingCategory.ID != id {
			return nil, errors.New("category name is already taken")
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		category.Name = utils.SanitizeString(*name)
	}

	if description != nil {
		if *description != "" {
			if err := utils.ValidateMaxLength(*description, "description", 500); err != nil {
				return nil, err
			}
		}
		category.Description = utils.SanitizeString(*description)
	}

	if err := s.categoryRepo.Update(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

// DeleteCategory deletes a category
func (s *CategoryService) DeleteCategory(ctx context.Context, id uint) error {
	// Check if category exists
	if _, err := s.categoryRepo.GetByID(ctx, id); err != nil {
		return err
	}

	return s.categoryRepo.Delete(ctx, id)
}

// ListCategories retrieves categories with pagination
func (s *CategoryService) ListCategories(ctx context.Context, params repositories.PaginationParams) (*repositories.PaginatedResult[app.Category], error) {
	return s.categoryRepo.List(ctx, params)
}
