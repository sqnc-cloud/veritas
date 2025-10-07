package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateClaimOutputDTO struct {
	ID primitive.ObjectID `json:"id"`
}

type UpdateClaimOutputDTO struct {
	ID primitive.ObjectID `json:"id"`
}
