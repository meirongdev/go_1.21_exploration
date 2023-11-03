package main

import (
	_ "github.com/meirongdev/go121expl/pkg_init/b"

	_ "github.com/meirongdev/go121expl/pkg_init/a"
)

// main -> a
// main -> b
// a -> c
// a -> d
// b -> c
// b -> e
// c -> e
// c -> f
// e -> f
// output:
// d init
// f init
// e init
// c init
// a init
// b init
// main init
func init() {
	println("main init")
}

func main() {
}
