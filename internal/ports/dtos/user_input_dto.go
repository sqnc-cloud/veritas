package dtos

type CreateUserInputDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdateUserInputDTO struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty,email"`
	Password string `json:"password,omitempty,min=8"`
}
