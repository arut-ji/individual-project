package sample

import (
	"context"
	"errors"
	"github.com/google/go-github/v32/github"
	"sync"
)

const (
	KubernetesQueryString = "apiVersion+in:file language:YAML"
)

func (s *sampler) NewSampleFromAPI(ctx context.Context, opts *SamplingOptions) (*Samples, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// TODO: Implement a mechanism to use pagination feature when the sample size >= 50

	result, _, err := s.ghc.Search.Code(
		ctx,
		KubernetesQueryString,
		&github.SearchOptions{
			ListOptions: github.ListOptions{
				PerPage: (int)(opts.Size),
			},
		})
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	samples := make(Samples, 0)
	var mux sync.Mutex
	for _, file := range result.CodeResults {
		wg.Add(1)
		cch := make(chan string, 1)
		go func() {
			defer wg.Done()
			var repo = file.GetRepository()
			content, err := s.fetchContent(
				ctx,
				repo.GetOwner().GetLogin(),
				repo.GetName(),
				file.GetPath(),
			)
			if err != nil {
				panic(err)
			}
			if content != nil {
				cch <- *content
			}
		}()
		mux.Lock()
		samples = append(samples, Sample{
			FileName:     file.GetName(),
			Path:         file.GetPath(),
			Repository:   file.GetRepository().GetFullName(),
			RepositoryId: file.GetRepository().GetID(),
			Fork:         file.GetRepository().GetFork(),
			Content:      <-cch,
		})
		mux.Unlock()
	}

	wg.Wait()

	return &samples, nil
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
