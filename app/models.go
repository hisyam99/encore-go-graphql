package app

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// BlogStatus represents the status of a blog post
type BlogStatus string

const (
	BlogStatusDraft     BlogStatus = "draft"
	BlogStatusPublished BlogStatus = "published"
)

// Scan implements the Scanner interface for database reads
func (bs *BlogStatus) Scan(value interface{}) error {
	if value == nil {
		*bs = BlogStatusDraft
		return nil
	}
	if s, ok := value.(string); ok {
		*bs = BlogStatus(s)
		return nil
	}
	return errors.New("cannot scan BlogStatus")
}

// Value implements the Valuer interface for database writes
func (bs BlogStatus) Value() (driver.Value, error) {
	return string(bs), nil
}

// StringArray represents a JSON array of strings in the database
type StringArray []string

// Scan implements the Scanner interface for database reads
func (sa *StringArray) Scan(value interface{}) error {
	if value == nil {
		*sa = StringArray{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, sa)
	case string:
		return json.Unmarshal([]byte(v), sa)
	}
	return errors.New("cannot scan StringArray")
}

// Value implements the Valuer interface for database writes
func (sa StringArray) Value() (driver.Value, error) {
	if len(sa) == 0 {
		return "[]", nil
	}
	return json.Marshal(sa)
}

// User represents a user in the system
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null;size:255" json:"name"`
	Email     string         `gorm:"uniqueIndex;not null;size:255" json:"email"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Projects []Project `json:"projects,omitempty"`
}

// Category represents a category for resume content
type Category struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null;uniqueIndex;size:100" json:"name"`
	Description string         `gorm:"size:500" json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	ResumeContents []ResumeContent `json:"resumeContents,omitempty"`
}

// ResumeContent represents content for resume sections
type ResumeContent struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null;size:255" json:"title"`
	Description string         `gorm:"size:500" json:"description"`
	Detail      string         `gorm:"type:text" json:"detail"`
	CategoryID  uint           `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"categoryId"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Category Category `json:"category,omitempty"`
}

// Project represents a portfolio project
type Project struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null;size:255" json:"title"`
	Description string         `gorm:"size:1000" json:"description"`
	UserID      *uint          `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"userId,omitempty"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	User *User `json:"user,omitempty"`
}

// Blog represents a blog post
type Blog struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Title           string         `gorm:"not null;size:255" json:"title"`
	Content         string         `gorm:"type:text;not null" json:"content"`
	Summary         string         `gorm:"size:500" json:"summary"`
	Slug            string         `gorm:"uniqueIndex;not null;size:255" json:"slug"`
	Author          string         `gorm:"size:255" json:"author"`
	PublishedAt     *time.Time     `json:"publishedAt,omitempty"`
	Status          BlogStatus     `gorm:"default:'draft';index" json:"status"`
	Tags            StringArray    `gorm:"type:jsonb" json:"tags"`
	MetaDescription string         `gorm:"size:160" json:"metaDescription"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate sets published timestamp for published blogs
func (b *Blog) BeforeCreate(tx *gorm.DB) error {
	if b.Status == BlogStatusPublished && b.PublishedAt == nil {
		now := time.Now()
		b.PublishedAt = &now
	}
	return nil
}

// BeforeUpdate sets published timestamp when status changes to published
func (b *Blog) BeforeUpdate(tx *gorm.DB) error {
	if b.Status == BlogStatusPublished && b.PublishedAt == nil {
		now := time.Now()
		b.PublishedAt = &now
	}
	return nil
}
