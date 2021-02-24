package main

import (
	"context"
	"fmt"
	"github.com/arut-ji/individual-project/database"
	"github.com/arut-ji/individual-project/linter/smells_detector"
	"github.com/arut-ji/individual-project/sample"
	"github.com/arut-ji/individual-project/util"
	"github.com/reactivex/rxgo/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

type DetectionRecord struct {
	FileName          string                          `json:"fileName,omitempty" bson:"fileName,omitempty"`
	Path              string                          `json:"path,omitempty" bson:"path,omitempty"`
	RepositoryId      int64                           `json:"repositoryId,omitempty" bson:"repositoryId,omitempty"`
	DetectionResult   smells_detector.DetectionResult `json:"detectionResult,omitempty" bson:"detectionResult, omitempty"`
	LineOfCodes       int                             `json:"lineOfCodes,omitempty" bson:"lineOfCodes,omitempty"`
	NumberOfResources int                             `json:"numberOfResources,omitempty" bson:"numberOfResources,omitempty"`
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
		Map(decodeContent). // Decode base64 content into string
		Map(
			detectImplementationSmells, // Feed each content to smells detection pipeline
		).
		Map(mongoResultSink). // Save the result into a mongo's collection named "detections"
		DoOnError(func(err error) {
			fmt.Println(err)
		})
}

func decodeContent(_ context.Context, i interface{}) (interface{}, error) {
	s := i.(sample.Sample)
	content, err := util.DecodeContent(s.Content)
	if err != nil {
		return nil, err
	}
	s.Content = string(content)
	return s, nil
}

func detectImplementationSmells(_ context.Context, i interface{}) (interface{}, error) {
	s := i.(sample.Sample)
	lineOfCodes := countLineOfCodes(s.Content)
	numberOfResources := util.GetNumberOfResources(s.Content)
	fmt.Printf("Repository Id: %v\nLine of Codes: %v\nNumber of resources: %v\n\n", s.RepositoryId, lineOfCodes, numberOfResources)
	detectionResult, err := smells_detector.Detect(s.Content)
	if err != nil {
		return nil, err
	}
	return DetectionRecord{
		FileName:          s.FileName,
		Path:              s.Path,
		RepositoryId:      s.RepositoryId,
		DetectionResult:   detectionResult,
		LineOfCodes:       lineOfCodes,
		NumberOfResources: numberOfResources,
	}, nil
}

func countLineOfCodes(script string) int {
	return len(strings.Split(script, "\n"))
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
		s := item.(DetectionRecord)
		opts := options.FindOneAndUpdate().SetUpsert(true)
		filter := bson.D{
			{"fileName", s.FileName},
			{"path", s.Path},
			{"repositoryId", s.RepositoryId},
		}
		update := bson.D{
			{"$set", bson.D{{"detectionResult", s.DetectionResult}}},
			{"$set", bson.D{{"lineOfCodes", s.LineOfCodes}}},
			{"$set", bson.D{{"numberOfResources", s.NumberOfResources}}},
		}
		var updatedDocument bson.M
		err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedDocument)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return updatedDocument, nil
			}
			return nil, err
		}
		return updatedDocument, nil
	}
}
