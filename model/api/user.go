package api

type UserCreateOrUpdateRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
}
