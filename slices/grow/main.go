package main

import (
	"fmt"
	"slices"
)

func main() {
	s := make([]int, 0, 10)
	s = append(s, 1, 2, 3)
	fmt.Printf("%v len=%d cap=%d %p\n", s, len(s), cap(s), s)
	// output
	// [1 2 3] len=3 cap=10 0xc000120050
	afterGrow := slices.Grow(s, 20)
	fmt.Printf("%v len=%d cap=%d %p\n", afterGrow, len(afterGrow), cap(afterGrow), afterGrow)
	// output
	// [1 2 3] len=3 cap=24 0xc00012a000
}
