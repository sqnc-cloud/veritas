package input

import (
	"context"
	"veritas/core/domain"
	"veritas/core/usecases"
	"veritas/internal/ports/dtos"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInputPort interface {
	CreateUser(ctx context.Context, input dtos.CreateUserInputDTO) (primitive.ObjectID, error)
	GetUser(ctx context.Context, id string) (*domain.User, error)
	UpdateUser(ctx context.Context, id string, input usecases.UpdateUserInput) (*domain.User, error)
	DeleteUser(ctx context.Context, id string) error
}
