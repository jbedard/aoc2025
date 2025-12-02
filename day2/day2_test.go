package main

import (
	"testing"
)

func TestCountInvalidIds(t *testing.T) {
	if c := addInvalidIds([]string{"10-12", "15-15", "20-22"}); c != 33 {
		t.Errorf("Expected 2 invalid IDs adding to 33, got %d", c)
	}
}

func TestIsInvalidId(t *testing.T) {
	tests := []struct {
		id       int
		expected bool
	}{
		{1234, false},
		{1212, true},
		{111, false},
		{1111, true},
		{1112, false},
		{11111, false},
		{1231234, false},
		{1010, true},
		{101101, true},
		{1188511885, true},
		{222222, true},
		{446446, true},
		{38593859, true},
	}

	for _, test := range tests {
		if result := isInvalidId(test.id); result != test.expected {
			t.Errorf("isInvalidId(%d) = %v; want %v", test.id, result, test.expected)
		}
	}
}

func TestPart1Example(t *testing.T) {
	if c := addInvalidIds([]string{"11-22"}); c != 33 {
		t.Errorf("Expected 33, got %d", c)
	}

	if c := addInvalidIds([]string{"95-115"}); c != 210 {
		t.Errorf("Expected 99, got %d", c)
	}

	if c := addInvalidIds([]string{"998-1012"}); c != 2009 {
		t.Errorf("Expected 1010, got %d", c)
	}

	if c := addInvalidIds([]string{"1188511880-1188511890"}); c != 1188511885 {
		t.Errorf("Expected 1188511885, got %d", c)
	}

	if c := addInvalidIds([]string{"222220-222224"}); c != 222222 {
		t.Errorf("Expected 222222, got %d", c)
	}

	if c := addInvalidIds([]string{"1698522-1698528"}); c != 0 {
		t.Errorf("Expected 0 invalid IDs, got %d", c)
	}

	if c := addInvalidIds([]string{"446443-446449"}); c != 446446 {
		t.Errorf("Expected 446446, got %d", c)
	}

	if c := addInvalidIds([]string{"38593856-38593862"}); c != 38593859 {
		t.Errorf("Expected 38593859, got %d", c)
	}

	if c := addInvalidIds([]string{"565653-565659", "824824821-824824827", "2121212118-2121212124"}); c != 2946602601 {
		t.Errorf("Expected 2946602601, got %d", c)
	}
}
