package services

import (
	"errors"

	"USATUKirill96/AcrhitecturePatterns/pkg/batches"
)

func Allocate(orderline *batches.OrderLine, batches []*batches.Batch) (*batches.Batch, error) {
	for _, batch := range batches {
		if batch.CanAllocate(orderline) {
			batch.Allocate(orderline)
			return batch, nil
		}
	}
	return nil, errors.New("Can't allocate batch")
}
