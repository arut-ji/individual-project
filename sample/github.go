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

	sch := s.createSource(ctx)

	ob := rxgo.
		FromChannel(sch, rxgo.WithPublishStrategy()).
		Map(mapToContent(s), rxgo.WithCPUPool()).
		Map(mapWithLintingResult).
		Take(uint(opts.Size))

	ob.Connect(ctx)

	samples := make(Samples, 0)

	for item := range ob.Observe() {
		sample := item.V.(Sample)
		log.Println("Getting a file from: ", sample.RepositoryId)
		samples = append(samples, sample)
	}

	return &samples, nil
}

func (s *sampler) createSource(ctx context.Context) <-chan rxgo.Item {

	ch := make(chan rxgo.Item)

	perPage := 100

	go func(och chan rxgo.Item) {
		for page := 0; perPage*page <= 1000; page++ {
			log.Printf("Fetching %v page ...", page)
			result, _, err := s.ghc.Search.Code(
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
				och <- rxgo.Of(codeResult)
			}
			time.Sleep(time.Second * 1)
		}
	}(ch)
	return ch
}

func (s *sampler) fetchContent(ctx context.Context, owner, repo, path string) (*string, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	fileContent, _, _, err := s.ghc.Repositories.GetContents(ctx, owner, repo, path, nil)
	if err != nil && fileContent != nil {
		return nil, err
	}
	if fileContent != nil {
		return fileContent.Content, nil
	}
	return nil, errors.New("file content not found")
}

func mapToContent(s *sampler) rxgo.Func {
	return func(ctx context.Context, item interface{}) (interface{}, error) {
		result := item.(*github.CodeResult)
		repo := result.GetRepository()
		content, err := s.fetchContent(
			ctx,
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
