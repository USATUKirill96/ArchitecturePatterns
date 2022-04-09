package batches

import (
	"testing"
	"time"
)

func makeBatchAndLine(sku string, batchQuantity, lineQuantity int) (*Batch, OrderLine) {
	batch := NewBatch("batch-001", sku, time.Now(), batchQuantity)
	line := NewOrderLine("order-123", sku, lineQuantity)

	return batch, line
}

func TestAllocatingToABatchReducesTheAvailableQuantity(t *testing.T) {
	batch, line := makeBatchAndLine("SMALL-TABLE", 20, 2)

	batch.Allocate(line)

	if batch.AvailableQuantity() != 18 {
		t.Errorf("Expected %v, got %v", 18, batch.AvailableQuantity())
	}

}

func TestCanAllocate(t *testing.T) {

	testCases := []struct {
		batchQuantity int
		lineQuantity  int
		expected      bool
	}{
		{20, 2, true},
		{2, 20, false},
		{2, 2, true},
	}

	for _, testCase := range testCases {
		batch, line := makeBatchAndLine("ELEGANT-LAMP", testCase.batchQuantity, testCase.lineQuantity)
		result := batch.CanAllocate(line)
		if result != testCase.expected {
			t.Errorf(
				"batchQTY: %v, lineQTY: %v, expected: %v, got: %v",
				batch.AvailableQuantity(), line.Quantity, testCase.expected, result,
			)
		}
	}
}

func TestCannotAllocateIfSQUsDoNotMatch(t *testing.T) {
	batch := NewBatch("batch-001", "CHAIT", time.Time{}, 100)
	line := NewOrderLine("order-123", "TOASTER", 10)
	if batch.CanAllocate(line) != false {
		t.Error("Expected False, got True")
	}
}

func TestCanOnlyDeallocateAllocatedLines(t *testing.T) {
	batch, unallocatedLine := makeBatchAndLine("DECORATIVE-TRINKET", 20, 2)
	batch.Deallocate(unallocatedLine)
	if batch.AvailableQuantity() != 20 {
		t.Errorf("Expected 20, got %v", batch.AvailableQuantity())
	}
}

func TestAllocationIsIdempotent(t *testing.T) {
	batch, line := makeBatchAndLine("ANGULAR-DESK", 20, 2)
	batch.Allocate(line)
	batch.Allocate(line)
	if batch.AvailableQuantity() != 18 {
		t.Errorf("Expected 18, got %v", batch.AvailableQuantity())
	}
}
