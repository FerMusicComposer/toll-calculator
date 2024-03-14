package distanceaggregator

import (
	"fmt"

	"github.com/FerMusicComposer/toll-calculator/models"
)

type Aggregator interface {
	AggregateDistance(models.Distance) error
}

type Storer interface {
	Insert(models.Distance) error
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
