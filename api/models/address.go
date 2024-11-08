package models

import (
	"time"
)

type Address struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"` // Foreign key to User
	Line1     string `gorm:"type:varchar(255);not null"`
	Line2     string `gorm:"type:varchar(255)"`
	City      string `gorm:"type:varchar(100);not null"`
	State     string `gorm:"type:varchar(100)"`
	Country   string `gorm:"type:varchar(100);not null"`
	ZipCode   string `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
