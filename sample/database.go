package sample

import "context"

func (s *sampler) NewSampleFromDB(ctx context.Context, opts *SamplingOptions) (*Samples, error) {
	panic("implement me")
}
