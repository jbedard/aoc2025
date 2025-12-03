package main

import (
	"testing"
)

func TestMaxJoltage(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1991", 99},
		{"9119", 99},
		{"1234", 34},
		{"9876", 98},
		{"13579", 79},
	}

	for _, test := range tests {
		if got := maxJoltage(test.input); got != test.expected {
			t.Errorf("maxJoltage(%q) = %d; want %d", test.input, got, test.expected)
		}
	}
}
