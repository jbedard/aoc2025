package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	d := 50
	zeroCount := 0

	for line := range strings.SplitSeq(strings.TrimSpace(content), "\n") {
		n, _ := strconv.Atoi(line[1:])

		for i := 0; i < n; i++ {
			if line[0] == 'L' {
				d--
			} else {
				d++
			}

			if d%100 == 0 {
				zeroCount++
			}
		}

		d = d % 100
	}
	fmt.Printf("Zero count: %d\n", zeroCount)
}
