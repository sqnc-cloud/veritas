package db

import (
	"context"
	"fmt"
	"time"
	"veritas/core/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const roleCollectionName = "roles"

type RoleRepository struct {
	db *mongo.Database
}

func NewRoleRepository(db *mongo.Database) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) CreateRole(ctx context.Context, role *domain.Role) (primitive.ObjectID, error) {
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()

	result, err := r.db.Collection(roleCollectionName).InsertOne(ctx, role)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("failed to insert role: %w", err)
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.ObjectID{}, fmt.Errorf("failed to convert inserted id to ObjectID")
	}

	return objectID, nil
}

func (r *RoleRepository) GetRole(ctx context.Context, id primitive.ObjectID) (*domain.Role, error) {
	var role domain.Role
	filter := bson.M{"_id": id}

	err := r.db.Collection(roleCollectionName).FindOne(ctx, filter).Decode(&role)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("role not found")
		}
		return nil, fmt.Errorf("failed to get role: %w", err)
	}

	return &role, nil
}

func (r *RoleRepository) GetRoleByName(ctx context.Context, name string) (*domain.Role, error) {
	var role domain.Role
	filter := bson.M{"name": name}

	err := r.db.Collection(roleCollectionName).FindOne(ctx, filter).Decode(&role)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("role not found")
		}
		return nil, fmt.Errorf("failed to get role: %w", err)
	}

	return &role, nil
}

func (r *RoleRepository) UpdateRole(ctx context.Context, id primitive.ObjectID, role *domain.Role) error {
	role.UpdatedAt = time.Now()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": role}

	_, err := r.db.Collection(roleCollectionName).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update role: %w", err)
	}

	return nil
}

func (r *RoleRepository) DeleteRole(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := r.db.Collection(roleCollectionName).DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete role: %w", err)
	}

	return nil
}

func (r *RoleRepository) GetAllRoles(ctx context.Context) ([]*domain.Role, error) {
	var roles []*domain.Role

	cursor, err := r.db.Collection(roleCollectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get all roles: %w", err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &roles); err != nil {
		return nil, fmt.Errorf("failed to decode roles: %w", err)
	}

	return roles, nil
}
