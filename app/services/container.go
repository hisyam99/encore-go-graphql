package services

import (
	"encore.app/app/repositories"
	"gorm.io/gorm"
)

// Services contains all application services
type Services struct {
	User          *UserService
	Category      *CategoryService
	ResumeContent *ResumeContentService
	Project       *ProjectService
	Blog          *BlogService
}

// NewServices creates a new services container with all dependencies
func NewServices(db *gorm.DB) *Services {
	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	resumeContentRepo := repositories.NewResumeContentRepository(db)
	projectRepo := repositories.NewProjectRepository(db)
	blogRepo := repositories.NewBlogRepository(db)

	// Initialize services
	userService := NewUserService(userRepo)
	categoryService := NewCategoryService(categoryRepo)
	resumeContentService := NewResumeContentService(resumeContentRepo, categoryRepo)
	projectService := NewProjectService(projectRepo)
	blogService := NewBlogService(blogRepo)

	return &Services{
		User:          userService,
		Category:      categoryService,
		ResumeContent: resumeContentService,
		Project:       projectService,
		Blog:          blogService,
	}
}
