package utility

import (
	"testing"
)

func TestBezierCurve2_CalculateBezierPointAtTime(t *testing.T) {
	bezierCurve := BezierCurve2{
		ControlPoints: []Vector2{
			{X: 0, Y: 0},
			{X: 1, Y: 1},
			{X: 2, Y: 0},
		},
	}

	tests := []struct {
		name     string
		t        float64
		expected Vector2
	}{
		{
			name:     "Start point",
			t:        0.0,
			expected: Vector2{X: 0.0, Y: 0.0},
		},
		{
			name:     "Mid point",
			t:        0.5,
			expected: Vector2{X: 1.0, Y: 0.5},
		},
		{
			name:     "End point",
			t:        1.0,
			expected: Vector2{X: 2.0, Y: 0.0},
		},
		{
			name:     "Custom point1",
			t:        0.2,
			expected: Vector2{X: 0.4000000000000001, Y: 0.32000000000000006},
		},
		{
			name:     "Custom point2",
			t:        1.2,
			expected: Vector2{X: 2.4, Y: -0.47999999999999987},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test
			got := bezierCurve.Interpolate(tt.t)

			// Compare
			if got != tt.expected {
				t.Errorf("Name: %s, got %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}
