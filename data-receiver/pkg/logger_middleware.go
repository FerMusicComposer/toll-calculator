package datareceiver

import (
	"github.com/FerMusicComposer/toll-calculator/models"
	"go.uber.org/zap"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{next: next}
}

func (logMdw *LogMiddleware) ProduceData(data models.OBUData) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("Producing data to Kafka", "data", data)

	return logMdw.next.ProduceData(data)
}
