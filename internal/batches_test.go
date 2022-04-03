package internal

import (
	"testing"
	"time"
)

func TestAllocatingToABatchReducesTheAvailableQuantity(t *testing.T) {
	batch := &Batch{"batch-001", "SMALL-TABLE", 20, time.Now()}
	line := OrderLine{"order-ref", "SMALL-TABLE", 2}

	batch.Allocate(line)

	if batch.AvailableQuantity != 18 {
		t.Errorf("Error in AllocatingToABatchReducesTheAvailableQuantity: expected %v, got %v", 18, batch.AvailableQuantity)
	}

}
