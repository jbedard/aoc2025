package main

import (
	"strings"
	"testing"

	"github.com/jbedard/aoc2025/lib"
)

// Practice problem test case
const TEST_INPUT = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`

func TestPart1(t *testing.T) {
	points := []lib.Pos2d{}
	for line := range lib.ReadLines(strings.TrimSpace(TEST_INPUT)) {
		points = append(points, lib.ReadPos2d(line))
	}

	r := part1(points)
	if r != 50 {
		t.Errorf("expected 50, got %d", r)
	}

	r2 := part2_b(points)
	if r2 != 24 {
		t.Errorf("expected 24, got %d", r2)
	}
}

// Custom test cases
const TEST_INPUT_2 = `7,1
11,1
11,3
13,3
13,5
11,5
11,7
9,7
9,5
2,5
2,3
7,3
`

const TEST_INPUT_3 = `
1,2
8,2
8,4
6,4
6,5
9,5
9,8
1,8
`

var TESTS = map[string]int{
	TEST_INPUT_2: 36,
	TEST_INPUT_3: 56,
}

func Test2(t *testing.T) {
	for input, expected := range TESTS {
		points := []lib.Pos2d{}
		for line := range lib.ReadLines(strings.TrimSpace(input)) {
			points = append(points, lib.ReadPos2d(line))
		}

		r2 := part2_b(points)
		if r2 != expected {
			t.Errorf("expected %d, got %d", expected, r2)
		}
	}
}
