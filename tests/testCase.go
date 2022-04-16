package tests

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"USATUKirill96/AcrhitecturePatterns/pkg/batches"
	"USATUKirill96/AcrhitecturePatterns/utils"
)

type TestCase struct {
	db        *sql.DB
	Container *utils.Container
}

func NewIntegrityTestCase() TestCase {

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
		db:        db,
		Container: container,
	}
}

func NewFakeTestCase() TestCase {
	batchRepository := &FakeBatchRepository{}
	orderLineRepository := &FakeOrderLineRepository{}
	container := &utils.Container{
		Batches:    batchRepository,
		OrderLines: orderLineRepository,
	}
	return TestCase{Container: container}
}
