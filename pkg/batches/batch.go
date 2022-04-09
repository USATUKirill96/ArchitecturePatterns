package batches

import (
	"time"
)

type Batch struct {
	ID        int
	Reference string
	SKU       string
	ETA       time.Time

	purchasedQuantity int
	allocations       Allocations
}

func NewBatch(reference, sku string, eta time.Time, purchasedQuantity int) *Batch {
	return &Batch{Reference: reference, SKU: sku, ETA: eta, purchasedQuantity: purchasedQuantity}
}

func (b *Batch) Allocate(line OrderLine) {
	idx := b.allocations.getIdx(line)
	if idx == -1 {
		b.allocations = append(b.allocations, line)
	}
}

func (b *Batch) Deallocate(line OrderLine) {
	idx := b.allocations.getIdx(line)
	if idx != -1 {
		b.allocations.remove(idx)
	}

}

func (b *Batch) CanAllocate(line OrderLine) bool {
	return line.SKU == b.SKU && b.AvailableQuantity() >= line.Quantity
}

func (b *Batch) AllocatedQuantity() int {
	var sum int
	for _, allocation := range b.allocations {
		sum += allocation.Quantity
	}
	return sum
}

func (b *Batch) AvailableQuantity() int {
	return b.purchasedQuantity - b.AllocatedQuantity()
}

type Allocations []OrderLine

func (a Allocations) getIdx(line OrderLine) int {
	for i, allocation := range a {
		if allocation == line {
			return i
		}
	}
	return -1
}

func (a Allocations) remove(idx int) {
	a = append(a[:idx], a[idx+1:]...)
}
