package integrity

import (
	"testing"
	"time"

	"USATUKirill96/AcrhitecturePatterns/pkg/batches"
	"USATUKirill96/AcrhitecturePatterns/tests"
)

func TestBatchesRepositoryCanGetBatch(t *testing.T) {
	testCase := tests.NewTestCase()
	teardown := testCase.Setup(t)
	defer teardown(t)
	batch, err := testCase.Container.Batches.Get(1)
	if err != nil {
		t.Errorf("Error in TestBatchesRepository: %v", err)
	}

	expected := expectedType{"test-batch-1", "table", time.Time{}, 10}

	assertBatch(batch, expected, t)

}

func TestBatchesRepositoryCanCreateBatch(t *testing.T) {

	testCase := tests.NewTestCase()
	teardown := testCase.Setup(t)
	defer teardown(t)

	reference := "Test-batch-created"
	sku := "tested-good"
	eta := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	purchasedAuantity := 10

	batch := batches.NewBatch(reference, sku, eta, purchasedAuantity)

	id, err := testCase.Container.Batches.Insert(batch)
	if err != nil {
		t.Error(err)
		return
	}

	createdBatch, err := testCase.Container.Batches.Get(id)

	if err != nil {
		t.Error(err)
		return
	}
	expected := expectedType{reference: reference, sku: sku, eta: eta, availableQuantity: purchasedAuantity}

	assertBatch(createdBatch, expected, t)
}

func TestBatchRepositoryCanUpdateBatch(t *testing.T) {

	testCase := tests.NewTestCase()
	teardown := testCase.Setup(t)
	defer teardown(t)

	reference := "Test-batch-created"
	sku := "tested-good"
	eta := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	purchasedAuantity := 15

	batch := batches.NewBatch(reference, sku, eta, purchasedAuantity)

	batch.ID = 1

	err := testCase.Container.Batches.Update(batch)
	if err != nil {
		t.Error(err)
		return
	}

	updatedBatch, err := testCase.Container.Batches.Get(batch.ID)

	if err != nil {
		t.Error(err)
		return
	}

	expected := expectedType{reference: reference, sku: sku, eta: eta, availableQuantity: purchasedAuantity}

	assertBatch(updatedBatch, expected, t)

}

type expectedType struct {
	reference         string
	sku               string
	eta               time.Time
	availableQuantity int
}

func assertBatch(batch *batches.Batch, expected expectedType, t *testing.T) {

	if batch == nil {
		t.Error("Batch was not created in fixtures")
	}
	if batch.Reference != expected.reference {
		t.Errorf("Batch reference doesn't match. Expected: %v, got: %v", expected.reference, batch.Reference)
	}

	if batch.SKU != expected.sku {
		t.Errorf("Batch SKU doesn't match. Expected: %v, got: %v", expected.sku, batch.SKU)
	}
	if batch.ETA != expected.eta {
		t.Errorf("Batch ETA doesn't match. Expected: %v, got: %v", expected.eta, batch.ETA)
	}

	if batch.AvailableQuantity() != expected.availableQuantity {
		t.Errorf(
			"Batch available quantity doesn't match. Expected: %v, got: %v",
			expected.availableQuantity,
			batch.AvailableQuantity(),
		)
	}
}
