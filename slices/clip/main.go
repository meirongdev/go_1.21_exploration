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
	// [1 2 3] len=3 cap=10 0xc0000b4050
	afterClip := slices.Clip(s)
	fmt.Printf("%v len=%d cap=%d %p\n", afterClip, len(afterClip), cap(afterClip), afterClip)
	// output
	// [1 2 3] len=3 cap=3 0xc0000b4050
}
