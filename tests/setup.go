package tests

import (
	"testing"
)

func (tc TestCase) Setup(tb testing.TB) func() {
	tc.CreateBatches()
	tc.CreateOrderLines()
	return func() {
		tc.Delete()
	}
}
