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

const claimCollectionName = "claims"

type ClaimRepository struct {
	db *mongo.Database
}

func NewClaimRepository(db *mongo.Database) *ClaimRepository {
	return &ClaimRepository{db: db}
}

func (r *ClaimRepository) CreateClaim(ctx context.Context, claim *domain.Claim) (primitive.ObjectID, error) {
	claim.CreatedAt = time.Now()
	claim.UpdatedAt = time.Now()

	result, err := r.db.Collection(claimCollectionName).InsertOne(ctx, claim)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("failed to insert claim: %w", err)
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.ObjectID{}, fmt.Errorf("failed to convert inserted id to ObjectID")
	}

	return objectID, nil
}

func (r *ClaimRepository) GetClaim(ctx context.Context, id primitive.ObjectID) (*domain.Claim, error) {
	var claim domain.Claim
	filter := bson.M{"_id": id}

	err := r.db.Collection(claimCollectionName).FindOne(ctx, filter).Decode(&claim)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("claim not found")
		}
		return nil, fmt.Errorf("failed to get claim: %w", err)
	}

	return &claim, nil
}

func (r *ClaimRepository) GetClaimByName(ctx context.Context, name string) (*domain.Claim, error) {
	var claim domain.Claim
	filter := bson.M{"name": name}

	err := r.db.Collection(claimCollectionName).FindOne(ctx, filter).Decode(&claim)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("claim not found")
		}
		return nil, fmt.Errorf("failed to get claim: %w", err)
	}

	return &claim, nil
}

func (r *ClaimRepository) UpdateClaim(ctx context.Context, id primitive.ObjectID, claim *domain.Claim) error {
	claim.UpdatedAt = time.Now()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": claim}

	_, err := r.db.Collection(claimCollectionName).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update claim: %w", err)
	}

	return nil
}

func (r *ClaimRepository) DeleteClaim(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := r.db.Collection(claimCollectionName).DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete claim: %w", err)
	}

	return nil
}

func (r *ClaimRepository) GetAllClaims(ctx context.Context) ([]*domain.Claim, error) {
	var claims []*domain.Claim

	cursor, err := r.db.Collection(claimCollectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get all claims: %w", err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &claims); err != nil {
		return nil, fmt.Errorf("failed to decode claims: %w", err)
	}

	return claims, nil
}
