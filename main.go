package main

import (
	"context"
	"fmt"
	"github.com/arut-ji/individual-project/linter"
	"github.com/arut-ji/individual-project/sample"
	"github.com/arut-ji/individual-project/util"
	"os"
)

func main() {
	ctx := context.Background()
	sampler := sample.NewCodeSampler(ctx)
	ss, err := sampler.NewSampleFromAPI(ctx, &sample.SamplingOptions{
		Size: 2,
	})
	if err != nil {
		_ = fmt.Errorf("%v", err)
		os.Exit(1)
	}
	for _, s := range *ss {
		decodedContent, err := util.DecodeContent(s.Content)
		if err != nil {
			panic(err)
		}
		_, err = linter.Lint(decodedContent)
	}
}
