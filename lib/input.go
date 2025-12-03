package lib

import (
	"bufio"
	"io"
	"iter"
	"os"
	"strings"
)

func ReadInputList() []string {
	r, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	content, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimSpace(string(content)), ",")
}

func ReadInputLines() iter.Seq[string] {
	return func(yield func(string) bool) {
		r, err := os.Open("./input.txt")
		if err != nil {
			panic(err)
		}
		defer r.Close()

		scanner := bufio.NewScanner(r)

		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
}
