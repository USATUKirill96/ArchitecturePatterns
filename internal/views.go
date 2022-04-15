package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"USATUKirill96/AcrhitecturePatterns/utils"
)

type AllocationHandler struct {
	Container *utils.Container
}

func (ah AllocationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
			OrderID  string `json: "orderid"`
			SKU      string `json: "sku"`
			Quantity int    `json: "quantity"`
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")

		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(r.PostForm)
		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			fmt.Println(err)

			jsonResp, err := json.Marshal(map[string]int{"id": 1})
			if err != nil {
				fmt.Println(err)
			}
			w.Write(jsonResp)
			return
		}
	}
}
