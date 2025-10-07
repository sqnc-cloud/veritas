package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateRoleOutputDTO struct {
	ID primitive.ObjectID `json:"id"`
}

type UpdateRoleOutputDTO struct {
	ID primitive.ObjectID `json:"id"`
}
