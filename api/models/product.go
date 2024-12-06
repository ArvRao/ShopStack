package models

import (
	"time"

	"gorm.io/gorm"
)

// Product represents a product
type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"not null" json:"user_id"`           // Foreign key to User table
	CategoryID  uint           `gorm:"not null" json:"category_id"`       // Foreign key to Category table
	Name        string         `gorm:"not null" json:"name"`              // Product name
	Description string         `json:"description"`                       // Optional description
	Price       float64        `gorm:"not null" json:"price"`             // Product price
	Stock       int            `gorm:"not null" json:"stock"`             // Quantity available in stock
	CreatedAt   time.Time      `json:"created_at"`                        // Timestamp for creation
	UpdatedAt   time.Time      `json:"updated_at"`                        // Timestamp for updates
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"` // Soft delete field

	// Relationships
	Category Category `gorm:"foreignKey:CategoryID" json:"category"` // Linked category
	User     User     `gorm:"foreignKey:UserID" json:"user"`         // Linked vendor/admin
}
