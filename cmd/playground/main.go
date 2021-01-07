package main

import (
	"context"
	"github.com/arut-ji/individual-project/database"
	"github.com/arut-ji/individual-project/sample"
	"github.com/jinzhu/gorm"
	"github.com/reactivex/rxgo/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

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

	mongoSink := createMongoSink(mClient)

	<-sampleFromSqlite().
		Map(mongoSink).
		Run()
}

func sampleFromSqlite() rxgo.Observable {
	ch := make(chan rxgo.Item)
	go func(ch chan rxgo.Item) {
		db, err := gorm.Open("sqlite3", "samples.db")
		if err != nil {
			ch <- rxgo.Error(err)
		}
		db.AutoMigrate(&database.Sample{})
		ghrp := sample.NewRepository(db)
		samples, err := ghrp.GetAll()
		if err != nil {
			ch <- rxgo.Error(err)
		}
		for _, s := range *samples {
			ch <- rxgo.Item{
				V: s,
			}
		}
		close(ch)
	}(ch)
	return rxgo.FromEventSource(ch, rxgo.WithPublishStrategy())
}

func loadScripts(db *gorm.DB) (*sample.Samples, error) {
	results := make(sample.Samples, 0)
	err := db.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func createMongoSink(client *mongo.Client) rxgo.Func {

	collection := client.Database("kubernetes").Collection("samples")

	return func(ctx context.Context, item interface{}) (interface{}, error) {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		s := item.(sample.Sample)
		return collection.InsertOne(ctx, s)
	}
}
