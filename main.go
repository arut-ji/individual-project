package main

import (
	"context"
	"fmt"
	"github.com/arut-ji/individual-project/database"
	"github.com/arut-ji/individual-project/sample"
	"log"
	"sync"
)

func main() {
	ctx := context.Background()

	db, err := database.NewDatabase()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		panic(err)
	}

	sampler := sample.NewCodeSampler(ctx)
	samples, err := sampler.NewSampleFromAPI(ctx, &sample.SamplingOptions{
		Size: 1000,
	})
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}

	var wg sync.WaitGroup
	for _, s := range *samples {
		err := db.Create(&s).Error
		if err != nil {
			log.Fatalln("Error creating a sample: ", err)
		}
	}
	wg.Wait()
}
