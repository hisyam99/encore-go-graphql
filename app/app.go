package app

import (
	"encore.dev/storage/sqldb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//encore:service
type Service struct {
	db *gorm.DB
}

var blogDB = sqldb.NewDatabase("app", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

// initService initializes the site service.
// It is automatically called by Encore on service startup.
func initService() (*Service, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: blogDB.Stdlib(),
	}))
	if err != nil {
		return nil, err
	}
	return &Service{db: db}, nil
}

// New creates a new instance of Service (for external use, e.g., by graphql/service.go).
func New() (*Service, error) {
	return initService()
}

// DB returns the gorm.DB instance for external use.
func (s *Service) DB() *gorm.DB {
	return s.db
}