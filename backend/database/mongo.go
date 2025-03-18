package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Connect initializes the MongoDB connection
func Connect(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the database to ensure connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Could not ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")
	Client = client
}

// GetCollection returns a collection from the database
func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return Client.Database(databaseName).Collection(collectionName)
}

// Disconnect closes the MongoDB connection
func Disconnect() {
	if err := Client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Error disconnecting MongoDB: %v", err)
	}
	log.Println("MongoDB connection closed")
}
