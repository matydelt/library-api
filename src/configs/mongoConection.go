package configs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// Connect initializes a connection to MongoDB
func Connect(uri string, dbName string) *mongo.Database {
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle to the database
	db = client.Database(dbName)

	return db
}

// GetDB returns a handle to the database
func GetDB() *mongo.Database {
	return db
}
