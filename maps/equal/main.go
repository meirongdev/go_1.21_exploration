package main

import (
	"fmt"
	"maps"
	"strings"
)

func main() {
	m1 := map[int]string{
		1:    "one",
		10:   "Ten",
		1000: "THOUSAND",
	}
	m2 := map[int][]byte{
		1:    []byte("One"),
		10:   []byte("Ten"),
		1000: []byte("Thousand"),
	}
	eq := maps.EqualFunc(m1, m2, func(v1 string, v2 []byte) bool {
		return strings.EqualFold(v1, string(v2))
	})
	// 比较value使用 传入的函数
	fmt.Println(eq)

	m11 := map[string]int{"one": 1, "two": 2}
	m22 := map[string]int{"one": 1, "two": 2}
	// 比较value使用 !=
	equal := maps.Equal(m11, m22) // true
	fmt.Println(equal)
}
