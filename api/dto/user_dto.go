package dto

// UserProfileResponse is the response structure for fetching user profile
type UserProfileResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone,omitempty"`
	Role  string `json:"role"`
}

// UpdateUserProfileRequest is the request structure for updating user profile
type UpdateUserProfileRequest struct {
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// ChangePasswordRequest is the request structure for changing password
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}
