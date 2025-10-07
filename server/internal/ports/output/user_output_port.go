package output

import (
	"context"
	"veritas/core/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserOutputPort interface {
	CreateUser(ctx context.Context, user *domain.User) (primitive.ObjectID, error)
	GetUser(ctx context.Context, id primitive.ObjectID) (*domain.User, error)
	UpdateUser(ctx context.Context, id primitive.ObjectID, user *domain.User) error
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
}
