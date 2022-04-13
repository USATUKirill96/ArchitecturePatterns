package tests

import (
	"fmt"
	"time"
)

func (tc TestCase) CreateBatches() {
	stmt := `
	   INSERT INTO batch (id, reference, sku, eta, purchased_quantity) 
	   VALUES ($1, $2, $3, $4, $5) 
  RETURNING id
	`
	batches := []struct {
		ID                int
		Reference         string
		SKU               string
		ETA               time.Time
		purchasedQuantity int
	}{
		{1, "test-batch-1", "table", time.Time{}, 10},
		{2, "test-batch-2", "lamp", time.Time{}, 10},
		{3, "test-batch-3", "chair", time.Time{}, 10},
	}

	for _, batch := range batches {
		tc.DB.Exec(stmt, batch.ID, batch.Reference, batch.SKU, batch.ETA, batch.purchasedQuantity)
	}
}

func (tc TestCase) CreateOrderLines() {
	stmt := `
	   INSERT INTO order_line (id, order_id, sku, quantity, batch_id) 
	   VALUES ($1, $2, $3, $4, $5) 
	RETURNING id
	`
	orderLines := []struct {
		ID       int
		OrderID  string
		SKU      string
		Quantity int
		batchID  int
	}{
		{1, "1", "lamp", 3, 2},
		{2, "2", "lamp", 3, 2},
		{3, "3", "chair", 4, 3},
	}

	for _, orderLine := range orderLines {
		tc.DB.Exec(stmt, orderLine.ID, orderLine.OrderID, orderLine.SKU, orderLine.Quantity, orderLine.batchID)
	}
}

func (tc TestCase) Delete() {

	stmt := `
	DELETE FROM batch order_line 
	`
	_, err := tc.DB.Exec(stmt)
	if err != nil {
		fmt.Print("Ошибка при очистке базы", err)
	}
}
