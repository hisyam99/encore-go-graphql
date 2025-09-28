package repositories

import (
	"context"
	"math"

	"encore.app/app"
	"gorm.io/gorm"
)

// gormResumeContentRepository implements ResumeContentRepository
type gormResumeContentRepository struct {
	db *gorm.DB
}

// NewResumeContentRepository creates a new resume content repository
func NewResumeContentRepository(db *gorm.DB) ResumeContentRepository {
	return &gormResumeContentRepository{db: db}
}

func (r *gormResumeContentRepository) Create(ctx context.Context, content *app.ResumeContent) error {
	return r.db.WithContext(ctx).Create(content).Error
}

func (r *gormResumeContentRepository) GetByID(ctx context.Context, id uint) (*app.ResumeContent, error) {
	var content app.ResumeContent
	err := r.db.WithContext(ctx).Preload("Category").First(&content, id).Error
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (r *gormResumeContentRepository) Update(ctx context.Context, content *app.ResumeContent) error {
	return r.db.WithContext(ctx).Select("*").Omit("created_at").Updates(content).Error
}

func (r *gormResumeContentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&app.ResumeContent{}, id).Error
}

func (r *gormResumeContentRepository) List(ctx context.Context, params PaginationParams) (*PaginatedResult[app.ResumeContent], error) {
	var contents []app.ResumeContent
	var total int64

	query := r.db.WithContext(ctx).Model(&app.ResumeContent{}).Preload("Category")

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	offset := (params.Page - 1) * params.PageSize
	if params.SortBy != "" {
		order := convertToSnakeCase(params.SortBy)
		if params.SortDesc {
			order += " DESC"
		}
		query = query.Order(order)
	} else {
		query = query.Order("created_at DESC")
	}

	if err := query.Offset(offset).Limit(params.PageSize).Find(&contents).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &PaginatedResult[app.ResumeContent]{
		Data:       contents,
		Total:      total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
	}, nil
}

func (r *gormResumeContentRepository) ListByCategory(ctx context.Context, categoryID uint, params PaginationParams) (*PaginatedResult[app.ResumeContent], error) {
	var contents []app.ResumeContent
	var total int64

	query := r.db.WithContext(ctx).Model(&app.ResumeContent{}).Where("category_id = ?", categoryID).Preload("Category")

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	offset := (params.Page - 1) * params.PageSize
	if params.SortBy != "" {
		order := convertToSnakeCase(params.SortBy)
		if params.SortDesc {
			order += " DESC"
		}
		query = query.Order(order)
	} else {
		query = query.Order("created_at DESC")
	}

	if err := query.Offset(offset).Limit(params.PageSize).Find(&contents).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &PaginatedResult[app.ResumeContent]{
		Data:       contents,
		Total:      total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
	}, nil
}

// gormProjectRepository implements ProjectRepository
type gormProjectRepository struct {
	db *gorm.DB
}

// NewProjectRepository creates a new project repository
func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &gormProjectRepository{db: db}
}

func (r *gormProjectRepository) Create(ctx context.Context, project *app.Project) error {
	return r.db.WithContext(ctx).Create(project).Error
}

func (r *gormProjectRepository) GetByID(ctx context.Context, id uint) (*app.Project, error) {
	var project app.Project
	err := r.db.WithContext(ctx).Preload("User").First(&project, id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *gormProjectRepository) Update(ctx context.Context, project *app.Project) error {
	return r.db.WithContext(ctx).Select("*").Omit("created_at").Updates(project).Error
}

func (r *gormProjectRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&app.Project{}, id).Error
}

func (r *gormProjectRepository) List(ctx context.Context, params PaginationParams) (*PaginatedResult[app.Project], error) {
	var projects []app.Project
	var total int64

	query := r.db.WithContext(ctx).Model(&app.Project{}).Preload("User")

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	offset := (params.Page - 1) * params.PageSize
	if params.SortBy != "" {
		order := convertToSnakeCase(params.SortBy)
		if params.SortDesc {
			order += " DESC"
		}
		query = query.Order(order)
	} else {
		query = query.Order("created_at DESC")
	}

	if err := query.Offset(offset).Limit(params.PageSize).Find(&projects).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &PaginatedResult[app.Project]{
		Data:       projects,
		Total:      total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
	}, nil
}

func (r *gormProjectRepository) ListByUser(ctx context.Context, userID uint, params PaginationParams) (*PaginatedResult[app.Project], error) {
	var projects []app.Project
	var total int64

	query := r.db.WithContext(ctx).Model(&app.Project{}).Where("user_id = ?", userID).Preload("User")

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	offset := (params.Page - 1) * params.PageSize
	if params.SortBy != "" {
		order := convertToSnakeCase(params.SortBy)
		if params.SortDesc {
			order += " DESC"
		}
		query = query.Order(order)
	} else {
		query = query.Order("created_at DESC")
	}

	if err := query.Offset(offset).Limit(params.PageSize).Find(&projects).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &PaginatedResult[app.Project]{
		Data:       projects,
		Total:      total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
	}, nil
}
