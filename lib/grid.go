package lib

import (
	"iter"
)

type Coord struct {
	X int
	Y int
}

type Grid[T any] interface {
	W() int
	H() int
	At(x, y int) T

	Set(x, y int, r rune)
	Matches(pred func(int, int, T) bool) iter.Seq[Coord]
	CountMatches(pred func(int, int, T) bool) int
	CountAround(x, y int, pred func(int, int, T) bool) int
}

type charGrid struct {
	rows [][]rune
	w    int
}

var _ Grid[rune] = (*charGrid)(nil)

func NewCharGridFromSeq(seq iter.Seq[string]) Grid[rune] {
	rows := [][]rune{}
	w := -1
	for line := range seq {
		if w == -1 {
			w = len(line)
		} else if len(line) != w {
			panic("inconsistent row widths")
		}
		rows = append(rows, []rune(line))
	}

	return &charGrid{rows: rows, w: w}
}

func (g *charGrid) Set(x, y int, r rune) {
	g.rows[y][x] = r
}

func (g *charGrid) At(x, y int) rune {
	return g.rows[y][x]
}

func (g *charGrid) W() int {
	return g.w
}

func (g *charGrid) H() int {
	return len(g.rows)
}

func (g *charGrid) Matches(pred func(int, int, rune) bool) iter.Seq[Coord] {
	return func(yield func(Coord) bool) {
		for y := 0; y < g.H(); y++ {
			for x := 0; x < g.W(); x++ {
				if pred(x, y, g.At(x, y)) {
					if !yield(Coord{X: x, Y: y}) {
						return
					}
				}
			}
		}
	}
}

func (g *charGrid) CountMatches(pred func(int, int, rune) bool) int {
	count := 0
	for y := 0; y < g.H(); y++ {
		for x := 0; x < g.W(); x++ {
			if pred(x, y, g.At(x, y)) {
				count++
			}
		}
	}
	return count
}

func (g *charGrid) CountAround(x, y int, pred func(int, int, rune) bool) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < g.W() && ny >= 0 && ny < g.H() {
				if pred(nx, ny, g.At(nx, ny)) {
					count++
				}
			}
		}
	}
	return count
}
