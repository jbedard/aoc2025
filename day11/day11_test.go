package main

import (
	"strings"
	"testing"
)

const EXAMPLE_INPUT = `
aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`

func TestExample(t *testing.T) {
	d := parseInput(strings.TrimSpace(EXAMPLE_INPUT))

	r1 := part1(d)
	if r1 != 5 {
		t.Errorf("Part 1: expected 5, got %d", r1)
	}
}

const EXAMPLE_INPUT2 = `
svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out
`

func TestExample2(t *testing.T) {
	d := parseInput(strings.TrimSpace(EXAMPLE_INPUT2))
	r2 := part2(d)
	if r2 != 2 {
		t.Errorf("Part 2: expected 2, got %d", r2)
	}
}
