package main

import "github.com/jbedard/aoc2025/lib"

func main() {
	g := lib.NewCharGridFromSeq(lib.ReadInputLines())
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

	println("Count:", c)
}

func canMove(g lib.Grid[rune], x, y int) bool {
	r := g.At(x, y)
	if r != '@' {
		return false
	}

	return g.CountAround(x, y, func(x, y int, r rune) bool { return r == '@' }) < 4
}
