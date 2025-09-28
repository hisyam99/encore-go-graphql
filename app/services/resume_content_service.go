package services

import (
	"context"

	"encore.app/app"
	"encore.app/app/repositories"
	"encore.app/app/utils"
)

// ResumeContentService provides business logic for resume content operations
type ResumeContentService struct {
	resumeContentRepo repositories.ResumeContentRepository
	categoryRepo      repositories.CategoryRepository
}

// NewResumeContentService creates a new resume content service
func NewResumeContentService(resumeContentRepo repositories.ResumeContentRepository, categoryRepo repositories.CategoryRepository) *ResumeContentService {
	return &ResumeContentService{
		resumeContentRepo: resumeContentRepo,
		categoryRepo:      categoryRepo,
	}
}

// CreateResumeContent creates a new resume content with validation
func (s *ResumeContentService) CreateResumeContent(ctx context.Context, title, description, detail string, categoryID uint) (*app.ResumeContent, error) {
	// Validate required fields
	if err := utils.ValidateRequired(title, "title"); err != nil {
		return nil, err
	}

	// Validate lengths
	if err := utils.ValidateMaxLength(title, "title", 255); err != nil {
		return nil, err
	}
	if description != "" {
		if err := utils.ValidateMaxLength(description, "description", 500); err != nil {
			return nil, err
		}
	}

	// Sanitize input
	title = utils.SanitizeString(title)
	description = utils.SanitizeString(description)
	detail = utils.SanitizeString(detail)

	// Verify that category exists
	if _, err := s.categoryRepo.GetByID(ctx, categoryID); err != nil {
		return nil, err
	}

	// Create resume content
	resumeContent := &app.ResumeContent{
		Title:       title,
		Description: description,
		Detail:      detail,
		CategoryID:  categoryID,
	}

	if err := s.resumeContentRepo.Create(ctx, resumeContent); err != nil {
		return nil, err
	}

	return resumeContent, nil
}

// GetResumeContent retrieves a resume content by ID
func (s *ResumeContentService) GetResumeContent(ctx context.Context, id uint) (*app.ResumeContent, error) {
	return s.resumeContentRepo.GetByID(ctx, id)
}

// UpdateResumeContent updates an existing resume content
func (s *ResumeContentService) UpdateResumeContent(ctx context.Context, id uint, title, description, detail *string, categoryID *uint) (*app.ResumeContent, error) {
	// Get existing resume content
	resumeContent, err := s.resumeContentRepo.GetByID(ctx, id)
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
		resumeContent.Title = utils.SanitizeString(*title)
	}

	if description != nil {
		if *description != "" {
			if err := utils.ValidateMaxLength(*description, "description", 500); err != nil {
				return nil, err
			}
		}
		resumeContent.Description = utils.SanitizeString(*description)
	}

	if detail != nil {
		resumeContent.Detail = utils.SanitizeString(*detail)
	}

	if categoryID != nil {
		// Verify that category exists
		if _, err := s.categoryRepo.GetByID(ctx, *categoryID); err != nil {
			return nil, err
		}
		resumeContent.CategoryID = *categoryID
	}

	if err := s.resumeContentRepo.Update(ctx, resumeContent); err != nil {
		return nil, err
	}

	return resumeContent, nil
}

// DeleteResumeContent deletes a resume content
func (s *ResumeContentService) DeleteResumeContent(ctx context.Context, id uint) error {
	// Check if resume content exists
	if _, err := s.resumeContentRepo.GetByID(ctx, id); err != nil {
		return err
	}

	return s.resumeContentRepo.Delete(ctx, id)
}

// ListResumeContents retrieves resume contents with pagination
func (s *ResumeContentService) ListResumeContents(ctx context.Context, params repositories.PaginationParams) (*repositories.PaginatedResult[app.ResumeContent], error) {
	return s.resumeContentRepo.List(ctx, params)
}

// ListResumeContentsByCategory retrieves resume contents by category with pagination
func (s *ResumeContentService) ListResumeContentsByCategory(ctx context.Context, categoryID uint, params repositories.PaginationParams) (*repositories.PaginatedResult[app.ResumeContent], error) {
	// Verify that category exists
	if _, err := s.categoryRepo.GetByID(ctx, categoryID); err != nil {
		return nil, err
	}

	return s.resumeContentRepo.ListByCategory(ctx, categoryID, params)
}
