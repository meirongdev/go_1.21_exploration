package main

import (
	"fmt"
	"slices"
)

func main() {
	originalNames := []string{"Alice", "Bob", "Vera", "Zac"}
	fmt.Printf("%v\n", originalNames)

	// 替换[1, 3) 之间的元素，替换的元素个数没必要和被替换的元素个数相同
	names := slices.Replace(originalNames, 1, 3, "Bill", "Billie", "Cat")
	fmt.Printf("%v\n", names)

	names = slices.Replace(originalNames, 0, 0, "Bill", "Billie", "Cat")
	fmt.Printf("%v\n", names)

	// names = slices.Replace(originalNames, 0, -1, "Bill", "Billie", "Cat")
	// fmt.Printf("%v\n", names) // panic: runtime error: slice bounds out of range [:-1]
}
