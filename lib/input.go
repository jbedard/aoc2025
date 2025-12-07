package lib

import (
	"iter"
	"strings"
)

func ReadLines(content string) iter.Seq[string] {
	return strings.SplitSeq(strings.Trim(content, "\n"), "\n")
}
