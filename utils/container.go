package utils

import "USATUKirill96/AcrhitecturePatterns/pkg/batches"

type iBatchRepository interface {
	Get(int) (*batches.Batch, error)
	FliterBySKU(string) ([]*batches.Batch, error)
	LoadAllocations(*batches.Batch) error
	Insert(*batches.Batch) (int, error)
	Update(*batches.Batch) error
}

type iOrderLineRepository interface {
	Get(int) (*batches.OrderLine, error)
	Insert(*batches.OrderLine) (int, error)
	Update(*batches.OrderLine) error
}

type Container struct {
	Batches    iBatchRepository
	OrderLines iOrderLineRepository
	Logger     *Logger
}
