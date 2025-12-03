package main

import (
	"testing"
)

func TestMaxJoltage(t *testing.T) {
	tests := []struct {
		input    string
		expected uint64
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}

	for _, test := range tests {
		if got := maxJoltage(test.input); got != test.expected {
			t.Errorf("maxJoltage(%q) = %d; want %d", test.input, got, test.expected)
		}
	}
}
