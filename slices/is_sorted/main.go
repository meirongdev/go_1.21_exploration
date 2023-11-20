package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(slices.IsSorted([]int{0, 1, 2})) // true
	fmt.Println(slices.IsSorted([]int{0, 2, 1})) // false
	fmt.Println(slices.IsSorted([]int{2, 1, 0})) // false
}
