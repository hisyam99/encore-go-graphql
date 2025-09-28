package repositories

import (
	"context"
	"math"

	"encore.app/app"
	"gorm.io/gorm"
)

// gormUserRepository implements UserRepository
type gormUserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) Create(ctx context.Context, user *app.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *gormUserRepository) GetByID(ctx context.Context, id uint) (*app.User, error) {
	var user app.User
	err := r.db.WithContext(ctx).Preload("Projects").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) GetByEmail(ctx context.Context, email string) (*app.User, error) {
	var user app.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) Update(ctx context.Context, user *app.User) error {
	return r.db.WithContext(ctx).Select("*").Omit("created_at").Updates(user).Error
}

func (r *gormUserRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&app.User{}, id).Error
}

func (r *gormUserRepository) List(ctx context.Context, params PaginationParams) (*PaginatedResult[app.User], error) {
	var users []app.User
	var total int64

	query := r.db.WithContext(ctx).Model(&app.User{})

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	offset := (params.Page - 1) * params.PageSize
	if params.SortBy != "" {
		order := params.SortBy
		if params.SortDesc {
			order += " DESC"
		}
		query = query.Order(order)
	} else {
		query = query.Order("created_at DESC")
	}

	if err := query.Offset(offset).Limit(params.PageSize).Find(&users).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &PaginatedResult[app.User]{
		Data:       users,
		Total:      total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
	}, nil
}

// gormCategoryRepository implements CategoryRepository
type gormCategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new category repository
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &gormCategoryRepository{db: db}
}

func (r *gormCategoryRepository) Create(ctx context.Context, category *app.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *gormCategoryRepository) GetByID(ctx context.Context, id uint) (*app.Category, error) {
	var category app.Category
	err := r.db.WithContext(ctx).Preload("ResumeContents").First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *gormCategoryRepository) GetByName(ctx context.Context, name string) (*app.Category, error) {
	var category app.Category
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *gormCategoryRepository) Update(ctx context.Context, category *app.Category) error {
	return r.db.WithContext(ctx).Select("*").Omit("created_at").Updates(category).Error
}

func (r *gormCategoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&app.Category{}, id).Error
}

func (r *gormCategoryRepository) List(ctx context.Context, params PaginationParams) (*PaginatedResult[app.Category], error) {
	var categories []app.Category
	var total int64

	query := r.db.WithContext(ctx).Model(&app.Category{})

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	offset := (params.Page - 1) * params.PageSize
	if params.SortBy != "" {
		order := params.SortBy
		if params.SortDesc {
			order += " DESC"
		}
		query = query.Order(order)
	} else {
		query = query.Order("name ASC")
	}

	if err := query.Offset(offset).Limit(params.PageSize).Find(&categories).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &PaginatedResult[app.Category]{
		Data:       categories,
		Total:      total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
	}, nil
}
