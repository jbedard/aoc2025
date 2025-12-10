package main

import (
	"strings"
	"testing"
)

const EXAMPLE_INPUT = `
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`

func TestSchematicsUtil(t *testing.T) {
	s := readSchematics(strings.TrimSpace(EXAMPLE_INPUT))

	if len(s) != 3 {
		t.Fatalf("expected 3 schematics, got %d", len(s))
	}

	if s[0].String() != "on: 0110000000000000 switches: [0001000000000000 0101000000000000 0010000000000000 0011000000000000 1010000000000000 1100000000000000] levels: [3 5 4 7]" {
		t.Errorf("schematic string mismatch: %s", s[0].String())
	}

	if s[1].String() != "on: 0001000000000000 switches: [1011100000000000 0011000000000000 1000100000000000 1110000000000000 0111100000000000] levels: [7 5 12 7 2]" {
		t.Errorf("schematic string mismatch: %s", s[1].String())
	}

	if s[2].String() != "on: 0111010000000000 switches: [1111100000000000 1001100000000000 1110110000000000 0110000000000000] levels: [10 11 11 5 10 5]" {
		t.Errorf("schematic string mismatch: %s", s[2].String())
	}
}

func TestExample(t *testing.T) {
	schematics := readSchematics(strings.TrimSpace(EXAMPLE_INPUT))

	part1Result := part1(schematics)
	if part1Result != 7 {
		t.Errorf("Part 1: expected 7, got %d", part1Result)
	}

	part2Result := part2(schematics)
	if part2Result != 33 {
		t.Errorf("Part 2: expected 33, got %d", part2Result)
	}
}
