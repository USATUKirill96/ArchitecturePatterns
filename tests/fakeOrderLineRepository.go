package tests

import (
	"database/sql"

	"USATUKirill96/AcrhitecturePatterns/pkg/batches"
)

type FakeOrderLineRepository struct {
	OrderLines      []*batches.OrderLine
	lastOrderLineId int
}

func (r *FakeOrderLineRepository) Get(id int) (*batches.OrderLine, error) {
	for _, ol := range r.OrderLines {
		if ol.ID == id {
			return ol, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (r *FakeOrderLineRepository) Insert(ol *batches.OrderLine) (int, error) {
	ol.ID = r.lastOrderLineId + 1
	r.lastOrderLineId += 1
	r.OrderLines = append(r.OrderLines, ol)
	return ol.ID, nil
}

func (r *FakeOrderLineRepository) Update(orderLine *batches.OrderLine) error {

	idx := 0
	for i, ol := range r.OrderLines {
		if ol.ID == orderLine.ID {
			idx = i
			break
		}
	}

	if idx == 0 {
		return sql.ErrNoRows
	}

	r.OrderLines[idx] = orderLine
	return nil
}
