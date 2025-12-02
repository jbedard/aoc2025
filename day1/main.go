package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	content, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	d := 50
	zeroCount := 0

	for line := range strings.SplitSeq(strings.TrimSpace(string(content)), "\n") {
		n, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			d = (d - n) % 100
		} else {
			d = (d + n) % 100
		}

		if d%100 == 0 {
			zeroCount++
		}
	}
	fmt.Printf("Zero count: %d\n", zeroCount)
}
