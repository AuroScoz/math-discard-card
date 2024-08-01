package utility

import (
	"math"
	"testing"
)

func TestVector2_ToArray(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector2
		expected [2]float64
	}{
		{
			name:     "Test1",
			vector:   Vector2{X: 3, Y: 4},
			expected: [2]float64{3, 4},
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

func TestVector2_ToSlice(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector2
		expected []float64
	}{
		{
			name:     "Test1",
			vector:   Vector2{X: 3, Y: 4},
			expected: []float64{3, 4},
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

func TestVector2_Add(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector2
		other    Vector2
		expected Vector2
	}{
		{
			// Vector2{ 3+1, 4+2 } = Vector2{ 4, 6 }
			name:     "Test1",
			vector:   Vector2{X: 3, Y: 4},
			other:    Vector2{X: 1, Y: 2},
			expected: Vector2{X: 4, Y: 6},
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

func TestVector2_Sub(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector2
		other    Vector2
		expected Vector2
	}{
		{
			// Vector2{ 3-1, 4-2 } = Vector2{ 2, 2 }
			name:     "Test1",
			vector:   Vector2{X: 3, Y: 4},
			other:    Vector2{X: 1, Y: 2},
			expected: Vector2{X: 2, Y: 2},
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

func TestVector2_Mul(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector2
		scalar   float64
		expected Vector2
	}{
		{
			// Vector2{ 3*2, 4*2 } = Vector2{ 6, 8 }
			name:     "Test1",
			vector:   Vector2{X: 3, Y: 4},
			scalar:   2,
			expected: Vector2{X: 6, Y: 8},
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

func TestVector2_Cross(t *testing.T) {
	tests := []struct {
		name     string
		v1       Vector2
		v2       Vector2
		expected float64
	}{
		{
			name:     "OrthogonalVectors",
			v1:       Vector2{X: 1, Y: 0},
			v2:       Vector2{X: 0, Y: 1},
			expected: 1,
		},
		{
			name:     "OppositeVectors",
			v1:       Vector2{X: 1, Y: 0},
			v2:       Vector2{X: -1, Y: 0},
			expected: 0,
		},
		{
			name:     "ParallelVectors",
			v1:       Vector2{X: 1, Y: 0},
			v2:       Vector2{X: 2, Y: 0},
			expected: 0,
		},
		{
			name:     "ArbitraryVectors",
			v1:       Vector2{X: 3, Y: 4},
			v2:       Vector2{X: 5, Y: 6},
			expected: -2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			result := tc.v1.Cross(tc.v2)

			// Compare
			if result != tc.expected {
				t.Errorf("Got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestVector2_Dot(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector2
		other    Vector2
		expected float64
	}{
		{
			// 3*1 + 4*2 = 8
			name:     "Test1",
			vector:   Vector2{X: 3, Y: 4},
			other:    Vector2{X: 1, Y: 2},
			expected: 11,
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

func TestVector2_Magnitude(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector2
		expected float64
	}{
		{
			name:     "Test1",
			vector:   Vector2{X: 3, Y: 4},
			expected: 5,
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

func TestVector2_Normalize(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector2
		expected Vector2
	}{
		{
			name:     "Test1",
			vector:   Vector2{X: 3, Y: 4},
			expected: Vector2{X: 0.6, Y: 0.8},
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

func TestVector2_Distance(t *testing.T) {
	tests := []struct {
		name     string
		from     Vector2
		to       Vector2
		expected float64
	}{
		{
			// Sqrt((4-3)^2 + (4-4)^2) = 1
			name:     "Test1",
			from:     Vector2{X: 3, Y: 4},
			to:       Vector2{X: 4, Y: 4},
			expected: 1,
		},
		{
			// Sqrt((5-3)^2 + (7-4)^2) = Sqrt(13)
			name:     "Test2",
			from:     Vector2{X: 3, Y: 4},
			to:       Vector2{X: 5, Y: 7},
			expected: math.Sqrt(13),
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

func TestGetDistanceFromPointToLine(t *testing.T) {
	tests := []struct {
		name      string
		point     Vector2
		linePoint Vector2
		lineDir   Vector2
		expected  float64
	}{
		{
			name:      "Test1",
			point:     Vector2{X: 3, Y: 4},
			linePoint: Vector2{X: 1, Y: 1},
			lineDir:   Vector2{X: 1, Y: 1},
			expected:  math.Sqrt(2) / 2,
		},
		{
			name:      "Test2",
			point:     Vector2{X: 5, Y: 5},
			linePoint: Vector2{X: 1, Y: 1},
			lineDir:   Vector2{X: 1, Y: 1},
			expected:  0,
		},
		{
			name:      "Test3",
			point:     Vector2{X: 1, Y: 3},
			linePoint: Vector2{X: 1, Y: 1},
			lineDir:   Vector2{X: 0, Y: 1},
			expected:  0,
		},
		{
			name:      "Test4",
			point:     Vector2{X: 1, Y: 3},
			linePoint: Vector2{X: 1, Y: 1},
			lineDir:   Vector2{X: 1, Y: 0},
			expected:  2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			result := GetDistanceFromPointToLine(tc.point, tc.linePoint, tc.lineDir)

			// Compare
			if !approxEqual(result, tc.expected) {
				t.Errorf("Got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestAngleToVector(t *testing.T) {
	tests := []struct {
		name        string
		angleDeg    float64
		expected    Vector2
		approximate bool // 近似比較浮點數精度
	}{
		{
			name:        "Test1",
			angleDeg:    45,
			expected:    Vector2{X: math.Sqrt(2) / 2, Y: math.Sqrt(2) / 2},
			approximate: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test
			result := AngleToVector(tc.angleDeg)

			// Compare
			if !tc.approximate {
				if result != tc.expected {
					t.Errorf("Got %v, want %v", result, tc.expected)
				}
			} else {
				if !approxEqual(result.X, tc.expected.X) || !approxEqual(result.Y, tc.expected.Y) {
					t.Errorf("Got %v, want %v", result, tc.expected)
				}
			}
		})
	}
}

func approxEqual(a, b float64) bool {
	const epsilon = 1e-6
	return math.Abs(a-b) < epsilon
}
