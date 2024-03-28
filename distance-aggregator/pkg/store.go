package distanceaggregator

import (
	"fmt"

	"github.com/FerMusicComposer/toll-calculator/models"
)

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (ms *MemoryStore) Insert(data models.Distance) error {
	ms.data[data.OBUID] += data.Value
	return nil
}

func (ms *MemoryStore) Get(id int) (float64, error) {
	dist, ok := ms.data[id]
	if !ok {
		return 0, fmt.Errorf("could not find records for id: %d", id)
	}

	return dist, nil
}
