package app

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRole represents user roles in the system
type UserRole string

const (
	UserRoleAdmin  UserRole = "ADMIN"
	UserRoleEditor UserRole = "EDITOR"
	UserRoleViewer UserRole = "VIEWER"
	UserRoleGuest  UserRole = "GUEST"
)

// Scan implements the Scanner interface for database reads
func (ur *UserRole) Scan(value interface{}) error {
	if value == nil {
		*ur = UserRoleViewer
		return nil
	}
	if s, ok := value.(string); ok {
		*ur = UserRole(s)
		return nil
	}
	return errors.New("cannot scan UserRole")
}

// Value implements the Valuer interface for database writes
func (ur UserRole) Value() (driver.Value, error) {
	return string(ur), nil
}

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
	Password  string         `gorm:"not null;size:255" json:"-"` // Hidden from JSON
	Role      UserRole       `gorm:"not null;default:'VIEWER';type:varchar(20)" json:"role"`
	IsActive  bool           `gorm:"not null;default:true" json:"isActive"`
	LastLogin *time.Time     `json:"lastLogin"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Projects []Project `json:"projects,omitempty"`
}

// BeforeCreate hash password before creating user
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// BeforeUpdate hash password before updating user if password changed
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("Password") && u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// CheckPassword verifies password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// HasRole checks if user has specific role
func (u *User) HasRole(role UserRole) bool {
	return u.Role == role
}

// HasPermission checks if user has permission based on role hierarchy
func (u *User) HasPermission(requiredRole UserRole) bool {
	roleHierarchy := map[UserRole]int{
		UserRoleGuest:  0,
		UserRoleViewer: 1,
		UserRoleEditor: 2,
		UserRoleAdmin:  3,
	}

	userLevel := roleHierarchy[u.Role]
	requiredLevel := roleHierarchy[requiredRole]

	return userLevel >= requiredLevel
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
