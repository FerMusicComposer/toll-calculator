package distanceaggregator

import "github.com/FerMusicComposer/toll-calculator/models"

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
