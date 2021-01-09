package sample

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DatabaseRef   = "kubernetes"
	CollectionRef = "samples"
)

func LoadSamplesFromMongo(ctx context.Context, client *mongo.Client, ch chan rxgo.Item) {
	collection := client.Database(DatabaseRef).Collection(CollectionRef)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		ch <- rxgo.Error(err)
	}
	for cursor.Next(ctx) {
		var item Sample
		if err := cursor.Decode(&item); err != nil {
			ch <- rxgo.Error(err)
		}
		ch <- rxgo.Of(item)
	}
	if err = cursor.Close(ctx); err != nil {
		ch <- rxgo.Error(err)
	}
	close(ch)
}

func CreateMongoSource(ctx context.Context, client *mongo.Client) rxgo.Observable {
	ch := make(chan rxgo.Item)
	go LoadSamplesFromMongo(ctx, client, ch)
	return rxgo.FromChannel(ch)
}
