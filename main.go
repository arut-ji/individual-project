package main

import (
	"context"
	"fmt"
	"github.com/arut-ji/individual-project/sample"
	"os"
)

func main() {
	ctx := context.Background()
	sampler := sample.NewCodeSampler(ctx)
	result, err := sampler.NewSampleFromAPI(ctx, &sample.SamplingOptions{
		Size: 2,
	})
	if err != nil {
		_ = fmt.Errorf("%v", err)
		os.Exit(1)
	}

}
