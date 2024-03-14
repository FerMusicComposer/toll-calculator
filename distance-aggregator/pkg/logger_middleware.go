package distanceaggregator

import (
	"time"

	"github.com/FerMusicComposer/toll-calculator/models"
	"go.uber.org/zap"
)

type LoggerMiddleware struct {
	next Aggregator
}

func NewLoggerMiddleware(next Aggregator) Aggregator {
	return &LoggerMiddleware{next: next}
}

func (logMdw *LoggerMiddleware) AggregateDistance(data models.Distance) (err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	defer func(start time.Time) {
		sugar.Info("Distance Aggregator",
			zap.Duration("processing_time", time.Since(start)))
		zap.Error(err)
	}(time.Now())

	return logMdw.next.AggregateDistance(data)
}
