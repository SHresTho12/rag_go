package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect establishes a connection to the MongoDB database.
func Connect(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)

	// Create a new context for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database to check if connection is established
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("could not connect to the MongoDB server: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

// NewDatabase returns a connection to a specific MongoDB database.
func NewDatabase(client *mongo.Client, dbName string) *mongo.Database {
	return client.Database(dbName)
}
