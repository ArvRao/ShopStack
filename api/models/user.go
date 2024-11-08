package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleCustomer UserRole = "customer"
	RoleVendor   UserRole = "vendor"
	RoleAdmin    UserRole = "admin"
)

type UserStatus string

const (
	StatusActive  UserStatus = "active"
	StatusBlocked UserStatus = "blocked"
)

// User represents a user in the e-commerce system
type User struct {
	ID           uint       `gorm:"primaryKey"`
	Name         string     `gorm:"type:varchar(100);not null"`
	Email        string     `gorm:"uniqueIndex;type:varchar(100);not null"`
	PasswordHash string     `gorm:"type:varchar(255);not null"` // Hashed password
	Phone        string     `gorm:"type:varchar(20)"`
	Role         UserRole   `gorm:"type:varchar(20);not null;default:'customer'"`
	Status       UserStatus `gorm:"type:varchar(20);not null;default:'active'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"` // For soft deletes

	// Relationships
	Addresses []Address `gorm:"foreignKey:UserID"` // User addresses
	Orders    []Order   `gorm:"foreignKey:UserID"` // Order history
	Reviews   []Review  `gorm:"foreignKey:UserID"` // Product reviews

	// Additional fields for vendors
	StoreName        *string `gorm:"type:varchar(100)"` // Vendor's store name (nullable for non-vendors)
	StoreDescription *string `gorm:"type:text"`         // Description for vendor's store
}
