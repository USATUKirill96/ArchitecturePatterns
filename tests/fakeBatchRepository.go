package tests

import (
	"database/sql"

	"USATUKirill96/AcrhitecturePatterns/pkg/batches"
)

type FakeBatchRepository struct {
	Batches     []*batches.Batch
	lastBatchID int
}

func (r *FakeBatchRepository) Get(id int) (*batches.Batch, error) {
	for _, b := range r.Batches {
		if b.ID == id {
			return b, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (r *FakeBatchRepository) FliterBySKU(sku string) ([]*batches.Batch, error) {
	var batches []*batches.Batch
	for _, batch := range r.Batches {
		if batch.SKU == sku {
			batches = append(batches, batch)
		}
	}
	return batches, nil
}

func (r *FakeBatchRepository) LoadAllocations(b *batches.Batch) error {
	return nil
}

func (r *FakeBatchRepository) Insert(b *batches.Batch) (int, error) {
	b.ID = r.lastBatchID + 1
	r.lastBatchID += 1
	r.Batches = append(r.Batches, b)
	return b.ID, nil
}

func (r *FakeBatchRepository) Update(b *batches.Batch) error {
	idx := 0
	for i, batch := range r.Batches {
		if batch.ID == b.ID {
			idx = i
			break
		}
	}

	if idx == 0 {
		return sql.ErrNoRows
	}

	r.Batches[idx] = b
	return nil
}
