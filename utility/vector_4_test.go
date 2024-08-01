package utility

import (
	"testing"
)

func TestVector4_ToArray(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector4
		expected [4]float64
	}{
		{
			name:     "Test1",
			vector:   Vector4{X: 3, Y: 4, Z: 5, W: 6},
			expected: [4]float64{3, 4, 5, 6},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			arr := tc.vector.ToArray()

			// Compare
			if arr != tc.expected {
				t.Errorf("Got %v, want %v", arr, tc.expected)
			}
		})
	}
}

func TestVector4_ToSlice(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector4
		expected []float64
	}{
		{
			name:     "Test1",
			vector:   Vector4{X: 3, Y: 4, Z: 5, W: 6},
			expected: []float64{3, 4, 5, 6},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			slice := tc.vector.ToSlice()

			// Compare
			if len(slice) != len(tc.expected) {
				t.Errorf("Expected slice length %d, but got %d", len(tc.expected), len(slice))
			}

			for i := range slice {
				if slice[i] != tc.expected[i] {
					t.Errorf("Got %v, want %v", slice, tc.expected)
				}
			}
		})
	}
}

func TestVector4_Add(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector4
		other    Vector4
		expected Vector4
	}{
		{
			// Vector4{ 3+1, 4+2, 5+3, 6+4 } = Vector4{ 4, 6, 8, 10 }
			name:     "Test1",
			vector:   Vector4{X: 3, Y: 4, Z: 5, W: 6},
			other:    Vector4{X: 1, Y: 2, Z: 3, W: 4},
			expected: Vector4{X: 4, Y: 6, Z: 8, W: 10},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			result := tc.vector.Add(tc.other)

			// Compare
			if result != tc.expected {
				t.Errorf("Got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestVector4_Sub(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector4
		other    Vector4
		expected Vector4
	}{
		{
			// Vector4{ 3-1, 4-2, 5-3, 6-4 } = Vector4{ 2, 2, 2, 2 }
			name:     "Test1",
			vector:   Vector4{X: 3, Y: 4, Z: 5, W: 6},
			other:    Vector4{X: 1, Y: 2, Z: 3, W: 4},
			expected: Vector4{X: 2, Y: 2, Z: 2, W: 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			result := tc.vector.Sub(tc.other)

			// Compare
			if result != tc.expected {
				t.Errorf("Got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestVector4_Mul(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector4
		scalar   float64
		expected Vector4
	}{
		{
			// Vector4{ 3*2, 4*2, 5*2, 6*2 } = Vector4{ 6, 8, 10, 12 }
			name:     "Test1",
			vector:   Vector4{X: 3, Y: 4, Z: 5, W: 6},
			scalar:   2,
			expected: Vector4{X: 6, Y: 8, Z: 10, W: 12},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			result := tc.vector.Mul(tc.scalar)

			// Compare
			if result != tc.expected {
				t.Errorf("Got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestVector4_Lerp(t *testing.T) {
	tests := []struct {
		name     string
		v1       Vector4
		v2       Vector4
		t        float64
		expected Vector4
	}{
		{
			name:     "Test1",
			v1:       Vector4{X: 3, Y: 4, Z: 5, W: 6},
			v2:       Vector4{X: 6, Y: 8, Z: 10, W: 12},
			t:        0.5,
			expected: Vector4{X: 4.5, Y: 6, Z: 7.5, W: 9},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			result := tc.v1.Lerp(tc.v2, tc.t)

			// Compare
			if result != tc.expected {
				t.Errorf("Got %v, want %v", result, tc.expected)
			}
		})
	}
}
