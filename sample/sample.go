package sample

import (
	"context"
	"github.com/google/go-github/v32/github"
	"github.com/jinzhu/gorm"
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
}

func NewCodeSampler(ctx context.Context, db *gorm.DB) Sampler {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	return &sampler{
		ghc: github.NewClient(tc),
		db:  db,
	}
}
