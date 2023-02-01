package stringy

import (
	"fmt"
	"io"
)

// Your StringifyTo function goes here!
func StringifyTo[T fmt.Stringer](w io.Writer, v T) {
	fmt.Fprintln(w, v)
}
