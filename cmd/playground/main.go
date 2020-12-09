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
)

func main() {

	ctx := context.Background()

	db, err := database.NewDatabase()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		panic(err)
	}
	log.Println("Creating a new sampler ...")
	sampler := sample.NewCodeSampler(ctx, db)
	samples, err := sampler.NewSampleFromDB(ctx, &sample.SamplingOptions{
		Size: 1000,
	})

	if err != nil {
		_ = fmt.Errorf("%v", err)
	}

	log.Println("Decoding contents ...")

	scriptsWithError := make([]sample.Sample, 0)

	for _, s := range *samples {
		content, err := util.DecodeContent(s.Content)
		if err != nil {
			os.Exit(1)
		}
		results, err := linter.Lint(content)

		numErrors := 0

		for _, result := range results {
			//fmt.Printf("Resource kind: %v\n", result.Kind)
			numErrors += len(result.Errors)
		}
		if numErrors != 0 {
			scriptsWithError = append(scriptsWithError, s)
			fmt.Printf("Filename: %v\n", s.FileName)
			fmt.Printf("Path: %v\n", s.Path)
			fmt.Println("Repository name: " + s.Repository)
			fmt.Printf("Number of errors: %d\n\n", numErrors)
		}
	}
	fmt.Println("Number of samples:", len(*samples))
	fmt.Println("Number of scripts containing errors:", len(scriptsWithError))
}
