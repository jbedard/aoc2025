package main

import (
	_ "embed"
	"slices"

	"github.com/jbedard/aoc2025/lib"
)

//go:embed input.txt
var content string

func main() {
	points := []lib.Pos3d{}
	for line := range lib.ReadLines(content) {
		points = append(points, lib.ReadPos3d(line))
	}
	connectCount := 1000

	defer lib.CpuProfile()()

	p1, p2 := part1(points, connectCount)
	println("Part 1:", p1)
	println("Part 2:", p2)
}

type PointPairDist = [3]int

func calculatePairs(points []lib.Pos3d) []PointPairDist {
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
	return distances
}

func part1(points []lib.Pos3d, connectCount int) (int, int) {
	// Calculate distances between all points in shortest order
	distances := calculatePairs(points)

	// Start with each point in its own circuit
	circuitCount := len(points)
	pointCircuits := make([]int, circuitCount)
	circuits := make([][]int, circuitCount)
	for i := range pointCircuits {
		pointCircuits[i] = -1
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

		if c1 == -1 && c2 == -1 {
			// Neither point is in a circuit: create a new circuit
			circuits[i1] = []int{i1, i2}
			pointCircuits[i1] = i1
			pointCircuits[i2] = i1
			circuitCount--
		} else if c1 == -1 {
			// Point 1 is not in a circuit: add it to point 2's circuit
			circuits[c2] = append(circuits[c2], i1)
			pointCircuits[i1] = c2
			circuitCount--
		} else if c2 == -1 {
			// Point 2 is not in a circuit: add it to point 1's circuit
			circuits[c1] = append(circuits[c1], i2)
			pointCircuits[i2] = c1
			circuitCount--
		} else if c1 != c2 {
			// Not already in the same circuit: merge two circuits
			fromId := c2
			toId := c1
			for _, pi := range circuits[fromId] {
				pointCircuits[pi] = toId
				circuits[toId] = append(circuits[toId], pi)
			}
			circuitCount--
			circuits[fromId] = nil
		}

		// TODO: why are we incrementing when we may not have connected anything!?
		cc++
		part2Result = d
		if cc == connectCount {
			part1Circuites := slices.Clone(circuits)
			slices.SortFunc(part1Circuites, func(a, b []int) int {
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
