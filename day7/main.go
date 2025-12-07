package main

import "github.com/jbedard/aoc2025/lib"

func main() {
	g := lib.NewCharGridFromSeq(lib.ReadInputLines())

	count, forks := part1(g)
	println("Part 1:", count)
	println("Part 2:", forks)
}

func part1(g lib.Grid[rune]) (int, int) {
	beams := make([]int, g.W())
	splits := 0
	for _, row := range g.Rows() {
		for x, cell := range row {
			switch cell {
			case 'S':
				beams[x]++
			case '^':
				if c := beams[x]; c != 0 {
					beams[x] = 0
					if 0 < x {
						beams[x-1] += c
					}
					if x < g.W()-1 {
						beams[x+1] += c
					}
					splits++
				}
			}
		}
	}

	beamCount := 0
	for _, b := range beams {
		beamCount += b
	}

	return splits, beamCount
}
