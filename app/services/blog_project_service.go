package services

import (
	"context"
	"errors"
	"strings"

	"encore.app/app"
	"encore.app/app/repositories"
	"encore.app/app/utils"
	"gorm.io/gorm"
)

// normalizeBlogStatus normalizes blog status to lowercase to ensure consistency
func normalizeBlogStatus(status app.BlogStatus) app.BlogStatus {
	switch strings.ToLower(string(status)) {
	case strings.ToLower(string(app.BlogStatusPublished)):
		return app.BlogStatusPublished
	default:
		return app.BlogStatusDraft
	}
}

// BlogService provides business logic for blog operations
type BlogService struct {
	blogRepo repositories.BlogRepository
}

// NewBlogService creates a new blog service
func NewBlogService(blogRepo repositories.BlogRepository) *BlogService {
	return &BlogService{blogRepo: blogRepo}
}

// CreateBlog creates a new blog with validation
func (s *BlogService) CreateBlog(ctx context.Context, title, content, summary, slug, author, metaDescription string, status app.BlogStatus, tags []string) (*app.Blog, error) {
	// Validate required fields
	if err := utils.ValidateRequired(title, "title"); err != nil {
		return nil, err
	}
	if err := utils.ValidateRequired(content, "content"); err != nil {
		return nil, err
	}
	if err := utils.ValidateRequired(slug, "slug"); err != nil {
		return nil, err
	}

	// Validate lengths
	if err := utils.ValidateMaxLength(title, "title", 255); err != nil {
		return nil, err
	}
	if err := utils.ValidateMaxLength(summary, "summary", 500); err != nil {
		return nil, err
	}
	if err := utils.ValidateMaxLength(author, "author", 255); err != nil {
		return nil, err
	}
	if err := utils.ValidateMaxLength(metaDescription, "metaDescription", 160); err != nil {
		return nil, err
	}

	// Validate slug format
	if err := utils.ValidateSlug(slug); err != nil {
		return nil, err
	}

	// Sanitize input
	title = utils.SanitizeString(title)
	content = utils.SanitizeString(content)
	summary = utils.SanitizeString(summary)
	author = utils.SanitizeString(author)
	metaDescription = utils.SanitizeString(metaDescription)

	// Check if slug already exists
	if _, err := s.blogRepo.GetBySlug(ctx, slug); err == nil {
		return nil, errors.New("blog with this slug already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Normalize status to ensure consistency
	status = normalizeBlogStatus(status)

	// Create blog
	blog := &app.Blog{
		Title:           title,
		Content:         content,
		Summary:         summary,
		Slug:            slug,
		Author:          author,
		Status:          status,
		Tags:            app.StringArray(tags),
		MetaDescription: metaDescription,
	}

	if err := s.blogRepo.Create(ctx, blog); err != nil {
		return nil, err
	}

	return blog, nil
}

// GetBlog retrieves a blog by ID
func (s *BlogService) GetBlog(ctx context.Context, id uint) (*app.Blog, error) {
	return s.blogRepo.GetByID(ctx, id)
}

// GetBlogBySlug retrieves a blog by slug
func (s *BlogService) GetBlogBySlug(ctx context.Context, slug string) (*app.Blog, error) {
	return s.blogRepo.GetBySlug(ctx, slug)
}

// UpdateBlog updates an existing blog
func (s *BlogService) UpdateBlog(ctx context.Context, id uint, title, content, summary, slug, author, metaDescription *string, status *app.BlogStatus, tags *[]string) (*app.Blog, error) {
	// Get existing blog
	blog, err := s.blogRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if title != nil {
		if err := utils.ValidateRequired(*title, "title"); err != nil {
			return nil, err
		}
		if err := utils.ValidateMaxLength(*title, "title", 255); err != nil {
			return nil, err
		}
		blog.Title = utils.SanitizeString(*title)
	}

	if content != nil {
		if err := utils.ValidateRequired(*content, "content"); err != nil {
			return nil, err
		}
		blog.Content = utils.SanitizeString(*content)
	}

	if summary != nil {
		if *summary != "" {
			if err := utils.ValidateMaxLength(*summary, "summary", 500); err != nil {
				return nil, err
			}
		}
		blog.Summary = utils.SanitizeString(*summary)
	}

	if slug != nil {
		if err := utils.ValidateRequired(*slug, "slug"); err != nil {
			return nil, err
		}
		if err := utils.ValidateSlug(*slug); err != nil {
			return nil, err
		}

		// Check if slug is already taken by another blog
		if existingBlog, err := s.blogRepo.GetBySlug(ctx, *slug); err == nil && existingBlog.ID != id {
			return nil, errors.New("slug is already taken")
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		blog.Slug = *slug
	}

	if author != nil {
		if *author != "" {
			if err := utils.ValidateMaxLength(*author, "author", 255); err != nil {
				return nil, err
			}
		}
		blog.Author = utils.SanitizeString(*author)
	}

	if metaDescription != nil {
		if *metaDescription != "" {
			if err := utils.ValidateMaxLength(*metaDescription, "metaDescription", 160); err != nil {
				return nil, err
			}
		}
		blog.MetaDescription = utils.SanitizeString(*metaDescription)
	}

	if status != nil {
		blog.Status = normalizeBlogStatus(*status)
	}

	if tags != nil {
		blog.Tags = app.StringArray(*tags)
	}

	if err := s.blogRepo.Update(ctx, blog); err != nil {
		return nil, err
	}

	return blog, nil
}

// DeleteBlog deletes a blog
func (s *BlogService) DeleteBlog(ctx context.Context, id uint) error {
	// Check if blog exists
	if _, err := s.blogRepo.GetByID(ctx, id); err != nil {
		return err
	}

	return s.blogRepo.Delete(ctx, id)
}

// ListBlogs retrieves blogs with pagination
func (s *BlogService) ListBlogs(ctx context.Context, params repositories.PaginationParams) (*repositories.PaginatedResult[app.Blog], error) {
	return s.blogRepo.List(ctx, params)
}

// ListPublishedBlogs retrieves published blogs with pagination
func (s *BlogService) ListPublishedBlogs(ctx context.Context, params repositories.PaginationParams) (*repositories.PaginatedResult[app.Blog], error) {
	return s.blogRepo.ListPublished(ctx, params)
}

// ListBlogsByStatus retrieves blogs by status with pagination
func (s *BlogService) ListBlogsByStatus(ctx context.Context, status app.BlogStatus, params repositories.PaginationParams) (*repositories.PaginatedResult[app.Blog], error) {
	return s.blogRepo.ListByStatus(ctx, normalizeBlogStatus(status), params)
}

// FixBlogStatus normalizes existing blog statuses and sets published_at for published blogs
func (s *BlogService) FixBlogStatus(ctx context.Context) error {
	return s.blogRepo.FixBlogStatus(ctx)
}

// ProjectService provides business logic for project operations
type ProjectService struct {
	projectRepo repositories.ProjectRepository
}

// NewProjectService creates a new project service
func NewProjectService(projectRepo repositories.ProjectRepository) *ProjectService {
	return &ProjectService{projectRepo: projectRepo}
}

// CreateProject creates a new project with validation
func (s *ProjectService) CreateProject(ctx context.Context, title, description string, userID *uint) (*app.Project, error) {
	// Validate required fields
	if err := utils.ValidateRequired(title, "title"); err != nil {
		return nil, err
	}

	// Validate lengths
	if err := utils.ValidateMaxLength(title, "title", 255); err != nil {
		return nil, err
	}
	if description != "" {
		if err := utils.ValidateMaxLength(description, "description", 1000); err != nil {
			return nil, err
		}
	}

	// Sanitize input
	title = utils.SanitizeString(title)
	description = utils.SanitizeString(description)

	// Create project
	project := &app.Project{
		Title:       title,
		Description: description,
		UserID:      userID,
	}

	if err := s.projectRepo.Create(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}

// GetProject retrieves a project by ID
func (s *ProjectService) GetProject(ctx context.Context, id uint) (*app.Project, error) {
	return s.projectRepo.GetByID(ctx, id)
}

// UpdateProject updates an existing project
func (s *ProjectService) UpdateProject(ctx context.Context, id uint, title, description *string, userID *uint) (*app.Project, error) {
	// Get existing project
	project, err := s.projectRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if title != nil {
		if err := utils.ValidateRequired(*title, "title"); err != nil {
			return nil, err
		}
		if err := utils.ValidateMaxLength(*title, "title", 255); err != nil {
			return nil, err
		}
		project.Title = utils.SanitizeString(*title)
	}

	if description != nil {
		if *description != "" {
			if err := utils.ValidateMaxLength(*description, "description", 1000); err != nil {
				return nil, err
			}
		}
		project.Description = utils.SanitizeString(*description)
	}

	if userID != nil {
		project.UserID = userID
	}

	if err := s.projectRepo.Update(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}

// DeleteProject deletes a project
func (s *ProjectService) DeleteProject(ctx context.Context, id uint) error {
	// Check if project exists
	if _, err := s.projectRepo.GetByID(ctx, id); err != nil {
		return err
	}

	return s.projectRepo.Delete(ctx, id)
}

// ListProjects retrieves projects with pagination
func (s *ProjectService) ListProjects(ctx context.Context, params repositories.PaginationParams) (*repositories.PaginatedResult[app.Project], error) {
	return s.projectRepo.List(ctx, params)
}

// ListProjectsByUser retrieves projects by user with pagination
func (s *ProjectService) ListProjectsByUser(ctx context.Context, userID uint, params repositories.PaginationParams) (*repositories.PaginatedResult[app.Project], error) {
	return s.projectRepo.ListByUser(ctx, userID, params)
}
