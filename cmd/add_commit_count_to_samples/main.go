package main

import (
	"context"
	"github.com/arut-ji/individual-project/database"
	"github.com/arut-ji/individual-project/sample"
	"github.com/google/go-github/v32/github"
	"github.com/imdario/mergo"
	"github.com/reactivex/rxgo/v2"
	"os"
	"strings"
)

func main() {
	ctx := context.Background()
	mClient, mClose, err := database.NewMongoClient(ctx, "mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	ghClient := sample.NewGithubClient(ctx, os.Getenv("GITHUB_TOKEN"))

	<-sample.CreateMongoSource(ctx, mClient).
		Take(1).
		Map(mapWithOwner).
		Map(mapWithCommitCount(ghClient)).
		Run()

	err = mClose()
	if err != nil {
		panic(err)
	}
}

func mapWithOwner(_ context.Context, i interface{}) (interface{}, error) {
	s := i.(sample.Sample)
	patch := sample.Sample{
		Owner: strings.Split(s.Repository, "/")[0],
	}

	if err := mergo.Merge(&s, patch); err != nil {
		return nil, err
	}
	return s, nil
}

func mapWithCommitCount(ghc *github.Client) rxgo.Func {
	return func(ctx context.Context, i interface{}) (interface{}, error) {
		s := i.(sample.Sample)
		commitCount, err := sample.FetchCommitCount(ctx, ghc, s.Owner, s.Repository)
		if err != nil {
			return nil, err
		}
		patch := sample.Sample{
			CommitCount: int64(commitCount),
		}
		if err := mergo.Merge(&s, patch); err != nil {
			return nil, err
		}

		return s, nil
	}
}
