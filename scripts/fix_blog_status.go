package main

import (
	"context"
	"log"

	"encore.app/app/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialize database connection
	// You'll need to replace this with your actual database connection string
	dsn := "host=localhost user=postgres password=password dbname=portfolio port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create repository
	blogRepo := repositories.NewBlogRepository(db)

	// Execute the fix
	ctx := context.Background()
	err = blogRepo.FixBlogStatus(ctx)
	if err != nil {
		log.Fatal("Failed to fix blog status:", err)
	}

	log.Println("Successfully fixed blog statuses and published_at fields")
}
