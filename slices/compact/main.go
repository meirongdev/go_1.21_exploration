package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println("slices.Compact")
	seq := []int{0, 1, 1, 2, 1, 3, 5, 8}
	fmt.Printf("%v len(seq)=%d\n", seq, len(seq))
	seq = slices.Compact(seq)
	fmt.Printf("%v len(seq)=%d\n", seq, len(seq))
	// output
	// [0 1 1 2 1 3 5 8] len(seq)=8
	// [0 1 2 1 3 5 8] len(seq)=7

	fmt.Println("slices.CompactFunc")
	names := []string{"bob", "Bob", "alice", "bob", "Vera", "VERA"}
	fmt.Printf("%v len(names)=%d\n", names, len(names))
	names = slices.CompactFunc(names, func(a, b string) bool {
		return strings.ToLower(a) == strings.ToLower(b)
	})
	fmt.Printf("%v len(names)=%d\n", names, len(names))
	// output
	// [bob Bob alice bob Vera VERA] len(names)=6
	// [bob alice bob Vera] len(names)=4
}
