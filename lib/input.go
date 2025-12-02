package lib

import (
	"io"
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
