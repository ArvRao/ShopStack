package models

import (
	"time"
)

type OrderStatus string

const (
	OrderPending   OrderStatus = "pending"
	OrderCompleted OrderStatus = "completed"
	OrderCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID         uint        `gorm:"primaryKey"`
	UserID     uint        `gorm:"not null"`         // Foreign key to User
	Status     OrderStatus `gorm:"type:varchar(20)"` // Order status
	TotalPrice float64     `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
