package sample

import (
	"context"
)

// TODO: Caching implementation
func (s *sampler) NewSampleFromCache(ctx context.Context, opts *SamplingOptions) (*Samples, error) {
	panic("implement me")
}
