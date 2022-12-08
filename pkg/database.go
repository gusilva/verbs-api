package pkg

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"time"
)

type MongoDatabase struct {
	DSN    string
	Logger *zap.Logger
	client *mongo.Client
}

func (db *MongoDatabase) ConnectToMongo() error {
	clientOptions := options.Client().ApplyURI(db.DSN)

	client, errorClient := mongo.Connect(context.TODO(), clientOptions)
	if errorClient != nil {
		return fmt.Errorf("error connecting to mongodb: %w", errorClient)
	}

	db.client = client

	return nil
}

func (db *MongoDatabase) PingDB() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if errorPingDB := db.client.Ping(ctx, readpref.Primary()); errorPingDB != nil {
		db.Logger.Warn("db unavailable", zap.String("error", errorPingDB.Error()))
	} else {
		db.Logger.Debug("database connected successfully")
	}
}

func (db *MongoDatabase) GetCollection(collName string) (*mongo.Collection, error) {
	if db.client == nil {
		return nil, fmt.Errorf("no database client set")
	}

	db.Logger.Debug("getting collection", zap.String("name", collName))

	return db.client.Database("verbs").Collection(collName), nil
}
