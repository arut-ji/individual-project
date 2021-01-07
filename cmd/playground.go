package cmd

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewPlaygroundCmd(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "playground",
		Short: "Execute code implemented in scratch pad.",
		Long:  "Execute code implemented in scratch pad.",
		Run: func(cmd *cobra.Command, args []string) {
			mongoPlayground(ctx)
		},
	}
}

func mongoPlayground(ctx context.Context) {
	// Create Mongo Client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// Connect Mongo Client to a Mongo instance
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	// Deferred connection closing
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	err = insertItem(ctx, client)
	if err != nil {
		panic(err)
	}

}

func insertItem(ctx context.Context, client *mongo.Client) error {
	collection := client.Database("testing").Collection("numbers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		return err
	}
	return nil
}
