package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jbedard/aoc2025/lib"
)

func main() {
	c := addInvalidIds(lib.ReadInputList())
	fmt.Printf("Invalid invalids total: %d\n", c)
}

func addInvalidIds(ids []string) int {
	s := 0
	for _, item := range ids {
		rng := strings.Split(item, "-")
		if len(rng) != 2 {
			panic(fmt.Sprintf("invalid range: %s from %v", item, ids))
		}
		low, _ := strconv.Atoi(rng[0])
		high, _ := strconv.Atoi(rng[1])
		s += addRangeInvalidIds(low, high)
	}
	return s
}

func addRangeInvalidIds(low, high int) int {
	n := 0
	for i := low; i <= high; i++ {
		if isInvalidId(i) {
			n += i
		}
	}
	return n
}

func isInvalidId(id int) bool {
	s := strconv.Itoa(id)
	return len(s)%2 == 0 && s[0:len(s)/2] == s[len(s)/2:]
}
