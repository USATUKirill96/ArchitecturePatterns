package integrity

import (
	"testing"

	"USATUKirill96/AcrhitecturePatterns/pkg/batches"
	"USATUKirill96/AcrhitecturePatterns/tests"
)

type orderLineExpected struct {
	OrderID  string
	SKU      string
	Quantity int
	BatchID  int
}

func TestOrderLineRepositoryGet(t *testing.T) {
	testCase := tests.NewIntegrityTestCase()
	batchFixtures := testCase.CreateBatches()
	orderLineFixtures := testCase.CreateOrderLines(batchFixtures)
	t.Cleanup(testCase.Delete)

	orderLine, err := testCase.Container.OrderLines.Get(orderLineFixtures[0].ID)
	if err != nil {
		t.Errorf("Error in TestBatchesRepository: %v", err)
	}

	expected := orderLineExpected{"1", "lamp", 3, batchFixtures[1].ID}

	assertOrderLine(orderLine, expected, t)
}

func TestOrderLineRepositoryInsert(t *testing.T) {
	testCase := tests.NewIntegrityTestCase()
	batchFixtures := testCase.CreateBatches()
	t.Cleanup(testCase.Delete)

	expected := orderLineExpected{
		OrderID:  "4444",
		SKU:      "lamp",
		Quantity: 5,
		BatchID:  batchFixtures[1].ID,
	}

	orderLine := batches.NewOrderLine(expected.OrderID, expected.SKU, expected.Quantity, expected.BatchID)

	id, err := testCase.Container.OrderLines.Insert(orderLine)
	if err != nil {
		t.Errorf("Error in TestOrderLineRepository: %v", err)
		return
	}

	savedOrderLine, err := testCase.Container.OrderLines.Get(id)

	if err != nil {
		t.Errorf("Error in TestOrderLineRepository: %v", err)
		return
	}
	assertOrderLine(savedOrderLine, expected, t)

}

func TestOrderLineRepositoryUpdate(t *testing.T) {

	testCase := tests.NewIntegrityTestCase()
	batchFixtures := testCase.CreateBatches()
	orderLineFixtures := testCase.CreateOrderLines(batchFixtures)
	t.Cleanup(testCase.Delete)

	expected := orderLineExpected{
		OrderID:  "4444",
		SKU:      "lamp",
		Quantity: 5,
		BatchID:  batchFixtures[2].ID,
	}

	orderLine := orderLineFixtures[0]

	orderLine.OrderID = expected.OrderID
	orderLine.SKU = expected.SKU
	orderLine.Quantity = expected.Quantity
	orderLine.BatchID = expected.BatchID

	err := testCase.Container.OrderLines.Update(orderLine)

	if err != nil {
		t.Errorf("Error in TestOrderLineRepository: %v", err)
		return
	}

	savedOrderLine, err := testCase.Container.OrderLines.Get(orderLine.ID)

	if err != nil {
		t.Errorf("Error in TestOrderLineRepository: %v", err)
		return
	}

	assertOrderLine(savedOrderLine, expected, t)
}

func assertOrderLine(ol *batches.OrderLine, expected orderLineExpected, t *testing.T) {
	if ol.OrderID != expected.OrderID {
		t.Errorf("OrderLine OrderID doesn't match. Expected: %v, got: %v", expected.OrderID, ol.OrderID)
	}
	if ol.SKU != expected.SKU {
		t.Errorf("OrderLine SKU doesn't match. Expected: %v, got: %v", expected.SKU, ol.SKU)
	}
	if ol.Quantity != expected.Quantity {
		t.Errorf("OrderLine Quantity doesn't match. Expected: %v, got: %v", expected.Quantity, ol.Quantity)
	}
	if ol.BatchID != expected.BatchID {
		t.Errorf("OrderLine BatchID doesn't match. Expected: %v, got: %v", expected.BatchID, ol.BatchID)
	}
}
