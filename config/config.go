package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDatabase *mongo.Database

// InitMongoDB initializes the MongoDB connection
func InitMongoDB() {
	mongoURI := os.Getenv("MONGOSTRING")
	if mongoURI == "" {
		log.Fatal("MONGOSTRING environment variable is not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Error creating MongoDB client: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB: ", err)
	}

	MongoClient = client
	MongoDatabase = client.Database("Ruteangkot")
	log.Println("Connected to MongoDB!")
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatal("MongoClient is not initialized")
	}
	return MongoDatabase.Collection(collectionName)
}
