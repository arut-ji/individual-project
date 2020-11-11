package main

import (
	"context"
	"fmt"
	"github.com/arut-ji/individual-project/database"
	"github.com/arut-ji/individual-project/linter"
	"github.com/arut-ji/individual-project/sample"
	"github.com/arut-ji/individual-project/util"
	"log"
	"os"
	"sync"
)

func main() {
	ctx := context.Background()
	sampler := sample.NewCodeSampler(ctx)
	ss, err := sampler.NewSampleFromAPI(ctx, &sample.SamplingOptions{
		Size: 50,
	})
	if err != nil {
		_ = fmt.Errorf("%v", err)
		os.Exit(1)
	}

	kubeSamples := make(sample.Samples, 0)
	for _, s := range *ss {
		decodedContent, err := util.DecodeContent(s.Content)
		if err != nil {
			panic(err)
		}
		_, err = linter.Lint(decodedContent)
		if err == nil {
			kubeSamples = append(kubeSamples, s)
		}
	}

	db, err := database.NewDatabase()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		panic(err)
	}

	var wg sync.WaitGroup
	for _, s := range *ss {
		s := s
		go func() {
			wg.Add(1)
			log.Println("Saving: ", s)
			err := db.Create(&s).Error
			if err != nil {
				log.Fatalln("Error creating a sample: ", err)
			}
		}()
	}
	wg.Wait()
}
