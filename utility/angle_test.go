package utility

import (
	"math"
	"testing"
)

func TestAngleDiff(t *testing.T) {
	tests := []struct {
		a1, a2   float64
		expected float64
	}{
		{0, 0, 0},
		{0, 90, 90},
		{0, 180, 180},
		{0, 270, 90},
		{0, 360, 0},
		{90, 180, 90},
		{180, 270, 90},
		{90, 270, 180},
		{360, 360, 0},
		{720, 360, 0},
		{-90, 90, 180},
		{-180, 180, 0},
		{450, 90, 0},
	}

	for _, test := range tests {
		// Test
		result := AngleDiff(test.a1, test.a2)

		// Compare
		if math.Abs(result-test.expected) > 1e-9 {
			t.Errorf("AngleDiff(%f, %f) = %f; expected %f", test.a1, test.a2, result, test.expected)
		}
	}
}
