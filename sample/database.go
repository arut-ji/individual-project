package sample

import (
	"context"
)

func (s *sampler) NewSampleFromDB(ctx context.Context, opts *SamplingOptions) (*Samples, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	samples := make(Samples, 0)
	limit := opts.Size
	err := s.db.
		Find(&samples).
		Error
	if err != nil {
		return nil, err
	}
	samples = samples[:limit]
	return &samples, nil

}

func (s *sampler) save(sample *Sample) error {
	err := s.db.Create(&sample).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *sampler) loadValidScripts() (*Samples, error) {
	results := make(Samples, 0)
	err := s.db.Where("lintingResult <> ?", true).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return &results, nil
}
