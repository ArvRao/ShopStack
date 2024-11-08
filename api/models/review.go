package models

import (
	"time"
)

type Review struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`                                            // Foreign key to User
	ProductID uint   `gorm:"not null"`                                            // Foreign key to Product
	Rating    int    `gorm:"type:int;not null;check:rating >= 1 AND rating <= 5"` // Rating from 1 to 5
	Comment   string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
