package internal

import (
	"time"
)

type OrderLine struct {
	OrderID  string
	SKU      string
	Quantity int
}

type Batch struct {
	Reference         string
	SKU               string
	AvailableQuantity int
	ETA               time.Time
}

func (b *Batch) Allocate(line OrderLine) {
	b.AvailableQuantity -= line.Quantity
}
