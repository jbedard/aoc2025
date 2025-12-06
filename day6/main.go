package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jbedard/aoc2025/lib"
)

func parseInput1() ([][]int, []string) {
	g := [][]int{}
	o := []string{}

	for line := range lib.ReadInputLines() {
		line = strings.TrimLeft(line, " ")

		if strings.HasPrefix(line, "*") || strings.HasPrefix(line, "+") {
			for field := range strings.FieldsSeq(line) {
				o = append(o, field)
			}
		} else {
			r := []int{}
			for field := range strings.FieldsSeq(line) {
				n, _ := strconv.Atoi(field)
				r = append(r, int(n))
			}
			g = append(g, r)
		}
	}

	return g, o
}

func main() {
	g, o := parseInput1()
	println("Part 1:", part1(g, o))

	g, o = parseInput2()
	println("Part 2:", part1(g, o))
}

func part1(g [][]int, o []string) int {
	t := 0

	for i, op := range o {
		v := 0
		if op == "*" {
			v = 1
		}

		for _, row := range g {
			// ignore zero values which may be padding for numbers
			// of less digits
			if row[i] == 0 {
				continue
			}

			switch op {
			case "*":
				v = v * row[i]
			case "+":
				v = v + row[i]
			}
		}

		t += v
	}

	return t
}

func parseInput2() ([][]int, []string) {
	/*
		123 328  51 64
		 45 64  387 23
		  6 98  215 314
		*   +   *   +

		must transform to:

		g: {
			{4,   431, 623},
			{175, 581, 32},
			{8,   248, 369},
			{356, 24,  1}
		}
		o: "*", "+", "*", "+"
	*/

	lines := slices.Collect(lib.ReadInputLines())

	ops := []string{}
	offsets := []int{}
	for i, char := range lines[len(lines)-1] {
		if char == '*' || char == '+' {
			ops = append(ops, string(char))
			offsets = append(offsets, i)
		}
	}

	// Drop the operations line
	lines = lines[:len(lines)-1]

	// Preallocade grid of chars and ints
	sgrid := make([][]rune, len(lines))

	// Convert to a grid of chars
	for r, line := range lines {
		sgrid[r] = make([]rune, len(line))
		for c, ch := range line {
			sgrid[r][c] = ch
		}
	}

	/* sgrid [][]rune:
		123 328  51 64				 1	  369
		 45 64  387 23			=>	 24   248 ...
		  6 98  215 314			     356  8

	   o []string: "*", "+", "*", "+"
	*/

	g := [][]int{}
	for col, offset := range offsets {
		opNums := []int{}
		for colOffset := offset; (col == len(offsets)-1 && colOffset < len(sgrid[0])) || (col < len(offsets)-1 && colOffset < offsets[col+1]-1); colOffset++ {
			v := 0
			for _, row := range sgrid {
				if row[colOffset] != ' ' {
					v = v*10 + int(row[colOffset]-'0')
				}
			}

			opNums = append(opNums, v)
		}

		for c, n := range opNums {
			for len(g) <= c {
				g = append(g, make([]int, len(offsets)))
			}
			if g[c][col] != 0 {
				panic("unexpected non-zero value")
			}
			g[c][col] = n
		}
	}

	return g, ops
}
