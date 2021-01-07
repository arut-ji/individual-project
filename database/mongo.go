package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CloseClientFn func() error

func closeFn(ctx context.Context, client *mongo.Client) CloseClientFn {
	return func() error {
		return client.Disconnect(ctx)
	}
}

func NewMongoClient(ctx context.Context, uri string) (*mongo.Client, CloseClientFn, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, err
	}

	return client, closeFn(ctx, client), nil
}
