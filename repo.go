package main

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoClient(ctx context.Context) (*mongo.Client, error) {
	mongoURL := viper.GetString("mongo.prod")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getUsersCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(viper.GetString("mongo.db")).Collection("users")
}
