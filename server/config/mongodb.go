package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDBClient(ctx context.Context) (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
		log.Println("MONGO_URI not set, using default: mongodb://localhost:27017")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}
	return client, nil
}

func GetDatabaseName() string {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "housekeeper"
		log.Println("DB_NAME not set, using default: housekeeper")
	}
	return dbName
}
