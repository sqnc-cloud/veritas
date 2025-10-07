package usecases

import (
	"context"
	"fmt"
	"time"
	"veritas/core/domain"
	"veritas/internal/ports/output"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClaimUsecase struct {
	repo output.ClaimOutputPort
}

func NewClaimUsecase(repo output.ClaimOutputPort) *ClaimUsecase {
	return &ClaimUsecase{repo: repo}
}

type CreateClaimInput struct {
	Name        string
	Description string
}

func (uc *ClaimUsecase) CreateClaim(ctx context.Context, input CreateClaimInput) (primitive.ObjectID, error) {
	claim := &domain.Claim{
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return uc.repo.CreateClaim(ctx, claim)
}

func (uc *ClaimUsecase) ReadClaim(ctx context.Context, id string) (*domain.Claim, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}
	return uc.repo.GetClaim(ctx, objectID)
}

type UpdateClaimInput struct {
	Name        string
	Description string
}

func (uc *ClaimUsecase) UpdateClaim(ctx context.Context, id string, input UpdateClaimInput) (*domain.Claim, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	existingClaim, err := uc.repo.GetClaim(ctx, objectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get claim: %w", err)
	}

	if input.Name != "" {
		existingClaim.Name = input.Name
	}
	if input.Description != "" {
		existingClaim.Description = input.Description
	}

	existingClaim.UpdatedAt = time.Now()

	err = uc.repo.UpdateClaim(ctx, objectID, existingClaim)
	if err != nil {
		return nil, err
	}

	return existingClaim, nil
}

func (uc *ClaimUsecase) DeleteClaim(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}
	return uc.repo.DeleteClaim(ctx, objectID)
}

func (uc *ClaimUsecase) GetAllClaims(ctx context.Context) ([]*domain.Claim, error) {
	return uc.repo.GetAllClaims(ctx)
}
