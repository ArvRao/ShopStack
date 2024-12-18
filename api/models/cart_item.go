package models

import (
	"time"

	"gorm.io/gorm"
)

// CartItem represents a product in a user's cart
type CartItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`    // Foreign key to User
	ProductID uint           `gorm:"not null" json:"product_id"` // Foreign key to Product
	Quantity  int            `gorm:"not null" json:"quantity"`   // Quantity of the product
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relationships
	Product Product `gorm:"foreignKey:ProductID" json:"product"` // Linked product
}
