package main

import (
	"context"
	"fmt"
	"github.com/arut-ji/individual-project/database"
	"github.com/arut-ji/individual-project/sample"
	"log"
)

func main() {
	ctx := context.Background()

	db, err := database.NewDatabase()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		panic(err)
	}

	sampler := sample.NewCodeSampler(ctx, db)
	//_, err = sampler.NewSampleFromAPI(ctx, &sample.SamplingOptions{
	//	Size: 1000,
	//})
	//if err != nil {
	//	_ = fmt.Errorf("%v", err)
	//}
	samples, err := sampler.NewSampleFromDB(ctx, &sample.SamplingOptions{
		Size: 10,
	})
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}

}
