package dtos

type CreateRoleInputDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateRoleInputDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
