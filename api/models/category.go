package models

import (
	"time"

	"gorm.io/gorm"
)

// Category represents the category of products
type Category struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"unique;not null" json:"name"`       // Category name must be unique
	Description string         `json:"description"`                       // Optional description
	CreatedAt   time.Time      `json:"created_at"`                        // Timestamp for creation
	UpdatedAt   time.Time      `json:"updated_at"`                        // Timestamp for updates
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"` // Soft delete field
}
