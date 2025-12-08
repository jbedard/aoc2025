package lib

import (
	"math"
	"strconv"
	"strings"
)

type Pos struct {
	X, Y, Z int
}

func (p1 Pos) Dist(p2 Pos) int {
	return int(math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y) + (p1.Z-p2.Z)*(p1.Z-p2.Z))))
}

func (p1 Pos) MDist(p2 Pos) int {
	return int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)) + math.Abs(float64(p1.Z-p2.Z)))
}

func ReadPoint(s string) Pos {
	a := strings.SplitN(s, ",", 3)
	x, _ := strconv.Atoi(a[0])
	y, _ := strconv.Atoi(a[1])
	z, _ := strconv.Atoi(a[2])
	return Pos{
		X: x,
		Y: y,
		Z: z,
	}
}
