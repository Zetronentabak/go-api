package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Timeout  int
	MongoUri string
	DBname   string
}

func Connect(c MongoConfig) (*mongo.Database, error) {
	clientUrl := c.MongoUri
	clientOptions := options.Client().ApplyURI(clientUrl)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(),
		time.Duration(c.Timeout)*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client.Database(c.DBname), err
}
