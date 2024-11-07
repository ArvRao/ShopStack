package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string
type UserStatus string

const (
	RoleAdmin    UserRole = "admin"
	RoleVendor   UserRole = "vendor"
	RoleCustomer UserRole = "customer"

	StatusActive  UserStatus = "active"
	StatusBlocked UserStatus = "blocked"
)

type User struct {
	ID           uuid.UUID  `json:"id"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"`
	Role         UserRole   `json:"role"`
	Status       UserStatus `json:"status"`
	FullName     string     `json:"full_name"`
	Phone        string     `json:"phone,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type CreateUserInput struct {
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required,min=6"`
	Role     UserRole `json:"role" validate:"required,oneof=admin vendor customer"`
	FullName string   `json:"full_name" validate:"required"`
	Phone    string   `json:"phone" validate:"omitempty"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
