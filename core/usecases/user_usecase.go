package usecases

import (
	"context"
	"fmt"
	"time"
	"veritas/core/domain"
	"veritas/internal/ports/output"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	repo output.UserOutputPort
}

func NewUserUsecase(repo output.UserOutputPort) *UserUsecase {
	return &UserUsecase{repo: repo}
}

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

func (uc *UserUsecase) CreateUser(ctx context.Context, input CreateUserInput) (primitive.ObjectID, error) {
	user := &domain.User{
		Username:  input.Name,
		Email:     input.Email,
		Password:  input.Password, // TODO: hash and salt this!
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUsecase) ReadUser(ctx context.Context, id string) (*domain.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}
	return uc.repo.GetUser(ctx, objectID)
}

type UpdateUserInput struct {
	Name     string
	Email    string
	Password string
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, id string, input UpdateUserInput) (*domain.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	existingUser, err := uc.repo.GetUser(ctx, objectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if input.Name != "" {
		existingUser.Username = input.Name
	}
	if input.Email != "" {
		existingUser.Email = input.Email
	}
	if input.Password != "" {
		existingUser.Password = input.Password // TODO: hash and salt this!
	}
	existingUser.UpdatedAt = time.Now()

	err = uc.repo.UpdateUser(ctx, objectID, existingUser)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}
	return uc.repo.DeleteUser(ctx, objectID)
}

func (uc *UserUsecase) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	return uc.repo.GetAllUsers(ctx)
}

func (uc *UserUsecase) VerifyUser(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := uc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if user.Password != password { // TODO: Compare hashed passwords
		return nil, fmt.Errorf("invalid credentials")
	}

	return user, nil
}
