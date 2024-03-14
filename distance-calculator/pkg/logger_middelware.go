package distancecalculator

import (
	"time"

	"github.com/FerMusicComposer/toll-calculator/models"
	"go.uber.org/zap"
)

type LoggerMiddlewware struct {
	next DistCalculator
}

func NewLoggerMiddleware(next DistCalculator) DistCalculator {
	return &LoggerMiddlewware{next: next}
}

func (logMdw *LoggerMiddlewware) CalculateDistance(data models.OBUData) (dist float64, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	defer func(start time.Time) {
		sugar.Info("Distance Calculated",
			zap.Float64p("distance", &dist),
			zap.Duration("processing_time", time.Since(start)),
			zap.Error(err))
	}(time.Now())

	return logMdw.next.CalculateDistance(data)
}
