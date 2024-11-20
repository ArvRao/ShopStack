package services

import (
	"errors"

	"github.com/ArvRao/shopstack/api/dto"
	"github.com/ArvRao/shopstack/api/models"
	"github.com/ArvRao/shopstack/api/utils"
	"gorm.io/gorm"
)

// UserService struct to encapsulate user-related methods
type UserService struct {
	db *gorm.DB // Database connection (optional, can be passed as dependency)
}

// NewUserService creates a new instance of UserService
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// GetUserProfile retrieves the user's profile based on their user ID
func (s *UserService) GetUserProfile(userID uint) (*dto.UserProfileResponse, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}

	profile := &dto.UserProfileResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Role:  string(user.Role),
	}
	return profile, nil
}

// UpdateUserProfile updates a user's profile based on provided data
func (s *UserService) UpdateUserProfile(userID uint, updateData *dto.UpdateUserProfileRequest) error {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	if updateData.Name != nil {
		user.Name = *updateData.Name
	}
	if updateData.Phone != nil {
		user.Phone = *updateData.Phone
	}

	if err := s.db.Save(&user).Error; err != nil {
		return errors.New("failed to update user profile")
	}
	return nil
}

// ChangePassword changes the user's password after verifying the current password
func (s *UserService) ChangePassword(userID uint, currentPassword, newPassword string) error {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	// Verify current password
	if !utils.CheckPasswordHash(currentPassword, user.PasswordHash) {
		return errors.New("current password is incorrect")
	}

	// Hash the new password and update it
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("failed to hash new password")
	}
	user.PasswordHash = hashedPassword

	if err := s.db.Save(&user).Error; err != nil {
		return errors.New("failed to update password")
	}
	return nil
}
