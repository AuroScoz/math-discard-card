package utility

import (
	"testing"
)

func TestRandomFloatBetweenInts(t *testing.T) {
	tests := []struct {
		min      int
		max      int
		expected bool // Expecting no error
	}{
		{0, 10, true},
		{10, 10, true},
		{10, 0, false},
	}

	for _, tt := range tests {
		// Test
		result, err := RandomFloatBetweenInts(tt.min, tt.max)

		// Compare
		if (err == nil) != tt.expected {
			t.Errorf("Got %v , expected %v", (err == nil), tt.expected)
		}

		if err == nil && (result < float64(tt.min) || result > float64(tt.max)) {
			t.Errorf("Result: %v, out of range", result)
		}
	}
}

func TestGetRandomIntFromMinMax(t *testing.T) {
	tests := []struct {
		min      int
		max      int
		expected bool // Expecting no error
	}{
		{0, 100, true},
		{100, 100, true},
		{100, 0, false},
	}

	for _, tt := range tests {
		// Test
		result, err := GetRandomIntFromMinMax(tt.min, tt.max)

		// Compare
		if (err == nil) != tt.expected {
			t.Errorf("Got %v , expected %v", (err == nil), tt.expected)
		}

		if err == nil && (result < tt.min || result > tt.max) {
			t.Errorf("Result: %v, out of range", result)
		}
	}
}

func TestGetRandomTFromSlice(t *testing.T) {
	tests := []struct {
		slice    []int
		expected bool // Expecting no error
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{}, false},
	}

	for _, tt := range tests {
		// Test
		_, err := GetRandomTFromSlice(tt.slice)

		// Compare
		if (err == nil) != tt.expected {
			t.Errorf("Got %v , expected %v", (err == nil), tt.expected)
		}
	}
}

func TestGetRndIntFromRangeStr(t *testing.T) {
	tests := []struct {
		input     string
		delimiter string
		expected  bool // Expecting no error
	}{
		{"10~20", "~", true},
		{"20~10", "~", false},
		{"10~", "~", false},
		{"~20", "~", false},
	}

	for _, tt := range tests {
		// Test
		_, err := GetRndIntFromRangeStr(tt.input, tt.delimiter)

		// Compare
		if (err == nil) != tt.expected {
			t.Errorf("Got %v , expected %v", (err == nil), tt.expected)
		}
	}
}

func TestGetRndIntFromString(t *testing.T) {
	tests := []struct {
		input     string
		delimiter string
		expected  bool // Expecting no error
	}{
		{"100,200,300", ",", true},   // No Error
		{"100,200,300,", ",", true},  // Error
		{"100,abc,300", ",", true},   // Error
		{",,,,,,,,,,,,", ",", false}, // Error
		{"", ",", false},             // Error
	}

	for _, tt := range tests {
		// Test
		_, err := GetRndIntFromString(tt.input, tt.delimiter)

		// Compare
		if (err == nil) != tt.expected {
			t.Errorf("Got %v , expected %v", (err == nil), tt.expected)
		}
	}
}

func TestGetRndStrFromString(t *testing.T) {
	tests := []struct {
		input     string
		delimiter string
		expected  bool // Expecting no error
	}{
		{"100,200,300", ",", true},   // No Error
		{"100,200,300,", ",", true},  // No Error
		{",,,,,,,,,,,,", ",", false}, // Error
		{"", ",", false},             // Error
	}

	for _, tt := range tests {
		// Test
		_, err := GetRndStrFromString(tt.input, tt.delimiter)

		// Compare
		if (err == nil) != tt.expected {
			t.Errorf("Got %v , expected %v", (err == nil), tt.expected)
		}
	}
}
