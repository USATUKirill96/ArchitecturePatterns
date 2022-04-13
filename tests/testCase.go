package tests

import (
	"USATUKirill96/AcrhitecturePatterns/pkg/batches"
	"USATUKirill96/AcrhitecturePatterns/utils"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type TestCase struct {
	DB        *sql.DB
	Container *utils.Container
}

func NewTestCase() TestCase {

	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	dsn := os.Getenv("DATABASE_TEST_URL")

	db, err := utils.OpenDB(dsn)
	if err != nil {
		fmt.Println(err)
	}

	batchRepository := batches.NewBatchRepository(db)
	orderLineRepository := batches.NewOrderLineRepository(db)
	container := &utils.Container{
		Batches:    batchRepository,
		OrderLines: orderLineRepository,
	}
	return TestCase{
		DB:        db,
		Container: container,
	}
}
