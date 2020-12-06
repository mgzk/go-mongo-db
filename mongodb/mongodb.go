package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	uri        string = "mongodb://localhost:27017"
	database   string = "test"
	collection string = "peoples"
)

func Client() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	var err error

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

func Collection(client *mongo.Client) *mongo.Collection {
	return client.Database(database).Collection(collection)
}
