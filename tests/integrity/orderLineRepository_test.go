package integrity

import (
	"testing"

	"USATUKirill96/AcrhitecturePatterns/pkg/batches"
	"USATUKirill96/AcrhitecturePatterns/tests"
)

type orderLineExpected struct {
	ID       int
	OrderID  string
	SKU      string
	Quantity int
	BatchID  int
}

func TestOrderLineRepositoryGet(t *testing.T) {
	testCase := tests.NewTestCase()
	teardown := testCase.Setup(t)
	defer teardown(t)
	orderLine, err := testCase.Container.OrderLines.Get(1)
	if err != nil {
		t.Errorf("Error in TestBatchesRepository: %v", err)
	}

	expected := orderLineExpected{1, "1", "lamp", 3, 2}

	assertOrderLine(orderLine, expected, t)
}

func TestOrderLineRepositoryInsert(t *testing.T) {
	testCase := tests.NewTestCase()
	teardown := testCase.Setup(t)
	defer teardown(t)

	expected := orderLineExpected{
		OrderID:  "4444",
		SKU:      "lamp",
		Quantity: 5,
		BatchID:  2,
	}

	orderLine := batches.NewOrderLine(expected.OrderID, expected.SKU, expected.Quantity, expected.BatchID)

	id, err := testCase.Container.OrderLines.Insert(&orderLine)
	if err != nil {
		t.Errorf("Error in TestOrderLineRepository: %v", err)
	}

	savedOrderLine, err := testCase.Container.OrderLines.Get(id)

	if err != nil {
		t.Errorf("Error in TestOrderLineRepository: %v", err)
	}
	expected.ID = id
	assertOrderLine(savedOrderLine, expected, t)

}

func TestOrderLineRepositoryUpdate(t *testing.T) {

	testCase := tests.NewTestCase()
	teardown := testCase.Setup(t)
	defer teardown(t)

	expected := orderLineExpected{
		OrderID:  "4444",
		SKU:      "lamp",
		Quantity: 5,
		BatchID:  3,
	}

	orderLine, err := testCase.Container.OrderLines.Get(1)
	if err != nil {
		panic(err)
	}

	orderLine.OrderID = expected.OrderID
	orderLine.SKU = expected.SKU
	orderLine.Quantity = expected.Quantity
	orderLine.BatchID = expected.BatchID

	err = testCase.Container.OrderLines.Update(orderLine)

	if err != nil {
		t.Errorf("Error in TestOrderLineRepository: %v", err)
	}

	savedOrderLine, err := testCase.Container.OrderLines.Get(1)

	if err != nil {
		t.Errorf("Error in TestOrderLineRepository: %v", err)
	}
	expected.ID = savedOrderLine.ID
	assertOrderLine(savedOrderLine, expected, t)
}

func assertOrderLine(ol *batches.OrderLine, expected orderLineExpected, t *testing.T) {
	if ol.ID != expected.ID {
		t.Errorf("OrderLine ID doesn't match. Expected: %v, got: %v", expected.ID, ol.ID)
	}

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
