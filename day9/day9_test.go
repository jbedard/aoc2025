package main

import (
	"strings"
	"testing"

	"github.com/jbedard/aoc2025/lib"
)

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
