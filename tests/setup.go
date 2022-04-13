package tests

import (
	"testing"
)

func (tc TestCase) Setup(tb testing.TB) func() {
	tc.createBatches()
	tc.createOrderLines()
	return func() {
		tc.Delete()
	}
}
