package main

import (
	"fmt"
	"strings"

	_ "embed"
)

//go:embed input.txt
var content string

type present uint16

func newPresent(p uint16) present {
	if p&(^presentMask) != 0 {
		panic("invalid present")
	}
	return present(p)
}

const presentMask = uint16(0b0000000111111111)

func (p present) size() int {
	n := 0
	for i := 0; i < 9; i++ {
		if (p & (1 << i)) != 0 {
			n++
		}
	}
	return n
}
func (p present) String() string {
	s := ""
	for i := 0; i < 9; i++ {
		if (p & (1 << (8 - i))) != 0 {
			s += "#"
		} else {
			s += "."
		}
		if i%3 == 2 && i != 8 {
			s += "\n"
		}
	}
	return s
}

type region struct {
	width, height int
	presentCounts []int
}

type dataSet struct {
	regions []region

	presents []present
}

func parsePresent(input string) present {
	block := strings.ReplaceAll(input, "\n", "")
	if len(block) != 9 {
		panic("invalid present block")
	}
	p := uint16(0)
	for _, c := range block {
		p <<= 1
		if c == '#' {
			p |= 1
		}
	}
	np := newPresent(p)
	nps := np.String()
	if strings.TrimSpace(input) != nps {
		panic("present parse/string mismatch")
	}
	return np
}

func parseInput(input string) dataSet {
	presents := []present{}
	regions := []region{}
	for _, block := range strings.Split(strings.TrimSpace(input), "\n\n") {
		if strings.Contains(block, "#") {
			presents = append(presents, parsePresent(strings.TrimSpace(block[strings.Index(block, "\n")+1:])))
		} else {
			for line := range strings.SplitSeq(block, "\n") {
				var x, y int
				ss := make([]int, len(presents))
				fmt.Sscanf(line, "%dx%d: %d %d %d %d %d %d", &x, &y, &ss[0], &ss[1], &ss[2], &ss[3], &ss[4], &ss[5])

				regions = append(regions, region{
					width:         x,
					height:        y,
					presentCounts: ss,
				})
			}
		}
	}

	return dataSet{
		presents: presents,
		regions:  regions,
	}
}

func main() {
	ds := parseInput(content)
	r1 := part1(ds)

	println("Part 1:", r1)
}

func part1(ds dataSet) int {
	r := 0

	for _, region := range ds.regions {
		if regionFits(region, ds.presents) {
			r++
		}
	}

	return r
}

func regionFits(r region, presents []present) bool {
	n := 0
	for i, pc := range r.presentCounts {
		s := presents[i].size()
		n += pc * s
	}
	return r.width*r.height >= n
}
