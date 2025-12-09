package main

import (
	_ "embed"
	"math"
	"strings"

	"github.com/jbedard/aoc2025/lib"
)

//go:embed input.txt
var content string

func main() {
	points := []lib.Pos2d{}
	for line := range lib.ReadLines(strings.TrimSpace(content)) {
		points = append(points, lib.ReadPos2d(line))
	}

	println("Part 1: ", part1(points))
}

func area(p1 lib.Pos2d, p2 lib.Pos2d) int {
	return int((math.Abs(float64(p1.X - p2.X + 1))) * (math.Abs(float64(p1.Y - p2.Y + 1))))
}

func part1(points []lib.Pos2d) int {
	p1, p2 := points[0], points[1]
	a := area(p1, p2)

	for i1 := 1; i1 < len(points); i1++ {
		for i2 := i1 + 1; i2 < len(points); i2++ {
			a2 := area(points[i1], points[i2])
			if a2 > a {
				p1, p2 = points[i1], points[i2]
				a = a2
			}
		}
	}

	return a
}
