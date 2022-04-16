package tests

import (
	"fmt"
	"time"

	"USATUKirill96/AcrhitecturePatterns/pkg/batches"
)

func (tc TestCase) CreateBatches() []*batches.Batch {
	batches := []*batches.Batch{
		batches.NewBatch("test-batch-1", "table", time.Time{}, 10),
		batches.NewBatch("test-batch-2", "lamp", time.Time{}, 10),
		batches.NewBatch("test-batch-3", "chair", time.Time{}, 10),
	}

	for _, batch := range batches {
		id, _ := tc.Container.Batches.Insert(batch)
		batch.ID = id
	}

	return batches

}

func (tc TestCase) CreateOrderLines(bathList []*batches.Batch) []*batches.OrderLine {
	orderLines := []*batches.OrderLine{
		batches.NewOrderLine("1", "lamp", 3, bathList[1].ID),
		batches.NewOrderLine("2", "lamp", 3, bathList[1].ID),
		batches.NewOrderLine("3", "chair", 4, bathList[2].ID),
	}

	for _, orderLine := range orderLines {
		id, err := tc.Container.OrderLines.Insert(orderLine)
		if err != nil {
			fmt.Println("Error during allocations creation, ", err)
		}
		orderLine.ID = id
	}

	return orderLines
}

func (tc TestCase) Delete() {

	if tc.db != nil {
		stmt := `
	DELETE FROM batch order_line 
	`
		_, err := tc.db.Exec(stmt)
		if err != nil {
			fmt.Print("Error during flushing database ", err)
		}
	} else {
		tc.Container.Batches = &FakeBatchRepository{}
		tc.Container.OrderLines = &FakeOrderLineRepository{}
	}
}
