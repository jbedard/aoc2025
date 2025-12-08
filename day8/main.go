package main

import (
	_ "embed"
	"slices"

	"github.com/jbedard/aoc2025/lib"
)

//go:embed input.txt
var content string

func main() {
	points := []lib.Pos{}
	for line := range lib.ReadLines(content) {
		points = append(points, lib.ReadPoint(line))
	}
	connectCount := 1000

	r1 := part1(points, connectCount)
	println("Part 1:", r1)
}

type PointPairDist = [3]int

func part1(points []lib.Pos, connectCount int) int {
	// Calculate distances between all points in shortest order
	distances := []PointPairDist{}
	for i1 := 0; i1 < len(points); i1++ {
		p1 := points[i1]
		for i2 := i1 + 1; i2 < len(points); i2++ {
			p2 := points[i2]
			distances = append(distances, PointPairDist{i1, i2, p1.Dist(p2)})
		}
	}
	slices.SortStableFunc(distances, func(a, b PointPairDist) int {
		return a[2] - b[2]
	})

	// Start with each point in its own circuit
	pointCircuits := make([]int, len(points))
	circuits := make([][]int, len(points))
	for i := range pointCircuits {
		pointCircuits[i] = i
		circuits[i] = []int{i}
	}

	cc := 0

	// Merge the closest N points
	for _, d := range distances {
		// Indexes of the points
		i1, i2 := d[0], d[1]

		// Circuits of the points
		c1, c2 := pointCircuits[i1], pointCircuits[i2]

		if c1 != c2 {
			// Not already in the same circuit: merge two circuits
			fromId := c2
			toId := c1
			for _, pi := range circuits[fromId] {
				pointCircuits[pi] = toId
				for _, pi2 := range circuits[fromId] {
					if !slices.Contains(circuits[toId], pi2) {
						circuits[toId] = append(circuits[toId], pi2)
					}
				}
			}
			circuits[fromId] = nil
		}

		// TODO: why are we incrementing when we may not have connected anything!?
		cc++
		if cc >= connectCount {
			break
		}
	}

	slices.SortStableFunc(circuits, func(a, b []int) int {
		return len(b) - len(a)
	})

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}
