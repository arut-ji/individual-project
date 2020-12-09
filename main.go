package main

import "github.com/arut-ji/individual-project/cmd"

func main() {
	cmd.Execute()
}

//func main() {
//	ctx := context.Background()
//
//	db, err := database.NewDatabase()
//	if err != nil {
//		_ = fmt.Errorf("%v", err)
//		panic(err)
//	}
//
//	sampler := sample.NewCodeSampler(ctx, db)
//	//_, err = sampler.NewSampleFromAPI(ctx, &sample.SamplingOptions{
//	//	Size: 1000,
//	//})
//	//if err != nil {
//	//	_ = fmt.Errorf("%v", err)
//	//}
//	samples, err := sampler.NewSampleFromDB(ctx, &sample.SamplingOptions{
//		Size: 10,
//	})
//
//	if err != nil {
//		_ = fmt.Errorf("%v", err)
//	}
//	jsonEncodedSamples, _ := json.MarshalIndent(samples, "", " ")
//
//	err = ioutil.WriteFile("www/samples/samples.json", jsonEncodedSamples, 0644)
//
//	if err != nil {
//		panic(err)
//	}
//
//}
