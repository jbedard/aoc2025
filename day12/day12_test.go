package main

import (
	"strings"
	"testing"
)

const EXAMPLE_INPUT = `
0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2
`

func TestParseExample(t *testing.T) {
	d := parseInput(EXAMPLE_INPUT)

	if len(d.presents) != 6 {
		t.Errorf("Expected 6 presents, got %d", len(d.presents))
	}

	for _, p := range d.presents {
		if p.size() != 7 {
			t.Errorf("Expected present size 7 in example, got %d", p.size())
		}
	}

	if len(d.regions) != 3 {
		t.Errorf("Expected 3 regions, got %d", len(d.regions))
	}

	if d.regions[0].width != 4 || d.regions[0].height != 4 {
		t.Errorf("Expected region 0 to be 4x4, got %dx%d", d.regions[0].width, d.regions[0].height)
	}

	if d.regions[1].width != 12 || d.regions[1].height != 5 {
		t.Errorf("Expected region 1 to be 12x5, got %dx%d", d.regions[1].width, d.regions[1].height)
	}

	if d.regions[2].width != 12 || d.regions[2].height != 5 {
		t.Errorf("Expected region 2 to be 12x5, got %dx%d", d.regions[2].width, d.regions[2].height)
	}

	for i, r := range d.regions {
		if len(r.presentCounts) != len(d.presents) {
			t.Errorf("Region %d: expected presents counts (%d) to align with presents total (%d)", i, len(r.presentCounts), len(d.presents))
		}
	}
}

func TestExample(t *testing.T) {
	d := parseInput(strings.TrimSpace(EXAMPLE_INPUT))

	r1 := part1(d)
	if r1 != 2 {
		t.Errorf("Part 1: expected 2, got %d", r1)
	}
}

func TestPresentParser(t *testing.T) {
	p0 := parsePresent(`#.#
###
..#`)
	if p0.size() != 6 {
		t.Errorf("Expected present size 6")
	}

	p1 := parsePresent(`..#
#.#
..#`)
	if p1.size() != 4 {
		t.Errorf("Expected present size 4")
	}
}
