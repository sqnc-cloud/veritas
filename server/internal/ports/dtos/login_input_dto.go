package dtos

// LoginInputDTO represents the data transfer object for user login.
type LoginInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
