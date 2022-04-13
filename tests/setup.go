package tests

import (
	"testing"
)

func (tc TestCase) Setup(tb testing.TB) func(tb testing.TB) {
	tc.createBatches()
	tc.createOrderLines()
	return func(tb testing.TB) {
		tc.delete()
	}
}
