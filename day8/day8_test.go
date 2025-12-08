package main

import (
	"strings"
	"testing"

	"github.com/jbedard/aoc2025/lib"
)

func TestPart1(t *testing.T) {
	input := `
162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
`
	points := []lib.Pos{}
	for line := range lib.ReadLines(strings.TrimSpace(input)) {
		points = append(points, lib.ReadPoint(line))
	}

	connectCount := 10
	result1, result2 := part1(points, connectCount)

	if 40 != result1 {
		t.Errorf("expected 40, got %d", result1)
	}

	if 25272 != result2 {
		t.Errorf("expected 25272, got %d", result2)
	}
}
