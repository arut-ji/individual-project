package main

import (
	"context"
	"fmt"
	"github.com/arut-ji/individual-project/database"
	"github.com/arut-ji/individual-project/repo"
	"github.com/arut-ji/individual-project/sample"
	"github.com/arut-ji/individual-project/util"
	"github.com/reactivex/rxgo/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Repository struct {
	Id    int64
	Name  string
	Owner string
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
	baseDirectory := "./repos/"
	<-createMongoSource(ctx, mClient).
		Take(20).
		Map(extractRepository).
		Distinct(util.EmptyPipe).
		ForEach(
			shallowClone(baseDirectory),
			onError,
			onComplete,
			rxgo.WithCPUPool(),
		)
}

func createMongoSource(ctx context.Context, client *mongo.Client) rxgo.Observable {
	ch := make(chan rxgo.Item)
	go sample.LoadSamplesFromMongo(ctx, client, ch)
	return rxgo.FromChannel(ch)
}

func extractRepository(_ context.Context, i interface{}) (interface{}, error) {
	s := i.(sample.Sample)

	return Repository{
		Id:    s.RepositoryId,
		Name:  s.Repository,
		Owner: s.Owner,
	}, nil
}

func shallowClone(destinationPath string) rxgo.NextFunc {
	return func(i interface{}) {
		r := i.(Repository)
		remoteUrl := fmt.Sprintf("https://github.com/%v", r.Name)
		err := repo.ShallowClone(fmt.Sprintf("%v/%v", destinationPath, r.Name), remoteUrl)
		if err != nil {
			panic(err)
		}
	}
}

func onNext(i interface{}) {
	log.Println(fmt.Sprintf("https://github.com/%v", i.(string)))
}

func onError(err error) { panic(err) }

func onComplete() {}
