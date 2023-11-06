package main

import (
	"fmt"
	"slices"
	"strings"
)

func Clone1[E any](s []E) []E {
	return append(s[:0:0], s...)
}

// We want the argument type is the same with the return type
// 此处?仅是占位符
// func Clone2[S ?](s S) S {
// 	return s
// }

// 我们知道S应该是个Slice，但是还需要定义E
// func Clone3[S []E](s S) S {
// 	return s
// }

// S是个Slice，且E是任意类型
func Clone4[S []E, E any](s S) S {
	return append(s[:0:0], s...)
}

func Clone5[S ~[]E, E any](s S) S {
	return append(s[:0:0], s...)
}

type MySlice []string

func (ms MySlice) String() string {
	return strings.Join(ms, "+")
}

func PrintSorted(ms MySlice) string {
	// c := Clone1(ms)
	// MySlice does not satisfy []string (possibly missing ~ for []string in []string)
	// c := Clone4(ms)
	c := Clone5(ms)
	slices.Sort(c)
	return c.String()
}

func main() {
	fmt.Println(PrintSorted(MySlice([]string{"a", "b", "c"})))
}
