package output

import (
	"context"
	"veritas/core/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClaimOutputPort interface {
	CreateClaim(ctx context.Context, claim *domain.Claim) (primitive.ObjectID, error)
	GetClaim(ctx context.Context, id primitive.ObjectID) (*domain.Claim, error)
	UpdateClaim(ctx context.Context, id primitive.ObjectID, claim *domain.Claim) error
	DeleteClaim(ctx context.Context, id primitive.ObjectID) error
	GetClaimByName(ctx context.Context, name string) (*domain.Claim, error)
	GetAllClaims(ctx context.Context) ([]*domain.Claim, error)
}
