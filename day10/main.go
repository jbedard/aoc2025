package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/jbedard/aoc2025/lib"
)

//go:embed input.txt
var content string

type switchState uint16

func (s switchState) String() string {
	return fmt.Sprintf("%016b", uint16(s))
}

type schematics struct {
	on       switchState
	switches []switchState

	// A map of enabled switches and how many combinations of switches lead to that state
	combos map[switchState]int
}

func (s *schematics) String() string {
	return fmt.Sprintf("on: %016b switches: %v", s.on, s.switches)
}

func readSchematics(input string) []*schematics {
	var result []*schematics

	for line := range lib.ReadLines(input) {
		chunks := strings.Fields(line)

		if len(chunks[0]) > 2+16 {
			panic("unexpected number of switches")
		}

		on := switchState(0)
		for n, c := range chunks[0][1:strings.IndexRune(chunks[0], ']')] {
			if c == '#' {
				on |= 1 << (15 - n)
			}
		}

		if len(chunks[1:len(chunks)-1]) > 16 {
			panic("unexpected number of switch groups")
		}

		switches := make([]switchState, len(chunks)-2)
		combos := make(map[switchState]int, len(chunks)<<1)
		for sn, flips := range chunks[1 : len(chunks)-1] {
			s := switchState(0)
			for _, num := range strings.Split(strings.Trim(flips, "()"), ",") {
				n, _ := strconv.Atoi(num)
				s |= 1 << (15 - n)
			}

			switches[sn] = s

			combos[s] = 1
			for combo, count := range combos {
				combined := combo ^ s
				if existing, found := combos[combined]; !found || count+1 < existing {
					combos[combined] = count + 1
				}
			}
		}

		result = append(result, &schematics{
			on:       on,
			switches: switches,
			combos:   combos,
		})
	}

	return result
}

func main() {
	schematics := readSchematics(content)

	fmt.Printf("Part 1: %d\n", part1(schematics))
}

func part1(schematics []*schematics) int {
	result := 0
	for _, s := range schematics {
		if count, found := s.combos[s.on]; found {
			result += count
		} else {
			panic("no solution found: " + s.String())
		}
	}
	return result
}
