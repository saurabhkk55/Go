package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}

func main() {
	client, err := connectMongoDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Disconnect(context.TODO())

	// Use the client for MongoDB operations...
	// Example: collection := client.Database("your_database").Collection("your_collection")
	// Perform CRUD operations on the collection...
}
