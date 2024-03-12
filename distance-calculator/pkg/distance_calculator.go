package distancecalculator

import (
	"math"

	"github.com/FerMusicComposer/toll-calculator/models"
)

type DistCalculator interface {
	CalculateDistance(models.OBUData) (float64, error)
}

type DistanceCalculator struct {
	points [][]float64
}

func NewDistanceCalculator() DistCalculator {
	return &DistanceCalculator{
		points: make([][]float64, 0),
	}
}

func (dc *DistanceCalculator) CalculateDistance(data models.OBUData) (float64, error) {
	distance := 0.0

	if len(dc.points) > 0 {
		prevPoint := dc.points[len(dc.points)-1]
		distance = calculateDistance(prevPoint[0], prevPoint[1], data.Latitude, data.Longitude)
	}

	dc.points = append(dc.points, []float64{data.Latitude, data.Longitude})

	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
