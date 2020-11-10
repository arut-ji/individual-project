package sample

import (
	"context"
	"github.com/arut-ji/individual-project/database"
	_ "github.com/jinzhu/gorm"
)

// TODO: Caching implementation
func (s *sampler) NewSampleFromCache(ctx context.Context, opts *SamplingOptions) (*Samples, error) {
	panic("implement me")
}

func (s *sampler) saveToCache(samples *Samples) error {
	for _, elem := range *samples {
		err := s.db.Create(&database.Sample{
			FileName:     elem.FileName,
			Path:         elem.Path,
			Repository:   elem.Repository,
			RepositoryId: elem.RepositoryId,
			Fork:         elem.Fork,
			Content:      elem.Content,
		}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
