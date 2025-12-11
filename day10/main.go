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
	jolts    []uint8

	switchMaxes []uint8

	// A map of enabled switches and how many combinations of switches lead to that state
	combos map[switchState]int
}

func (s *schematics) String() string {
	return fmt.Sprintf("on: %016b switches: %v levels: %v", s.on, s.switches, s.jolts)
}

func readSchematics(input string) []*schematics {
	var result []*schematics

	for line := range lib.ReadLines(input) {
		chunks := strings.Fields(line)

		if len(chunks[0]) > 2+16 {
			panic("unexpected number of switches")
		}

		initState := chunks[0][1:strings.IndexRune(chunks[0], ']')]
		on := switchState(0)
		for n, c := range initState {
			if c == '#' {
				on |= 1 << (15 - n)
			}
		}

		joltsStr := chunks[len(chunks)-1]
		joltsStr = joltsStr[1 : len(joltsStr)-1]
		jolts := make([]uint8, len(initState))
		for n, num := range strings.Split(joltsStr, ",") {
			i, _ := strconv.Atoi(num)
			if i > 255 {
				panic("jolt level too high")
			}
			jolts[n] = uint8(i)
		}

		if len(chunks[1:len(chunks)-1]) > 16 {
			panic("unexpected number of switch groups")
		}

		switches := make([]switchState, len(chunks)-2)
		switchMaxes := make([]uint8, len(switches))
		for i := range len(switchMaxes) {
			switchMaxes[i] = 255
		}
		combos := make(map[switchState]int, len(chunks)<<1)
		for sn, flips := range chunks[1 : len(chunks)-1] {
			s := switchState(0)
			for _, num := range strings.Split(strings.Trim(flips, "()"), ",") {
				n, _ := strconv.Atoi(num)

				m := jolts[n]
				if m < switchMaxes[sn] {
					switchMaxes[sn] = m
				}

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
			jolts:    jolts,

			combos:      combos,
			switchMaxes: switchMaxes,
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

func part2(schematics []*schematics) int {
	result := 0

	for _, s := range schematics {
		result += part2_a(s)
	}

	return result
}

func part2_a(scm *schematics) int {
	result := 0

	for _, s := range scm.switches {
		for n := range scm.switchMaxes[s] {
			// can use switch 's' at most 'n' times
		}
	}

	return result
}
