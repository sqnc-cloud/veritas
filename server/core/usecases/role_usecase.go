package usecases

import (
	"context"
	"fmt"
	"time"
	"veritas/core/domain"
	"veritas/internal/ports/output"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoleUsecase struct {
	repo output.RoleOutputPort
}

func NewRoleUsecase(repo output.RoleOutputPort) *RoleUsecase {
	return &RoleUsecase{repo: repo}
}

type CreateRoleInput struct {
	Name        string
	Description string
}

func (uc *RoleUsecase) CreateRole(ctx context.Context, input CreateRoleInput) (primitive.ObjectID, error) {
	role := &domain.Role{
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return uc.repo.CreateRole(ctx, role)
}

func (uc *RoleUsecase) ReadRole(ctx context.Context, id string) (*domain.Role, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}
	return uc.repo.GetRole(ctx, objectID)
}

type UpdateRoleInput struct {
	Name        string
	Description string
}

func (uc *RoleUsecase) UpdateRole(ctx context.Context, id string, input UpdateRoleInput) (*domain.Role, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	existingRole, err := uc.repo.GetRole(ctx, objectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get role: %w", err)
	}

	if input.Name != "" {
		existingRole.Name = input.Name
	}
	if input.Description != "" {
		existingRole.Description = input.Description
	}

	existingRole.UpdatedAt = time.Now()

	err = uc.repo.UpdateRole(ctx, objectID, existingRole)
	if err != nil {
		return nil, err
	}

	return existingRole, nil
}

func (uc *RoleUsecase) DeleteRole(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}
	return uc.repo.DeleteRole(ctx, objectID)
}

func (uc *RoleUsecase) GetAllRoles(ctx context.Context) ([]*domain.Role, error) {
	return uc.repo.GetAllRoles(ctx)
}
