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

func sampleFromGithub(ctx context.Context, ghc *github.Client) rxgo.Observable {
	ch := make(chan rxgo.Item)
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
				log.Println("Error fetching codes: ", err)
				return
			}

			for _, codeResult := range result.CodeResults {
				ch <- rxgo.Of(codeResult)
			}
			time.Sleep(time.Second * 1)
		}
		close(ch)
	}(ch)
	return rxgo.FromEventSource(
		ch,
		rxgo.WithContext(ctx),
		rxgo.WithPublishStrategy(),
	)
}

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

func (s *sampler) fetchCommitCounts(ctx context.Context) {

}
