package main

import "github.com/jbedard/aoc2025/lib"

func main() {
	g := lib.NewCharGridFromSeq(lib.ReadInputLines())

	println("Part 1:", part1(g))
}

func part1(g lib.Grid[rune]) int {
	beams := make([]bool, g.W())
	splits := 0
	for _, row := range g.Rows() {
		for x, cell := range row {
			switch cell {
			case 'S':
				beams[x] = true
			case '^':
				if beams[x] {
					beams[x] = false
					if 0 < x {
						beams[x-1] = true
					}
					if x < g.W()-1 && !beams[x+1] {
						beams[x+1] = true
					}
					splits++
				}
			}
		}
	}
	return splits
}
