package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Start Ingredient
	End   Ingredient
}

type RangeSet []Range

func (rs RangeSet) Contains(i Ingredient) bool {
	for _, r := range rs {
		if r.Start <= i && i <= r.End {
			return true
		}
	}
	return false
}

type Ingredient = int64

func parseInput() (RangeSet, []Ingredient) {
	r, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)

	doneRanges := false
	ranges := RangeSet{}
	items := []Ingredient{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			doneRanges = true
			continue
		}

		if doneRanges {
			n, _ := strconv.ParseInt(line, 10, 64)
			items = append(items, n)
		} else {
			parts := strings.SplitN(line, "-", 2)
			start, _ := strconv.ParseInt(parts[0], 10, 64)
			end, _ := strconv.ParseInt(parts[1], 10, 64)

			ranges = append(ranges, Range{Start: start, End: end})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		if a.Start < b.Start {
			return -1
		} else if a.Start > b.Start {
			return 1
		} else {
			return 0
		}
	})

	return ranges, items
}

func main() {
	fresh, available := parseInput()

	part1(fresh, available)
}

func part1(fresh RangeSet, available []Ingredient) {
	c := 0

	for _, item := range available {
		if fresh.Contains(item) {
			c++
		}
	}

	println("Part 1 fresh count:", c)
}
