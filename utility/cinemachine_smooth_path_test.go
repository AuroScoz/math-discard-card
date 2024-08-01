package utility

import (
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkCinemachineSmoothPath_EvaluateLocalPosition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = broothMother.EvaluateLocalPosition(5)
		_ = point1.EvaluateLocalPosition(5)
		_ = linear1.EvaluateLocalPosition(5)
		_ = linear2.EvaluateLocalPosition(5)
		_ = square1.EvaluateLocalPosition(5)
		_ = curve1.EvaluateLocalPosition(5)
		_ = curve2.EvaluateLocalPosition(5)
	}
}

func BenchmarkCinemachineSmoothPath_PathLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = broothMother.PathLength()
		_ = point1.PathLength()
		_ = linear1.PathLength()
		_ = linear2.PathLength()
		_ = square1.PathLength()
		_ = curve1.PathLength()
		_ = curve2.PathLength()
	}
}

func TestParseWayPoints(t *testing.T) {
	tests := []struct {
		input          string
		expected       WayPoints
		expectingError bool
	}{
		{
			// Valid case
			input: "1.0,2.0,3.0,4.0,5.0,6.0,1.0,2.0,3.0,4.0,5.0,6.0",
			expected: WayPoints{
				{Position: Vector3{X: 1.0, Y: 2.0, Z: 3.0}},
				{Position: Vector3{X: 4.0, Y: 5.0, Z: 6.0}},
				{Position: Vector3{X: 1.0, Y: 2.0, Z: 3.0}},
				{Position: Vector3{X: 4.0, Y: 5.0, Z: 6.0}},
			},
			expectingError: false,
		},
		{
			// Valid case
			input: "1.0,2.0,3.0,4.0,5.0,6.0",
			expected: WayPoints{
				{Position: Vector3{X: 1.0, Y: 2.0, Z: 3.0}},
				{Position: Vector3{X: 4.0, Y: 5.0, Z: 6.0}},
			},
			expectingError: false,
		},
		{
			// Invalid case
			input:          "1.0,2.0",
			expected:       nil,
			expectingError: true,
		},
		{
			// Invalid case
			input:          "a,2.0,3.0",
			expected:       nil,
			expectingError: true,
		},
		{
			// Valid case
			input: "1.0,2.0,3.0",
			expected: WayPoints{
				{Position: Vector3{X: 1.0, Y: 2.0, Z: 3.0}},
			},
			expectingError: false,
		},
	}

	for _, test := range tests {
		// Test
		result, err := ParseWayPoints(test.input)

		// Compare
		if (err != nil) != test.expectingError {
			t.Errorf("unexpected error state for input %q: got error %v", test.input, err)
		}
		if !test.expectingError && !reflect.DeepEqual(result, test.expected) {
			t.Errorf("for input %q: expected %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestCinemachineSmoothPath_ToNativePathUnits(t *testing.T) {
	tests := []struct {
		name     string
		pos      float64
		route    *CinemachineSmoothPath
		units    PositionUnits
		expected float64
	}{
		{
			name:     "Test PathUnits",
			pos:      0.2,
			route:    linear2,
			units:    PATH_UNITS,
			expected: 0.2,
		},
		{
			name:     "Test Distance",
			pos:      0.2,
			route:    linear2,
			units:    DISTANCE,
			expected: 0.20000192001627207,
		},
		{
			name:     "Test Normalize",
			pos:      0.2,
			route:    linear2,
			units:    NORMALIZE,
			expected: 0.20000192001627207,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test
			result := tt.route.ToNativePathUnits(tt.pos, tt.units)

			// Compare
			if result != tt.expected {
				t.Errorf("Got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCinemachineSmoothPath_PathLength(t *testing.T) {
	tests := []struct {
		name     string
		route    *CinemachineSmoothPath
		expected float64
	}{
		{
			name:     "BroothMother",
			route:    broothMother,
			expected: 47.456248976636694,
		},
		{
			name:     "Point",
			route:    point1,
			expected: 0,
		},
		{
			name:     "3D Linear",
			route:    linear1,
			expected: 3.464101615137753,
		},
		{
			name:     "3D Square",
			route:    square1,
			expected: 8.772396039422429,
		},
		{
			name:     "Unlooped Curve",
			route:    curve1,
			expected: 7.098198349819181,
		},
		{
			name:     "Looped Curve",
			route:    curve2,
			expected: 14.660377381402991,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test
			len := tt.route.PathLength()

			// Compare
			if len != tt.expected {
				t.Errorf("Got %v, want %v", len, tt.expected)
			}
		})
	}
}

func TestCinemachineSmoothPath_MinMax(t *testing.T) {
	type cmp struct {
		pos       [2]float64
		pathUnits [2]float64
		distance  [2]float64
		normalize [2]float64
	}

	tests := []struct {
		name     string
		route    *CinemachineSmoothPath
		expected cmp
	}{
		{
			name:  "Point",
			route: point1,
			expected: cmp{
				pos:       [2]float64{0, 0},
				pathUnits: [2]float64{0, 0},
				distance:  [2]float64{0, 0},
				normalize: [2]float64{0, 1},
			},
		},
		{
			name:  "3D Linear",
			route: linear1,
			expected: cmp{
				pos:       [2]float64{0, 2},
				pathUnits: [2]float64{0, 2},
				distance:  [2]float64{0, 3.464101615137753},
				normalize: [2]float64{0, 1},
			},
		},
		{
			name:  "3D Square",
			route: square1,
			expected: cmp{
				pos:       [2]float64{0, 4},
				pathUnits: [2]float64{0, 4},
				distance:  [2]float64{0, 8.772396039422429},
				normalize: [2]float64{0, 1},
			},
		},
		{
			name:  "Unlooped Curve",
			route: curve1,
			expected: cmp{
				pos:       [2]float64{0, 4},
				pathUnits: [2]float64{0, 4},
				distance:  [2]float64{0, 7.098198349819181},
				normalize: [2]float64{0, 1},
			},
		},
		{
			name:  "Looped Curve",
			route: curve2,
			expected: cmp{
				pos:       [2]float64{0, 8},
				pathUnits: [2]float64{0, 8},
				distance:  [2]float64{0, 14.660377381402991},
				normalize: [2]float64{0, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test
			actual := cmp{
				pos:       [2]float64{tt.route.MinPos(), tt.route.MaxPos()},
				pathUnits: [2]float64{tt.route.MinUnit(PATH_UNITS), tt.route.MaxUnit(PATH_UNITS)},
				distance:  [2]float64{tt.route.MinUnit(DISTANCE), tt.route.MaxUnit(DISTANCE)},
				normalize: [2]float64{tt.route.MinUnit(NORMALIZE), tt.route.MaxUnit(NORMALIZE)},
			}

			// Compare
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Got %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestCinemachineSmoothPath_EvaluateLocalPosition(t *testing.T) {
	tests := []struct {
		name     string
		route    *CinemachineSmoothPath
		pos      float64
		expected Vector3
	}{
		// Point
		{
			name:     "Point-1",
			route:    point1,
			pos:      0.5,
			expected: Vector3{0, 0, 0},
		},
		// 3D Linear
		{
			name:     "3D Linear-1",
			route:    linear1,
			pos:      0.5,
			expected: Vector3{0.5, 0.5, 0.5},
		},
		{
			name:     "3D Linear-2",
			route:    linear1,
			pos:      1,
			expected: Vector3{1, 1, 1},
		},
		{
			name:     "3D Linear-3",
			route:    linear1,
			pos:      1.5,
			expected: Vector3{1.5, 1.5, 1.5},
		},
		{
			name:     "3D Linear-4",
			route:    linear1,
			pos:      2,
			expected: Vector3{2, 2, 2},
		},
		// 3D Square
		{
			name:     "3D Square-1",
			route:    square1,
			pos:      0,
			expected: Vector3{1, 0, 1},
		},
		{
			name:     "3D Square-2",
			route:    square1,
			pos:      1,
			expected: Vector3{1, 0, -1},
		},
		{
			name:     "3D Square-3",
			route:    square1,
			pos:      2,
			expected: Vector3{-1, 0, -1},
		},
		{
			name:     "3D Square-4",
			route:    square1,
			pos:      3,
			expected: Vector3{-1, 0, 1},
		},
		{
			name:     "3D Square-5",
			route:    square1,
			pos:      3.5,
			expected: Vector3{0.009803921568627527, 0, 1.3849056603773584},
		},
		{
			name:     "3D Square-6",
			route:    square1,
			pos:      4,
			expected: Vector3{1, 0, 1},
		},
		{
			name:     "3D Square-7",
			route:    square1,
			pos:      4.5,
			expected: Vector3{1.3774509803921569, 0, -0.002830188679245338},
		},
		// Unlooped Curve
		{
			name:     "Unlooped Curve - 1",
			route:    curve1,
			pos:      0.25,
			expected: Vector3{0.24999999999999994, 0.24999999999999994, 0.3671875},
		},
		{
			name:     "Unlooped Curve - 2",
			route:    curve1,
			pos:      0.5,
			expected: Vector3{0.4999999999999999, 0.4999999999999999, 0.6875},
		},
		{
			name:     "Unlooped Curve - 3",
			route:    curve1,
			pos:      1,
			expected: Vector3{1, 1, 1},
		},
		{
			name:     "Unlooped Curve - 4",
			route:    curve1,
			pos:      4,
			expected: Vector3{4, 4, 0},
		},
		// Looped Curve
		{
			name:     "Looped Curve - 1",
			route:    curve2,
			pos:      0.25,
			expected: Vector3{0.09520064681272894, 0.09520064681272894, 0.3671876397017596},
		},
		{
			name:     "Looped Curve - 2",
			route:    curve2,
			pos:      1,
			expected: Vector3{1, 1, 1},
		},
		{
			name:     "Looped Curve - 3",
			route:    curve2,
			pos:      2,
			expected: Vector3{2, 2, 0},
		},
		{
			name:     "Looped Curve - 4",
			route:    curve2,
			pos:      3,
			expected: Vector3{3, 3, -1},
		},
		{
			name:     "Looped Curve - 5",
			route:    curve2,
			pos:      4,
			expected: Vector3{4, 4, 0},
		},
		{
			name:     "Looped Curve - 6",
			route:    curve2,
			pos:      5,
			expected: Vector3{3, 3, 1},
		},
		{
			name:     "Looped Curve - 7",
			route:    curve2,
			pos:      6,
			expected: Vector3{2, 2, 0},
		},
		{
			name:     "Looped Curve - 8",
			route:    curve2,
			pos:      7,
			expected: Vector3{1, 1, -1},
		},
		{
			name:     "Looped Curve - 9",
			route:    curve2,
			pos:      7.75,
			expected: Vector3{0.09525300266789713, 0.09525300266789713, -0.3697396404192234},
		},
		{
			name:     "Looped Curve - 10",
			route:    curve2,
			pos:      8,
			expected: Vector3{0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test
			result := tt.route.EvaluateLocalPosition(tt.pos)

			// Compare
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Got %v, want %v", result, tt.expected)
			}
		})
	}
}

var (
	// BroothMother
	broothMother = NewCinemachineSmoothPath(
		false,
		func(s string) WayPoints {
			ps, _ := ParseWayPoints(s)
			fmt.Printf("%v", ps)
			return ps
		}("0,0,0,-3.440392,0,-2.898927,1.447635,0,-7.49635,8.594374,0,-1.949572,3.946407,0,4.45786,-4.752399,0,6.805671,-11.723,0,2.915641"),
	)

	// Point
	point1 = NewCinemachineSmoothPath(
		false, // Un-Looped
		WayPoints{
			{Position: Vector3{X: 0, Y: 0, Z: 0}},
		},
	)

	// 3D Linear
	linear1 = NewCinemachineSmoothPath(
		false, // Un-Looped
		WayPoints{
			{Position: Vector3{X: 0, Y: 0, Z: 0}},
			{Position: Vector3{X: 1, Y: 1, Z: 1}},
			{Position: Vector3{X: 2, Y: 2, Z: 2}},
		},
	)

	linear2 = NewCinemachineSmoothPath(
		false, // Un-Looped
		WayPoints{
			{Position: Vector3{X: 0, Y: 0, Z: 0}},
			{Position: Vector3{X: 1, Y: 0, Z: 0}},
		},
	)

	// 3D Square
	square1 = NewCinemachineSmoothPath(
		true, // Looped
		WayPoints{
			{Position: Vector3{X: 1, Y: 0, Z: 1}},
			{Position: Vector3{X: 1, Y: 0, Z: -1}},
			{Position: Vector3{X: -1, Y: 0, Z: -1}},
			{Position: Vector3{X: -1, Y: 0, Z: 1}},
		},
	)

	// Unlooped Curve (S)
	curve1 = NewCinemachineSmoothPath(
		false,
		WayPoints{
			{Position: Vector3{X: 0, Y: 0, Z: 0}},
			{Position: Vector3{X: 1, Y: 1, Z: 1}},
			{Position: Vector3{X: 2, Y: 2, Z: 0}},
			{Position: Vector3{X: 3, Y: 3, Z: -1}},
			{Position: Vector3{X: 4, Y: 4, Z: 0}},
		},
	)

	// Looped Curve (8)
	curve2 = NewCinemachineSmoothPath(
		true,
		WayPoints{
			{Position: Vector3{X: 0, Y: 0, Z: 0}},
			{Position: Vector3{X: 1, Y: 1, Z: 1}},
			{Position: Vector3{X: 2, Y: 2, Z: 0}},
			{Position: Vector3{X: 3, Y: 3, Z: -1}},
			{Position: Vector3{X: 4, Y: 4, Z: 0}},
			{Position: Vector3{X: 3, Y: 3, Z: 1}},
			{Position: Vector3{X: 2, Y: 2, Z: 0}},
			{Position: Vector3{X: 1, Y: 1, Z: -1}},
		},
	)
)
