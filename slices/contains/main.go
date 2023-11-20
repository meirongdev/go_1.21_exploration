package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []string{"Apple", "Banana", "Orange"}
	fmt.Println("Have Apple?", slices.Contains(s, "Apple"))
}
