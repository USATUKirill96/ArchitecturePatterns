package utils

import "USATUKirill96/AcrhitecturePatterns/pkg/batches"

type Container struct {
	Batches    *batches.BatchRepository
	OrderLines *batches.OrderLineRepository
}
