package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("==============sync.OnceValue")
	onlyOnce := sync.OnceValue(func() int {
		// Expensive code here
		fmt.Println("OnceValue called")
		return 42 // Return your computed int
	})

	// 第一次调用的时候才执行初始化。
	firstCall := onlyOnce()
	fmt.Printf("%d\n", firstCall)
	secondCall := onlyOnce()
	fmt.Printf("%d\n", secondCall)

	fmt.Println("==============sync.OnceFunc")
	type person struct {
		name string
		age  int
	}
	var lazyObj person
	initFunc := sync.OnceFunc(func() {
		fmt.Println("OnceFunc called")
		lazyObj = person{
			name: "lazy",
			age:  18,
		}
	})

	initFunc()
	initFunc()
	fmt.Printf("%+v\n", lazyObj)

	fmt.Println("==============sync.OnceValues")
	onlyOnceValues := sync.OnceValues(func() (int, string) {
		fmt.Println("OnceValues called")
		return 1, "hello"
	})

	a, b := onlyOnceValues()
	fmt.Printf("%d, %s\n", a, b)
	a, b = onlyOnceValues()
	fmt.Printf("%d, %s\n", a, b)
}
