package main

import "fmt"

// 在Go1.21中，输出 panic: panic called with nil argument
// 在Go1.20中，输出 ok

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
