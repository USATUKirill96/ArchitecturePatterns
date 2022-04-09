package internal

import (
	"USATUKirill96/AcrhitecturePatterns/pkg/batches"
	"USATUKirill96/AcrhitecturePatterns/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BatchHandler struct {
	Container *utils.Container
}

func (bh BatchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")

		batch, err := bh.Container.Batches.Get(1)
		if err != nil {
			fmt.Println(err)
		}
		jsonResp, err := json.Marshal(batch)
		if err != nil {
			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}

	if r.Method == "POST" {

		var input struct {
			Reference         string `json: "reference"`
			SKU               string `json: "SKU"`
			PurchasedQuantity int    `json: "purchasedQuantity"`
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
		}
		b := batches.NewBatch(input.Reference, input.SKU, time.Now(), input.PurchasedQuantity)
		id, err := bh.Container.Batches.Insert(b)
		if err != nil {
			fmt.Println(err)
		}
		jsonResp, err := json.Marshal(map[string]int{"id": id})
		if err != nil {
			fmt.Println(err)
		}
		w.Write(jsonResp)
		return
	}
}
