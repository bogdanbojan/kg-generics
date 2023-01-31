package print

import (
	"fmt"
	"io"
)

// Your PrintAnythingTo function goes here!
func PrintAnythingTo[T any](w io.Writer, v T) {
	fmt.Fprintln(w, v)
}
