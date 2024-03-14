package distancecalculator

import (
	"math"

	"github.com/FerMusicComposer/toll-calculator/models"
)

type DistCalculator interface {
	CalculateDistance(models.OBUData) (float64, error)
}

type DistanceCalculator struct {
	prevpoint []float64
}

func NewDistanceCalculator() DistCalculator {
	return &DistanceCalculator{}
}

func (dc *DistanceCalculator) CalculateDistance(data models.OBUData) (float64, error) {
	distance := 0.0

	if len(dc.prevpoint) > 0 {
		distance = calculateDistance(dc.prevpoint[0], dc.prevpoint[1], data.Latitude, data.Longitude)
	}

	dc.prevpoint = []float64{data.Latitude, data.Longitude}

	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
