package main

import (
	"fmt"
	"maps"
)

func main() {
	// copy
	dst := map[string]int{"one": 1, "three": 3}
	src := map[string]int{"two": 2, "three": 33}
	maps.Copy(dst, src)
	fmt.Printf("%v\n", dst) // map[one:1 three:33 two:2]

	// clone
	m := map[string]int{"one": 1, "two": 2}
	fmt.Printf("m: %v %d %p\n", m, len(m), &m) // map[one:1 two:2] 2 0xc000056020
	cloned := maps.Clone(m)
	fmt.Printf("cloned: %v %d %p\n", cloned, len(cloned), &cloned) // map[one:1 two:2] 2 0xc000056030
}
