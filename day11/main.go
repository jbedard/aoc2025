package main

import (
	_ "embed"
	"strings"

	"github.com/jbedard/aoc2025/lib"
)

//go:embed input.txt
var content string

type node string
type path []node

func parseInput(input string) map[node][]node {
	result := make(map[node][]node)

	for line := range lib.ReadLines(input) {
		parts := strings.SplitN(line, ": ", 2)
		name := parts[0]
		deps := strings.Fields(parts[1])
		result[node(name)] = make([]node, len(deps))
		for n, d := range deps {
			result[node(name)][n] = node(d)
		}
	}

	return result
}

func main() {
	d := parseInput(content)

	r1 := part1(d)
	println("Part 1:", r1)
}

func part1(d map[node][]node) int {
	visited := make(map[node][]path)
	part1_visit(d, visited, []node{}, "you")
	return len(visited["out"])
}

func part1_visit(d map[node][]node, visited map[node][]path, p []node, todo node) {
	here := append(p[:], todo)
	visited[todo] = append(visited[todo], here)
	for _, next := range d[todo] {
		part1_visit(d, visited, here, next)
	}
}
