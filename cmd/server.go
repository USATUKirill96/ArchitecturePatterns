package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"

	"USATUKirill96/AcrhitecturePatterns/internal"
	"USATUKirill96/AcrhitecturePatterns/pkg/batches"
	"USATUKirill96/AcrhitecturePatterns/utils"
)

func runserver(c *cli.Context) error {

	variables := utils.ParseVariables()

	db, err := utils.OpenDB(variables.DSN)
	if err != nil {
		fmt.Println(err)
	}

	batchRepository := batches.NewBatchRepository(db)
	orderLineRepository := batches.NewOrderLineRepository(db)
	container := &utils.Container{
		Batches:    batchRepository,
		OrderLines: orderLineRepository,
	}

	r := mux.NewRouter()

	r.Handle("/", internal.AllocationHandler{Container: container})

	http.Handle("/", r)

	srv := &http.Server{
		Addr:    variables.Addr,
		Handler: r,

		IdleTimeout:  time.Minute,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 5,
	}

	err = srv.ListenAndServe()
	fmt.Println(err)
	return nil
}
