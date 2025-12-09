package main

import (
	_ "embed"
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
	println("Part 2: ", part2_b(points))
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

func part2_b(points []lib.Pos2d) int {
	startT := time.Now()

	lib.Progress(startT, "Calculating bounds...")
	defer lib.ProgressDone()

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

	lib.Progress(startT, "Calculating boxes...")

	// Calculate every box that could potentially be valid
	boxes := make([][3]int, 0, len(points)*len(points)/2)
	for i1 := 0; i1 < len(points); i1++ {
		for i2 := i1 + 1; i2 < len(points); i2++ {
			x0, x1 := points[i1].X, points[i2].X
			if x0 > x1 {
				x0, x1 = x1, x0
			}
			y0, y1 := points[i1].Y, points[i2].Y
			if y0 > y1 {
				y0, y1 = y1, y0
			}

			// Exclude as many boxes as possible that are known to be invalid
			hasNested := false
			for i3 := 0; i3 < len(points) && !hasNested; i3++ {
				if i3 == i1 || i3 == i2 {
					continue
				}
				x2, y2 := points[i3].X, points[i3].Y

				if x0 <= x2 && x2 <= x1 && y0 <= y2 && y2 <= y1 {
					// Overlapping, but must allow edges and corners overlapping
					if !((x2 == x0 || x2 == x1) && (y2 == y0 || y2 == y1)) {
						hasNested = true
					}
				}
			}

			if !hasNested {
				boxes = append(boxes, [3]int{i1, i2, (x1 - x0 + 1) * (y1 - y0 + 1)})
			}
		}
	}

	lib.Progress(startT, "Sorting boxes...")
	slices.SortFunc(boxes, func(a, b [3]int) int {
		return b[2] - a[2]
	})

	lib.Progress(startT, "Building grid...")

	// Construct grid of max bounds
	// Add an extra +1 so debugging shows an extra columns of 0s
	g := make([][]byte, max_y+2)
	for y := 0; y < len(g); y++ {
		g[y] = make([]byte, max_x+2)
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

	lib.Progress(startT, "Filling grid...")

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

	for i, ar := range boxes {
		if i%10 == 0 {
			lib.Progress(startT, "Checking %d/%d boxes...", i, len(boxes))
		}

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
