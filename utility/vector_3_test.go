package utility

import (
	"math"
	"testing"
)

func TestVector3_ToArray(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector3
		expected [3]float64
	}{
		{
			name:     "Test1",
			vector:   Vector3{X: 3, Y: 4, Z: 5},
			expected: [3]float64{3, 4, 5},
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

func TestVector3_ToSlice(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector3
		expected []float64
	}{
		{
			name:     "Test1",
			vector:   Vector3{X: 3, Y: 4, Z: 5},
			expected: []float64{3, 4, 5},
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

func TestVector3_Add(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector3
		other    Vector3
		expected Vector3
	}{
		{
			// Vector3{ 3+1, 4+2, 5+3 } = Vector3{ 4, 6, 8 }
			name:     "Test1",
			vector:   Vector3{X: 3, Y: 4, Z: 5},
			other:    Vector3{X: 1, Y: 2, Z: 3},
			expected: Vector3{X: 4, Y: 6, Z: 8},
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

func TestVector3_Sub(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector3
		other    Vector3
		expected Vector3
	}{
		{
			// Vector3{ 3-1, 4-2, 5-3 } = Vector3{ 2, 2, 2 }
			name:     "Test1",
			vector:   Vector3{X: 3, Y: 4, Z: 5},
			other:    Vector3{X: 1, Y: 2, Z: 3},
			expected: Vector3{X: 2, Y: 2, Z: 2},
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

func TestVector3_Mul(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector3
		scalar   float64
		expected Vector3
	}{
		{
			// Vector3{ 3*2, 4*2, 5*2 } = Vector3{ 6, 8, 10 }
			name:     "Test1",
			vector:   Vector3{X: 3, Y: 4, Z: 5},
			scalar:   2,
			expected: Vector3{X: 6, Y: 8, Z: 10},
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

func TestVector3_Cross(t *testing.T) {
	tests := []struct {
		name     string
		v1       Vector3
		v2       Vector3
		expected Vector3
	}{
		{
			name:     "OrthogonalVectors",
			v1:       Vector3{X: 1, Y: 0, Z: 0},
			v2:       Vector3{X: 0, Y: 1, Z: 0},
			expected: Vector3{X: 0, Y: 0, Z: 1},
		},
		{
			name:     "ArbitraryVectors",
			v1:       Vector3{X: 1, Y: 2, Z: 3},
			v2:       Vector3{X: 4, Y: 5, Z: 6},
			expected: Vector3{X: -3, Y: 6, Z: -3},
		},
		{
			name:     "ParallelVectors",
			v1:       Vector3{X: 1, Y: 0, Z: 0},
			v2:       Vector3{X: 2, Y: 0, Z: 0},
			expected: Vector3{X: 0, Y: 0, Z: 0},
		},
		{
			name:     "OppositeVectors",
			v1:       Vector3{X: 1, Y: 0, Z: 0},
			v2:       Vector3{X: -1, Y: 0, Z: 0},
			expected: Vector3{X: 0, Y: 0, Z: 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			result := tc.v1.Cross(tc.v2)

			// Compare
			if result.X != tc.expected.X || result.Y != tc.expected.Y || result.Z != tc.expected.Z {
				t.Errorf("Got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestVector3_Dot(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector3
		other    Vector3
		expected float64
	}{
		{
			// 3*1 + 4*2 + 5*3 = 26
			name:     "Test1",
			vector:   Vector3{X: 3, Y: 4, Z: 5},
			other:    Vector3{X: 1, Y: 2, Z: 3},
			expected: 26,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			result := tc.vector.Dot(tc.other)

			// Compare
			if result != tc.expected {
				t.Errorf("Got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestVector3_Magnitude(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector3
		expected float64
	}{
		{
			name:     "UnitVector",
			vector:   Vector3{X: 1, Y: 0, Z: 0},
			expected: 1,
		},
		{
			// Sqrt(3^2 + 4^2 + 5^2)
			name:     "NonUnitVector",
			vector:   Vector3{X: 3, Y: 4, Z: 5},
			expected: math.Sqrt(50),
		},
		{
			name:     "ZeroVector",
			vector:   Vector3{X: 0, Y: 0, Z: 0},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			result := tc.vector.Magnitude()

			// Compare
			if result != tc.expected {
				t.Errorf("Got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestVector3_Normalize(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector3
		expected Vector3
	}{
		{
			name:     "NonUnitVector",
			vector:   Vector3{X: 3, Y: 4, Z: 5},
			expected: Vector3{X: 3.0 / math.Sqrt(50), Y: 4.0 / math.Sqrt(50), Z: 5.0 / math.Sqrt(50)},
		},
		{
			name:     "ZeroVector",
			vector:   Vector3{X: 0, Y: 0, Z: 0},
			expected: Vector3{X: 0, Y: 0, Z: 0},
		},
		{
			name:     "NegativeValues",
			vector:   Vector3{X: -3, Y: -4, Z: -5},
			expected: Vector3{X: -3.0 / math.Sqrt(50), Y: -4.0 / math.Sqrt(50), Z: -5.0 / math.Sqrt(50)},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			result := tc.vector.Normalize()

			// Compare
			if result != tc.expected {
				t.Errorf("Got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestVector3_Distance(t *testing.T) {
	tests := []struct {
		name     string
		from     Vector3
		to       Vector3
		expected float64
	}{
		{
			// Sqrt((4-3)^2 + (4-4)^2 + (5-5)^2) = 1
			name:     "Test1",
			from:     Vector3{X: 3, Y: 4, Z: 5},
			to:       Vector3{X: 4, Y: 4, Z: 5},
			expected: 1,
		},
		{
			// Sqrt((5-3)^2 + (7-4)^2 + (13-5)^2) = Sqrt(77)
			name:     "Test2",
			from:     Vector3{X: 3, Y: 4, Z: 5},
			to:       Vector3{X: 5, Y: 7, Z: 13},
			expected: math.Sqrt(77),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			result := tc.from.Distance(tc.to)

			// Compare
			if result != tc.expected {
				t.Errorf("Got %v, want %v", result, tc.expected)
			}
		})
	}
}
