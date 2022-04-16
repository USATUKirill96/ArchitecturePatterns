package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"USATUKirill96/AcrhitecturePatterns/pkg/batches"
	"USATUKirill96/AcrhitecturePatterns/pkg/services"
	"USATUKirill96/AcrhitecturePatterns/utils"
)

type AllocationHandler struct {
	Container *utils.Container
}

func (h AllocationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "GET" {
	// 	w.WriteHeader(http.StatusCreated)
	// 	w.Header().Set("Content-Type", "application/json")

	// 	batch, err := ah.Container.Batches.Get(1)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	jsonResp, err := json.Marshal(batch)
	// 	if err != nil {
	// 		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	// 	}
	// 	w.Write(jsonResp)
	// 	return
	// }

	if r.Method == "POST" {

		var input struct {
			OrderID  string `json:"orderid"`
			SKU      string `json:"sku"`
			Quantity int    `json:"quantity"`
		}

		w.Header().Set("Content-Type", "application/json")

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			fmt.Println(err)
		}

		orderLine := batches.NewOrderLine(input.OrderID, input.SKU, input.Quantity, 0)

		batches, err := h.Container.Batches.FliterBySKU(input.SKU)
		if err != nil {
			h.Container.Logger.ServerError(err, w)
			return
		}

		batch, err := services.Allocate(orderLine, batches)
		if err != nil {
			w.WriteHeader(422)
			jsonResp, _ := json.Marshal(map[string]string{"error": err.Error()})
			w.Write(jsonResp)
			return
		}

		orderLine.BatchID = batch.ID

		id, err := h.Container.OrderLines.Insert(orderLine)
		if err != nil {
			h.Container.Logger.ServerError(err, w)
			return
		}

		jsonResp, err := json.Marshal(map[string]string{"message": "order allocated", "id": fmt.Sprint(id)})
		if err != nil {
			h.Container.Logger.ServerError(err, w)
			return
		}

		w.Write(jsonResp)
	}
}
