package cmd

import (
	"context"
	"fmt"
	"github.com/arut-ji/individual-project/database"
	"github.com/arut-ji/individual-project/linter"
	"github.com/arut-ji/individual-project/sample"
	"github.com/arut-ji/individual-project/util"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func NewPlaygroundCmd(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "playground",
		Short: "Execute code implemented in scratch pad.",
		Long:  "Execute code implemented in scratch pad.",
		Run: func(cmd *cobra.Command, args []string) {
			playground(ctx)
		},
	}
}

func playground(ctx context.Context) {
	db, err := database.NewDatabase()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		panic(err)
	}
	log.Println("Creating a new sampler ...")
	sampler := sample.NewCodeSampler(ctx, db)
	samples, err := sampler.NewSampleFromDB(ctx, &sample.SamplingOptions{
		Size: 500,
	})

	if err != nil {
		_ = fmt.Errorf("%v", err)
	}

	log.Println("Decoding contents ...")

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
			fmt.Printf("Filename: %v\n", s.FileName)
			fmt.Printf("Path: %v\n", s.Path)
			fmt.Println("Repository name: " + s.Repository)
			fmt.Printf("Number of errors: %d\n\n", numErrors)
		}
	}
}
