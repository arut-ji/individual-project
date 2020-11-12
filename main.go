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
	sampler := sample.NewCodeSampler(ctx)
	samples, err := sampler.NewSampleFromAPI(ctx, &sample.SamplingOptions{
		Size: 500,
	})
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}

	db, err := database.NewDatabase()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		panic(err)
	}

	var wg sync.WaitGroup
	for _, s := range *samples {
		s := s
		go func() {
			wg.Add(1)
			//log.Println("Saving: ", s)
			err := db.Create(&s).Error
			if err != nil {
				log.Fatalln("Error creating a sample: ", err)
			}
		}()
	}
	wg.Wait()
}
