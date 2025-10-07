package output

import (
	"context"
	"veritas/core/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoleOutputPort interface {
	CreateRole(ctx context.Context, role *domain.Role) (primitive.ObjectID, error)
	GetRole(ctx context.Context, id primitive.ObjectID) (*domain.Role, error)
	UpdateRole(ctx context.Context, id primitive.ObjectID, role *domain.Role) error
	DeleteRole(ctx context.Context, id primitive.ObjectID) error
	GetRoleByName(ctx context.Context, name string) (*domain.Role, error)
	GetAllRoles(ctx context.Context) ([]*domain.Role, error)
}
