package main

import (
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
}

func part1(g [][]int, o []string) int {
	t := 0

	for i, op := range o {
		v := 0
		if op == "*" {
			v = 1
		}

		for _, row := range g {
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
