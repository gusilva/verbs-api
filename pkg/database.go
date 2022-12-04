package pkg

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURL = "mongodb://localhost:27017"
)

func ConnectToMongo() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)

	client, errorClient := mongo.Connect(context.TODO(), clientOptions)
	if errorClient != nil {
		return nil, fmt.Errorf("error connecting to mongodb: %w", errorClient)
	}

	return client.Database("verbs"), nil
}
