package app

import "time"

// Contoh tabel untuk portofolio
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	CreatedAt time.Time
}

type Project struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	UserID      uint
}

type Blog struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Content   string
	CreatedAt time.Time
}

type Resume struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Category    string
}
