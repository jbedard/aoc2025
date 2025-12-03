package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jbedard/aoc2025/lib"
)

func main() {
	t := 0
	for s := range lib.ReadInputLines() {
		t += maxJoltage(s)
	}

	fmt.Printf("Total: %d\n", t)
}

func maxJoltage(s string) int {
	t := 0
	p := 2
	for d := byte('9'); d >= '0' && p > 0; {
		if di := strings.IndexByte(s, d); di != -1 && (len(s)-di >= p) {
			// Add this digit's place value to total
			t += int(math.Pow10(p-1)) * int(d-'0')

			p--           // the next digit is 1/10 the place value
			s = s[di+1:]  // truncate string to after found digit
			d = byte('9') // look for 9s again
		} else {
			d--
		}
	}
	return t
}
