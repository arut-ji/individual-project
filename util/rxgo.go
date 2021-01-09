package util

import (
	"context"
	"github.com/reactivex/rxgo/v2"
)

type Logger interface {
	Info(i interface{})
}

func WithLogger(logger Logger, f rxgo.Func) rxgo.Func{
	return func(ctx context.Context, i interface{}) (interface{}, error) {
		logger.Info(i)
		return f(ctx, i)
	}
}

func EmptyPipe(_ context.Context, i interface{}) (interface{}, error) {
	return i, nil
}
