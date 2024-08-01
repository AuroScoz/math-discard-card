package utility

import (
	"math"
	"testing"
)

func TestGetNextIdx(t *testing.T) {
	acc := NewAccumulator()

	// Test multiple calls
	for i := 0; i <= 10; i++ {
		if idx := acc.NextIndex("testKey"); idx != i {
			t.Errorf("Expected %d, got %d", i, idx)
		}
	}

	// Test max int overflow
	MAX_INT_KEY := "maxIntKey"
	acc.keyValueMap[MAX_INT_KEY] = math.MaxInt
	for i := 0; i <= 10; i++ {
		if idx := acc.NextIndex(MAX_INT_KEY); idx != i {
			t.Errorf("Expected %d, got %d", i, idx)
		}
	}
}
