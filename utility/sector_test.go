package utility

import (
	"testing"
)

func TestIsPointInSector(t *testing.T) {
	sector := Sector{
		Center:     Vector2{X: 0, Y: 0},
		Radius:     5,
		Angle:      0,
		AngleRange: 90,
	}

	tests := []struct {
		point  Vector2
		within bool
	}{
		{
			point:  Vector2{X: 3, Y: 3},
			within: true,
		},
		{
			point:  Vector2{X: 5, Y: 0},
			within: true,
		},
		{
			point:  Vector2{X: 0, Y: 5},
			within: false,
		},
		{
			point:  Vector2{X: 6, Y: 6},
			within: false,
		},
		{
			point:  Vector2{X: 0, Y: -6},
			within: false,
		},
	}

	for _, tt := range tests {
		// Test
		within := sector.IsPointInSector(tt.point)

		// Compare
		if tt.within != sector.IsPointInSector(tt.point) {
			t.Errorf("Got %v, expected %v", within, tt.within)
		}
	}
}
