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

func (logMdw *LoggerMiddleware) CalculateInvoice(obuID int) (invoice *models.Invoice, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	defer func(start time.Time) {
		if invoice != nil {
			sugar.Info("Invoice Calculated",
				zap.Int("obuID", obuID),
				zap.Float64("totalDistance", invoice.TotalDistance),
				zap.Float64("amountDue", invoice.AmountDue),
				zap.Duration("processing_time", time.Since(start)))
		} else {
			sugar.Info("Invoice Calculation Failed: Invoice not received",
				zap.Int("obuID", obuID),
				zap.Duration("processing_time", time.Since(start)))
		}

		if err != nil {
			sugar.Error("Error in Invoice Calculation", zap.Error(err))
		}
	}(time.Now())

	return logMdw.next.CalculateInvoice(obuID)
}
