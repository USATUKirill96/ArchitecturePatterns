package services

import (
	"testing"

	"USATUKirill96/AcrhitecturePatterns/pkg/batches"
	"USATUKirill96/AcrhitecturePatterns/tests"
)

func TestAllocate(t *testing.T) {
	testCase := tests.NewFakeTestCase()
	batchFixtures := testCase.CreateBatches()
	orderLineFixtures := testCase.CreateOrderLines(batchFixtures)
	orderLineFixtures = append(orderLineFixtures, batches.NewOrderLine("4", "incorrect", 4, 0))

	batchAvailables := make(map[int]int)
	for _, fixture := range batchFixtures {
		batchAvailables[fixture.ID] = fixture.AvailableQuantity()
	}

	for _, orderLine := range orderLineFixtures {
		batch, err := Allocate(orderLine, batchFixtures)
		if orderLine.SKU == "incorrect" {
			if err == nil {
				t.Errorf("Allocation had no matching batches, but no error returning")
			}
			return
		}

		if batch.AvailableQuantity()+orderLine.Quantity != batchAvailables[batch.ID] {
			t.Errorf("Incorrect calculation of batches left after Allocation being processed: Expected %v, got %v",
				batch.AvailableQuantity()+orderLine.Quantity, batchAvailables[batch.ID])
		}
		batchAvailables[batch.ID] = batch.AvailableQuantity()

	}
}
