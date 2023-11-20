package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strings"
)

func main() {
	// sort floating with NaN
	nums := []float64{1.0, 2.0, 3.0, 4.0, 5.0, math.NaN(), 6.0, 7.0}
	slices.Sort(nums)
	fmt.Printf("%v\n", nums)

	// sort func
	names := []string{"Bob", "alice", "VERA"}
	slices.SortFunc(names, func(a, b string) int {
		return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	fmt.Println(names)

	// sort stable
	names = []string{"Bob", "alice", "VERA"}
	slices.SortStableFunc(names, func(a, b string) int {
		return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	fmt.Println(names)
}
