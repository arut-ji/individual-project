package main

import (
	"context"
	"github.com/arut-ji/individual-project/database"
	"github.com/arut-ji/individual-project/linter/smells_detector"
	"github.com/arut-ji/individual-project/sample"
	"github.com/arut-ji/individual-project/util"
	"github.com/reactivex/rxgo/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type DetectionResult struct {
	Content string `json:"content,omitempty" bson:"content,omitempty"`
}

func main() {
	ctx := context.Background()
	mClient, mClose, err := database.NewMongoClient(ctx, "mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := mClose()
		if err != nil {
			panic(err)
		}
	}()

	// Create a mongo DB data sink.
	mongoResultSink := createMongoSink(mClient, "detections")
	// Pull kubernetes scripts from a mongo's collection named "samples"
	<-createMongoSource(ctx, mClient, "samples").
		Take(5).
		Map(decodeContent). // Decode base64 content into string
		Map(
			detectImplementationSmells, // Feed each content to smells detection pipeline
		).
		Map(mongoResultSink). // Save the result into a mongo's collection named "detections"
		Run()
}

func decodeContent(_ context.Context, i interface{}) (interface{}, error) {
	s := i.(sample.Sample)
	content, err := util.DecodeContent(s.Content)
	if err != nil {
		return nil, err
	}
	return string(content), nil
}

func detectImplementationSmells(_ context.Context, i interface{}) (interface{}, error) {
	detectionResult, err := smells_detector.Detect(i.(string))
	if err != nil {
		return nil, err
	}
	return detectionResult, nil
}

func createMongoSource(ctx context.Context, client *mongo.Client, collectionName string) rxgo.Observable {
	ch := make(chan rxgo.Item)
	collection := client.Database("kubernetes").Collection(collectionName)
	go loadScriptFromMongo(ctx, collection, ch)
	return rxgo.FromChannel(ch)
}

func loadScriptFromMongo(ctx context.Context, collection *mongo.Collection, ch chan rxgo.Item) {
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		ch <- rxgo.Error(err)
	}
	for cursor.Next(ctx) {
		var item sample.Sample
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

func createMongoSink(client *mongo.Client, collectionName string) rxgo.Func {

	collection := client.Database("kubernetes").Collection(collectionName)

	return func(ctx context.Context, item interface{}) (interface{}, error) {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		s := item.(smells_detector.DetectionResult)
		return collection.InsertOne(ctx, s)
	}
}
