package lib

import (
	"fmt"
	"time"
)

func Progress(t time.Time, m string, args ...interface{}) {
	fmt.Print("\x1b7")     // save the cursor position
	fmt.Print("\x1b[2K\r") // erase the current line
	fmt.Printf("\x1B[33m %s: %s\x1B[0m", time.Since(t), fmt.Sprintf(m, args...))
	fmt.Print("\x1b8") // restore the cursor position
}
func ProgressDone() {
	fmt.Print("\x1b7")     // save the cursor position
	fmt.Print("\x1b[2K\r") // erase the current line
	fmt.Print("\x1b8")     // restore the cursor position
}
