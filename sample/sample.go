package sample

import (
	"context"
	"github.com/google/go-github/v32/github"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
	"os"
	"sync"
)

type SamplingOptions struct {
	Size int32
	Seed int32
}

type Sample CodeContent
type Samples []Sample

type Sampler interface {
	NewSampleFromAPI(ctx context.Context, opts *SamplingOptions) (*Samples, error)
	NewSampleFromDB(ctx context.Context, opts *SamplingOptions) (*Samples, error)
}

type sampler struct {
	ghc *github.Client
	db  *gorm.DB
	mu  *sync.Mutex
	mc  *mongo.Client
}

func NewCodeSampler(ctx context.Context, db *gorm.DB, mc *mongo.Client) Sampler {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	return &sampler{
		ghc: github.NewClient(tc),
		db:  db,
		mc:  mc,
	}
}
