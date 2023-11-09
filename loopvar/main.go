package main

import (
	"fmt"
	"time"
)

func main() {
	var m = map[string]int{"a": 1, "b": 2}
	// for var pitfall
	fmt.Println("for var pitfall")
	for k, v := range m {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("k: %v, v: %v\n", k, v)
		}()
	}
	// output without GOEXPERIMENT=loopvar
	// k: b, v: 2
	// k: b, v: 2
	// output with GOEXPERIMENT=loopvar
	// k: a, v: 1
	// k: b, v: 2

	fmt.Println("avoid for var pitfall, use func params")
	for k, v := range m {
		go func(k string, v int) {
			time.Sleep(1 * time.Second)
			fmt.Printf("k: %v, v: %v\n", k, v)
		}(k, v)
	}
	// output
	// k: a, v: 1
	// k: b, v: 2

	fmt.Println("avoid for var pitfall, use local var")
	for k, v := range m {
		var k1, v1 = k, v
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("k: %v, v: %v\n", k1, v1)
		}()
	}
	// output
	// k: a, v: 1
	// k: b, v: 2
	time.Sleep(3 * time.Second)
}
