package repositories

import (
	"context"

	"encore.app/app"
)

// PaginationParams represents pagination parameters
type PaginationParams struct {
	Page     int
	PageSize int
	SortBy   string
	SortDesc bool
}

// PaginatedResult represents a paginated result
type PaginatedResult[T any] struct {
	Data       []T   `json:"data"`
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	TotalPages int   `json:"totalPages"`
}

// UserRepository defines the interface for user operations
type UserRepository interface {
	Create(ctx context.Context, user *app.User) error
	GetByID(ctx context.Context, id uint) (*app.User, error)
	GetByEmail(ctx context.Context, email string) (*app.User, error)
	Update(ctx context.Context, user *app.User) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, params PaginationParams) (*PaginatedResult[app.User], error)
}

// CategoryRepository defines the interface for category operations
type CategoryRepository interface {
	Create(ctx context.Context, category *app.Category) error
	GetByID(ctx context.Context, id uint) (*app.Category, error)
	GetByName(ctx context.Context, name string) (*app.Category, error)
	Update(ctx context.Context, category *app.Category) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, params PaginationParams) (*PaginatedResult[app.Category], error)
}

// ResumeContentRepository defines the interface for resume content operations
type ResumeContentRepository interface {
	Create(ctx context.Context, content *app.ResumeContent) error
	GetByID(ctx context.Context, id uint) (*app.ResumeContent, error)
	Update(ctx context.Context, content *app.ResumeContent) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, params PaginationParams) (*PaginatedResult[app.ResumeContent], error)
	ListByCategory(ctx context.Context, categoryID uint, params PaginationParams) (*PaginatedResult[app.ResumeContent], error)
}

// ProjectRepository defines the interface for project operations
type ProjectRepository interface {
	Create(ctx context.Context, project *app.Project) error
	GetByID(ctx context.Context, id uint) (*app.Project, error)
	Update(ctx context.Context, project *app.Project) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, params PaginationParams) (*PaginatedResult[app.Project], error)
	ListByUser(ctx context.Context, userID uint, params PaginationParams) (*PaginatedResult[app.Project], error)
}

// BlogRepository defines the interface for blog operations
type BlogRepository interface {
	Create(ctx context.Context, blog *app.Blog) error
	GetByID(ctx context.Context, id uint) (*app.Blog, error)
	GetBySlug(ctx context.Context, slug string) (*app.Blog, error)
	Update(ctx context.Context, blog *app.Blog) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, params PaginationParams) (*PaginatedResult[app.Blog], error)
	ListPublished(ctx context.Context, params PaginationParams) (*PaginatedResult[app.Blog], error)
	ListByStatus(ctx context.Context, status app.BlogStatus, params PaginationParams) (*PaginatedResult[app.Blog], error)
}
