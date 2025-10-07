package dtos

type CreateClaimInputDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateClaimInputDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
