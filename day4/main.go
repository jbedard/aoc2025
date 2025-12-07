package main

import (
	_ "embed"

	"github.com/jbedard/aoc2025/lib"
)

//go:embed input.txt
var content string

func main() {
	g := lib.NewCharGridFromSeq(lib.ReadLines(content))
	c := solution2(g)

	println("Count:", c)
}

func solution2(g lib.Grid[rune]) int {
	c := 0
	todo := []lib.Coord{}
	for y := 0; y < g.H(); y++ {
		for x := 0; x < g.W(); x++ {
			todo = append(todo, lib.Coord{X: x, Y: y})
		}
	}

	for len(todo) > 0 {
		now_todo := make([]lib.Coord, 0, len(todo))

		for _, coord := range todo {
			if canMove(g, coord.X, coord.Y) {
				g.Set(coord.X, coord.Y, '.')
				c++

				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						nx := coord.X + dx
						ny := coord.Y + dy
						if nx >= 0 && nx < g.W() && ny >= 0 && ny < g.H() && !(dx == 0 && dy == 0) {
							now_todo = append(now_todo, lib.Coord{X: nx, Y: ny})
						}
					}
				}
			}
		}

		todo = now_todo
	}

	return c
}

func solution1(g lib.Grid[rune]) int {
	c := 0
	for {
		moved := false

		for coord := range g.Matches(func(x, y int, r rune) bool {
			return canMove(g, x, y)
		}) {
			moved = true
			g.Set(coord.X, coord.Y, '.')
			c++
		}

		if !moved {
			break
		}
	}
	return c
}

func canMove(g lib.Grid[rune], x, y int) bool {
	r := g.At(x, y)
	if r != '@' {
		return false
	}

	return g.CountAround(x, y, func(x, y int, r rune) bool { return r == '@' }) < 4
}
