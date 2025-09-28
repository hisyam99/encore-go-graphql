package repositories

import (
	"context"
	"math"

	"encore.app/app"
	"gorm.io/gorm"
)

// gormBlogRepository implements BlogRepository
type gormBlogRepository struct {
	db *gorm.DB
}

// NewBlogRepository creates a new blog repository
func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &gormBlogRepository{db: db}
}

func (r *gormBlogRepository) Create(ctx context.Context, blog *app.Blog) error {
	return r.db.WithContext(ctx).Create(blog).Error
}

func (r *gormBlogRepository) GetByID(ctx context.Context, id uint) (*app.Blog, error) {
	var blog app.Blog
	err := r.db.WithContext(ctx).First(&blog, id).Error
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

func (r *gormBlogRepository) GetBySlug(ctx context.Context, slug string) (*app.Blog, error) {
	var blog app.Blog
	err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&blog).Error
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

func (r *gormBlogRepository) Update(ctx context.Context, blog *app.Blog) error {
	return r.db.WithContext(ctx).Select("*").Omit("created_at").Updates(blog).Error
}

func (r *gormBlogRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&app.Blog{}, id).Error
}

func (r *gormBlogRepository) List(ctx context.Context, params PaginationParams) (*PaginatedResult[app.Blog], error) {
	var blogs []app.Blog
	var total int64

	query := r.db.WithContext(ctx).Model(&app.Blog{})

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

	if err := query.Offset(offset).Limit(params.PageSize).Find(&blogs).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &PaginatedResult[app.Blog]{
		Data:       blogs,
		Total:      total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
	}, nil
}

func (r *gormBlogRepository) ListPublished(ctx context.Context, params PaginationParams) (*PaginatedResult[app.Blog], error) {
	return r.ListByStatus(ctx, app.BlogStatusPublished, params)
}

func (r *gormBlogRepository) ListByStatus(ctx context.Context, status app.BlogStatus, params PaginationParams) (*PaginatedResult[app.Blog], error) {
	var blogs []app.Blog
	var total int64

	query := r.db.WithContext(ctx).Model(&app.Blog{}).Where("status = ?", status)

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
		// For published blogs, order by published_at, for others by created_at
		if status == app.BlogStatusPublished {
			query = query.Order("published_at DESC")
		} else {
			query = query.Order("created_at DESC")
		}
	}

	if err := query.Offset(offset).Limit(params.PageSize).Find(&blogs).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &PaginatedResult[app.Blog]{
		Data:       blogs,
		Total:      total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
	}, nil
}
