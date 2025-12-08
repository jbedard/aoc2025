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

	defer lib.CpuProfile()()

	p1, p2 := part1(points, connectCount)
	println("Part 1:", p1)
	println("Part 2:", p2)
}

type PointPairDist = [3]int

func part1(points []lib.Pos, connectCount int) (int, int) {
	// Calculate distances between all points in shortest order
	distances := make([]PointPairDist, 0, len(points)*len(points)/2)
	for i1 := 0; i1 < len(points); i1++ {
		p1 := points[i1]
		for i2 := i1 + 1; i2 < len(points); i2++ {
			p2 := points[i2]
			distances = append(distances, PointPairDist{i1, i2, p1.Dist(p2)})
		}
	}
	slices.SortFunc(distances, func(a, b PointPairDist) int {
		return a[2] - b[2]
	})

	// Start with each point in its own circuit
	circuitCount := len(points)
	pointCircuits := make([]int, circuitCount)
	circuits := make([]map[int]struct{}, circuitCount)
	for i := range pointCircuits {
		pointCircuits[i] = i
		circuits[i] = map[int]struct{}{i: {}}
	}

	part1Answer := -1
	part2Result := distances[0]
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
			for pi := range circuits[fromId] {
				pointCircuits[pi] = toId
				for pi2 := range circuits[fromId] {
					circuits[toId][pi2] = struct{}{}
				}
			}
			circuitCount--
			circuits[fromId] = nil
			part2Result = d
		}

		// TODO: why are we incrementing when we may not have connected anything!?
		cc++
		if cc == connectCount {
			part1Circuites := slices.Clone(circuits)
			slices.SortFunc(part1Circuites, func(a, b map[int]struct{}) int {
				return len(b) - len(a)
			})
			part1Answer = len(part1Circuites[0]) * len(part1Circuites[1]) * len(part1Circuites[2])
		}

		if circuitCount == 1 {
			break
		}
	}

	return part1Answer, (points[part2Result[0]].X * points[part2Result[1]].X)
}
