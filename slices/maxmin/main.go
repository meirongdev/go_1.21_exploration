package main

import (
	"slices"
)

func main() {
	numbers := []int{}
	slices.Max(numbers) // panic: slices.Max: empty list
	slices.Min(numbers)
}
