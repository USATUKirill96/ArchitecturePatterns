package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"USATUKirill96/AcrhitecturePatterns/internal"
	"USATUKirill96/AcrhitecturePatterns/tests"
)

func TestAllocateBatch(t *testing.T) {
	testCase := tests.NewFakeTestCase()
	batchFixtures := testCase.CreateBatches()
	allocationHandler := internal.AllocationHandler{Container: testCase.Container}

	batch := batchFixtures[0]
	batchOriginQuantity := batch.AvailableQuantity()

	orderLineQuantity := 4

	body := fmt.Sprintf(`{"sku": "%v", "orderid": "123", "quantity": %v}`, batch.SKU, orderLineQuantity)

	req := httptest.NewRequest(http.MethodPost, "/allocate", strings.NewReader(body))
	w := httptest.NewRecorder()
	allocationHandler.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Errorf("Unexpected value in TestAllocateBatch. Expected: %v, got: %v", 200, resp.Status)
	}

	var response struct {
		Message string `json:"message"`
		ID      string `json:"id"`
	}

	err := json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Errorf(err.Error())
	}

	if batch.AvailableQuantity()+orderLineQuantity != batchOriginQuantity {
		t.Errorf(
			"Batch wasn't allocated correctly. Quantity before allocation: %v, after: %v",
			batchOriginQuantity,
			batch.AvailableQuantity(),
		)
	}

	orderLine, _ := testCase.Container.OrderLines.Get(1)
	if orderLine.BatchID != batch.ID {
		t.Errorf("Allocation saved for incorrect batch")
	}

	idFromResponse, _ := strconv.Atoi(response.ID)
	if orderLine.ID != idFromResponse {
		t.Error("ID from response is not the same as in saved orderLine")
	}
}

func TestCantAllocateBatch(t *testing.T) {
	testCase := tests.NewFakeTestCase()
	allocationHandler := internal.AllocationHandler{Container: testCase.Container}
	body := `
	{
		"sku": "something not existing",
		"orderid": "123",
		"quantity": 4,
	}
	`
	req := httptest.NewRequest(http.MethodPost, "/allocate", strings.NewReader(body))
	w := httptest.NewRecorder()
	allocationHandler.ServeHTTP(w, req)
	resp := w.Result()
	if resp.StatusCode != 422 {
		t.Errorf("Unexpected value in TestCantAllocateBatch. Expected: %v, got: %v", 422, resp.Status)
	}
}
