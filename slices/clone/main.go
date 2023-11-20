package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("%v %p\n", s, s)
	// output
	// [1 2 3] 0xc0000b4050
	clone := slices.Clone(s)
	fmt.Printf("%v %p\n", clone, clone)
	// output
	// [1 2 3] 0xc0000b40a0
	type Person struct {
		Name string
	}
	people := []Person{
		{Name: "bob"},
		{Name: "alice"},
	}
	clonePeople := slices.Clone(people)
	fmt.Printf("clonePeople %v %p\n", clonePeople, clonePeople)
	// output
	// [{bob} {alice}] 0xc0000b40e0
	clonePeople[0].Name = "Bob"
	fmt.Printf("clonePeople %v %p\n", clonePeople, clonePeople)
	// output
	// [{Bob} {alice}] 0xc0000b40e0
	fmt.Printf("people %v %p\n", people, people)
	// output
	// [{Bob} {alice}] 0xc0000b40e0
}
