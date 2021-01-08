package sample

import (
	"context"
	"errors"
	"github.com/arut-ji/individual-project/linter"
	"github.com/google/go-github/v32/github"
	"github.com/reactivex/rxgo/v2"
	"log"
	"time"
)

const (
	KubernetesQueryString = "apiVersion+in:file language:YAML"
)

func (s *sampler) NewSampleFromAPI(ctx context.Context, opts *SamplingOptions) (*Samples, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ob := sampleFromGithub(ctx, s.ghc).
		Map(mapToContent(s.ghc), rxgo.WithCPUPool()).
		Map(mapWithLintingResult).
		Take(uint(opts.Size))

	ob.Connect(ctx)

	samples := make(Samples, 0)

	var err error = nil

	for item := range ob.Observe() {
		sample := item.V.(Sample)
		log.Println("Getting a file from: ", sample.RepositoryId)
		samples = append(samples, sample)
		err = s.save(&sample)
		if err != nil {
			return nil, err
		}
	}

	return &samples, nil
}

// This function returns an observable emitting query results from Github API.
func sampleFromGithub(ctx context.Context, ghc *github.Client) rxgo.Observable {
	ch := make(chan rxgo.Item)
	/*
		Github API returns one single page results with a next page token.
		The API has a known limitation: the maximum number of item returned is 1000 items.
		In this case, we request 100 samples per page then iterate over the pages
		until it reach the defined upper bound.
	*/
	perPage := 100
	go func(ch chan rxgo.Item) {
		for page := 0; perPage*page <= 1000; page++ {
			log.Printf("Fetching %v page ...", page)
			result, _, err := ghc.Search.Code(
				ctx,
				KubernetesQueryString,
				&github.SearchOptions{
					ListOptions: github.ListOptions{
						PerPage: perPage,
						Page:    page,
					},
				})
			if err != nil {
				// log.Println is used because using either "log.Fatal"  or "log.Panic" is going to panic.
				log.Println("Error fetching codes: ", err)
				ch <- rxgo.Error(err)
			}
			// Feed each code result into the observable stream.
			for _, codeResult := range result.CodeResults {
				ch <- rxgo.Of(codeResult)
			}
			// Impose time interval to prevent fraud detection in Github API.
			time.Sleep(time.Second * 1)
		}
		close(ch)
	}(ch)
	return rxgo.FromChannel(
		ch,
		rxgo.WithContext(ctx),
		rxgo.WithPublishStrategy(),
	)
}

// Fetch file contents from each query result.
func fetchContent(ctx context.Context, ghc *github.Client, owner, repo, path string) (*string, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	fileContent, _, _, err := ghc.Repositories.GetContents(ctx, owner, repo, path, nil)
	if err != nil && fileContent != nil {
		return nil, err
	}
	if fileContent != nil {
		return fileContent.Content, nil
	}
	return nil, errors.New("file content not found")
}

// Map each query result to the content fetched
func mapToContent(ghc *github.Client) rxgo.Func {
	return func(ctx context.Context, item interface{}) (interface{}, error) {
		result := item.(*github.CodeResult)
		repo := result.GetRepository()
		content, err := fetchContent(
			ctx,
			ghc,
			repo.GetOwner().GetLogin(),
			repo.GetName(),
			result.GetPath(),
		)
		if err != nil {
			return nil, err
		}
		return Sample{
			FileName:     result.GetName(),
			Path:         result.GetPath(),
			Repository:   result.GetRepository().GetFullName(),
			RepositoryId: result.GetRepository().GetID(),
			Owner:        repo.GetOwner().GetLogin(),
			Fork:         result.GetRepository().GetFork(),
			Content:      *content,
		}, nil
	}
}

func mapWithLintingResult(_ context.Context, item interface{}) (interface{}, error) {
	sample := item.(Sample)
	lintingResult, err := linter.IsKubernetesScriptValid(sample.Content)
	if err != nil {
		return nil, err
	}
	return Sample{
		FileName:      sample.FileName,
		Path:          sample.Path,
		Repository:    sample.Repository,
		RepositoryId:  sample.RepositoryId,
		Fork:          false,
		LintingResult: lintingResult,
		Content:       sample.Content,
	}, nil
}

func fetchCommitCount(ctx context.Context, ghc *github.Client, owner, repo string) (int, error) {
	commits, _, err := ghc.Repositories.ListCommits(ctx, owner, repo, nil)
	if err != nil {
		return 1, err
	}
	return len(commits), nil
}
