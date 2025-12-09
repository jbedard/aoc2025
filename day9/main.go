package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/jbedard/aoc2025/lib"
)

//go:embed input.txt
var content string

func main() {
	points := []lib.Pos2d{}
	for line := range lib.ReadLines(strings.TrimSpace(content)) {
		points = append(points, lib.ReadPos2d(line))
	}

	defer lib.CpuProfile()()

	println("Part 1: ", part1(points))
	println("Part 2: ", part2_a(points))
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

func part2_a(points []lib.Pos2d) int {
	startT := time.Now()

	fmt.Printf("Start: %v\n", time.Since(startT))

	// Calculate the bounds
	min_x, max_x := math.MaxInt, 0
	min_y, max_y := math.MaxInt, 0
	for _, p := range points {
		if p.X < min_x {
			min_x = p.X
		}
		if p.X > max_x {
			max_x = p.X
		}
		if p.Y < min_y {
			min_y = p.Y
		}
		if p.Y > max_y {
			max_y = p.Y
		}
	}

	// Construct grid of max bounds
	g := make([][]byte, max_y+1)
	for y := 0; y <= max_y; y++ {
		g[y] = make([]byte, max_x+2) // Add an extra +1 so debugging shows an extra columns of 0s
	}

	// Fill between the points
	for pi := 1; pi <= len(points); pi++ {
		p0 := points[pi-1]
		p1 := points[pi%len(points)]

		if p0.X == p1.X {
			y0, y1 := p0.Y, p1.Y
			if y0 > y1 {
				y0, y1 = y1, y0
			}
			for y := y0; y <= y1; y++ {
				g[y][p0.X]++
			}
		} else if p0.Y == p1.Y {
			x0, x1 := p0.X, p1.X
			if x0 > x1 {
				x0, x1 = x1, x0
			}
			for x := x0; x <= x1; x++ {
				g[p0.Y][x]++
			}
		} else {
			panic("Bad input!")
		}
	}

	// Fill the insides
	for y := min_y; y <= max_y; y++ {
		acc := byte(0)
		for x := min_x; x <= max_x; x++ {
			if acc < g[y][x] {
				for ; x <= max_x; x++ {
					g[y][x]++
				}
			}
		}
	}

	fmt.Printf("Grid built: %v\n", time.Since(startT))

	areas := make([][3]int, 0, len(points)*len(points)/2)
	for i1 := 1; i1 < len(points); i1++ {
		for i2 := i1 + 1; i2 < len(points); i2++ {
			areas = append(areas, [3]int{i1, i2, area(points[i1], points[i2])})
		}
	}
	slices.SortFunc(areas, func(a, b [3]int) int {
		return b[2] - a[2]
	})

	fmt.Printf("Areas calculated: %v\n", time.Since(startT))

	for _, ar := range areas {
		i1, i2, a := ar[0], ar[1], ar[2]
		if isBlockInGrid(g, points[i1], points[i2]) {
			return a
		}
	}

	panic("No solution found!")
}

func isBlockInGrid(g [][]byte, p1 lib.Pos2d, p2 lib.Pos2d) bool {
	x0, x1 := p1.X, p2.X
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	y0, y1 := p1.Y, p2.Y
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	for y := y0; y <= y1; y++ {
		for x := x0; x <= x1; x++ {
			if g[y][x] == 0 {
				return false
			}
		}
	}
	return true
}
