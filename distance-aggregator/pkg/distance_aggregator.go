package distanceaggregator

import (
	"fmt"

	"github.com/FerMusicComposer/toll-calculator/models"
)

const basePrice = 3.15

type Aggregator interface {
	AggregateDistance(models.Distance) error
	CalculateInvoice(int) (*models.Invoice, error)
}

type Storer interface {
	Insert(models.Distance) error
	Get(int) (float64, error)
}

type DistanceAggregator struct {
	store Storer
}

func NewDistanceAggregator(store Storer) Aggregator {
	return &DistanceAggregator{store: store}
}

func (da *DistanceAggregator) AggregateDistance(data models.Distance) error {
	fmt.Printf("Inserting data: %+v\n", data)
	return da.store.Insert(data)
}

func (da *DistanceAggregator) CalculateInvoice(obuID int) (*models.Invoice, error) {
	dist, err := da.store.Get(obuID)
	if err != nil {
		return nil, fmt.Errorf("could not find records for id: %d", obuID)
	}

	inv := &models.Invoice{
		OBUID:         obuID,
		TotalDistance: dist,
		AmountDue:     basePrice * dist,
	}

	return inv, nil
}
