package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := os.Getenv("MONGODB_URL")
	fmt.Println("MongoDB URL:", MongoDb)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		fmt.Println("Error creating MongoDB client:", err)
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return nil
	}

	fmt.Println("Connected to MongoDB")
	return client
}

var Client *mongo.Client = DBinstance()

func openCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant_management").Collection(collectionName)
	return collection
}