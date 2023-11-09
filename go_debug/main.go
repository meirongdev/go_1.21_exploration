//go:debug panicnil=1
package main

import "fmt"

// will output
// ok

// compare to ../panic_nil/main.go
//
//	 which outputs
//	panic: panic called with nil argument

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
			return
		}
		fmt.Println("ok")
	}()
	panic(nil)
}
