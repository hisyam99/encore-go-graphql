package dataloader

import (
	"context"
	"time"

	"encore.app/app"
	"github.com/graph-gophers/dataloader/v7"
	"gorm.io/gorm"
)

// DataLoaders contains all data loaders for batching database queries
type DataLoaders struct {
	// User related loaders
	UserProjectLoader *dataloader.Loader[uint, []*app.Project]

	// Category related loaders
	CategoryResumeContentLoader *dataloader.Loader[uint, []*app.ResumeContent]

	// ResumeContent related loaders
	ResumeContentCategoryLoader *dataloader.Loader[uint, *app.Category]

	// Project related loaders
	ProjectUserLoader *dataloader.Loader[uint, *app.User]
}

// NewDataLoaders creates a new instance of all data loaders
func NewDataLoaders(db *gorm.DB) *DataLoaders {
	return &DataLoaders{
		UserProjectLoader: dataloader.NewBatchedLoader(
			userProjectBatchFunc(db),
			dataloader.WithWait[uint, []*app.Project](time.Millisecond),
		),
		CategoryResumeContentLoader: dataloader.NewBatchedLoader(
			categoryResumeContentBatchFunc(db),
			dataloader.WithWait[uint, []*app.ResumeContent](time.Millisecond),
		),
		ResumeContentCategoryLoader: dataloader.NewBatchedLoader(
			resumeContentCategoryBatchFunc(db),
			dataloader.WithWait[uint, *app.Category](time.Millisecond),
		),
		ProjectUserLoader: dataloader.NewBatchedLoader(
			projectUserBatchFunc(db),
			dataloader.WithWait[uint, *app.User](time.Millisecond),
		),
	}
}

// userProjectBatchFunc batches project lookups by user IDs
func userProjectBatchFunc(db *gorm.DB) dataloader.BatchFunc[uint, []*app.Project] {
	return func(ctx context.Context, userIDs []uint) []*dataloader.Result[[]*app.Project] {
		// Query all projects for the given user IDs in a single database call
		var projects []app.Project
		if err := db.WithContext(ctx).Where("user_id IN ?", userIDs).Find(&projects).Error; err != nil {
			// Return error for all user IDs
			results := make([]*dataloader.Result[[]*app.Project], len(userIDs))
			for i := range userIDs {
				results[i] = &dataloader.Result[[]*app.Project]{Error: err}
			}
			return results
		}

		// Group projects by user_id
		projectMap := make(map[uint][]*app.Project)
		for i := range projects {
			project := &projects[i]
			if project.UserID != nil {
				projectMap[*project.UserID] = append(projectMap[*project.UserID], project)
			}
		}

		// Create results for each requested user ID
		results := make([]*dataloader.Result[[]*app.Project], len(userIDs))
		for i, userID := range userIDs {
			results[i] = &dataloader.Result[[]*app.Project]{
				Data: projectMap[userID], // Will be nil/empty slice if no projects found
			}
		}
		return results
	}
}

// categoryResumeContentBatchFunc batches resume content lookups by category IDs
func categoryResumeContentBatchFunc(db *gorm.DB) dataloader.BatchFunc[uint, []*app.ResumeContent] {
	return func(ctx context.Context, categoryIDs []uint) []*dataloader.Result[[]*app.ResumeContent] {
		// Query all resume contents for the given category IDs in a single database call
		var resumeContents []app.ResumeContent
		if err := db.WithContext(ctx).Where("category_id IN ?", categoryIDs).Find(&resumeContents).Error; err != nil {
			// Return error for all category IDs
			results := make([]*dataloader.Result[[]*app.ResumeContent], len(categoryIDs))
			for i := range categoryIDs {
				results[i] = &dataloader.Result[[]*app.ResumeContent]{Error: err}
			}
			return results
		}

		// Group resume contents by category_id
		resumeMap := make(map[uint][]*app.ResumeContent)
		for i := range resumeContents {
			resumeContent := &resumeContents[i]
			resumeMap[resumeContent.CategoryID] = append(resumeMap[resumeContent.CategoryID], resumeContent)
		}

		// Create results for each requested category ID
		results := make([]*dataloader.Result[[]*app.ResumeContent], len(categoryIDs))
		for i, categoryID := range categoryIDs {
			results[i] = &dataloader.Result[[]*app.ResumeContent]{
				Data: resumeMap[categoryID], // Will be nil/empty slice if no resume contents found
			}
		}
		return results
	}
}

// resumeContentCategoryBatchFunc batches category lookups by category IDs
func resumeContentCategoryBatchFunc(db *gorm.DB) dataloader.BatchFunc[uint, *app.Category] {
	return func(ctx context.Context, categoryIDs []uint) []*dataloader.Result[*app.Category] {
		// Query all categories for the given category IDs in a single database call
		var categories []app.Category
		if err := db.WithContext(ctx).Where("id IN ?", categoryIDs).Find(&categories).Error; err != nil {
			// Return error for all category IDs
			results := make([]*dataloader.Result[*app.Category], len(categoryIDs))
			for i := range categoryIDs {
				results[i] = &dataloader.Result[*app.Category]{Error: err}
			}
			return results
		}

		// Create map for quick lookup
		categoryMap := make(map[uint]*app.Category)
		for i := range categories {
			category := &categories[i]
			categoryMap[category.ID] = category
		}

		// Create results for each requested category ID
		results := make([]*dataloader.Result[*app.Category], len(categoryIDs))
		for i, categoryID := range categoryIDs {
			results[i] = &dataloader.Result[*app.Category]{
				Data: categoryMap[categoryID], // Will be nil if category not found
			}
		}
		return results
	}
}

// projectUserBatchFunc batches user lookups by user IDs (for projects)
func projectUserBatchFunc(db *gorm.DB) dataloader.BatchFunc[uint, *app.User] {
	return func(ctx context.Context, userIDs []uint) []*dataloader.Result[*app.User] {
		// Query all users for the given user IDs in a single database call
		var users []app.User
		if err := db.WithContext(ctx).Where("id IN ?", userIDs).Find(&users).Error; err != nil {
			// Return error for all user IDs
			results := make([]*dataloader.Result[*app.User], len(userIDs))
			for i := range userIDs {
				results[i] = &dataloader.Result[*app.User]{Error: err}
			}
			return results
		}

		// Create map for quick lookup
		userMap := make(map[uint]*app.User)
		for i := range users {
			user := &users[i]
			userMap[user.ID] = user
		}

		// Create results for each requested user ID
		results := make([]*dataloader.Result[*app.User], len(userIDs))
		for i, userID := range userIDs {
			results[i] = &dataloader.Result[*app.User]{
				Data: userMap[userID], // Will be nil if user not found
			}
		}
		return results
	}
}
