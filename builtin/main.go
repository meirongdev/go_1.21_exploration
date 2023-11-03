package main

import (
	"fmt"
	"math"
)

type tToClear struct {
	a int
}
type tToClear2 []int
type tToClear3 map[string]int

func main() {
	var x, y, z int = 1, 2, 3
	fmt.Printf("max(x) :%v\n", max(x))
	fmt.Printf("max(x, y): %v\n", max(x, y))
	fmt.Printf("max(x, y, z): %v\n", max(x, y, z))
	fmt.Printf("max(1, NaN): %v\n", max(1, math.NaN()))
	// 输出
	// max(x) :1
	// max(x, y): 2
	// max(x, y, z): 3

	fmt.Printf("min(x): %v\n", min(x))
	fmt.Printf("min(x, y): %v\n", min(x, y))
	fmt.Printf("min(x, y, z): %v\n", min(x, y, z))
	fmt.Printf("min(1, NaN): %v\n", min(1, math.NaN()))
	// 输出
	// min(x): 1
	// min(x, y): 1
	// min(x, y, z): 1
	// max(1, NaN): NaN

	fmt.Printf("max(NaN, NaN): %v\n", max(math.NaN(), math.NaN()))
	// 输出
	// max(NaN, NaN): NaN

	var a, b, c string = "a", "b", "c"
	fmt.Printf("max(a): %v\n", max(a))
	fmt.Printf("max(a, b): %v\n", max(a, b))
	fmt.Printf("max(a, b, c): %v\n", max(a, b, c))
	// 输出
	// max(a): a
	// max(a, b): b
	// max(a, b, c): c

	fmt.Printf("min(a): %v\n", min(a))
	fmt.Printf("min(a, b): %v\n", min(a, b))
	fmt.Printf("min(a, b, c): %v\n", min(a, b, c))
	// 输出
	// min(a): a
	// min(a, b): a
	// min(a, b, c): a

	// var x64 int64 = 1
	// fmt.Println(max(x, x64)) // invalid argument: mismatched types int (previous argument) and int64 (type of x64)
	// fmt.Println(max(x, 1.1)) // 1.1 (untyped float constant) truncated to intcompilerTruncatedFloat

	// -------------clear----------------
	var slice1 = []int{1, 2, 3}
	fmt.Printf("before clear: slice1=%v, len(slice1)=%d, cap(slice1)=%d\n", slice1, len(slice1), cap(slice1))
	clear(slice1)
	fmt.Printf("after clear: slice1=%v, len(slice1)=%d, cap(slice1)=%d\n", slice1, len(slice1), cap(slice1))
	// 输出
	// before clear: slice1=[1 2 3], len(slice1)=3, cap(slice1)=3
	// after clear: slice1=[0 0 0], len(slice1)=0, cap(slice1)=3

	var slice2 = []tToClear{{1}, {2}, {3}}
	fmt.Printf("before clear: slice2=%v, len(slice2)=%d, cap(slice2)=%d\n", slice2, len(slice2), cap(slice2))
	clear(slice2)
	fmt.Printf("after clear: slice2=%v, len(slice2)=%d, cap(slice2)=%d\n", slice2, len(slice2), cap(slice2))
	// 输出
	// before clear: slice2=[{1} {2} {3}], len(slice2)=3, cap(slice2)=3
	// after clear: slice2=[{0} {0} {0}], len(slice2)=3, cap(slice2)=3
	var slice3 = tToClear2{1, 2, 3}
	fmt.Printf("before clear: slice3=%v, len(slice3)=%d, cap(slice3)=%d\n", slice3, len(slice3), cap(slice3))
	clear(slice3)
	fmt.Printf("after clear: slice3=%v, len(slice3)=%d, cap(slice3)=%d\n", slice3, len(slice3), cap(slice3))
	// 输出
	// before clear: slice3=[1 2 3], len(slice3)=3, cap(slice3)=3
	// after clear: slice3=[0 0 0], len(slice3)=0, cap(slice3)=3

	var map1 = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("before clear: map1=%v, len(map1)=%d\n", map1, len(map1))
	clear(map1)
	fmt.Printf("after clear: map1=%v, len(map1)=%d\n", map1, len(map1))
	// 输出
	// before clear: map1=map[a:1 b:2 c:3], len(map1)=3
	// after clear: map1=map[], len(map1)=0
	var map2 = map[string]tToClear{"a": {1}, "b": {2}, "c": {3}}
	fmt.Printf("before clear: map2=%v, len(map2)=%d\n", map2, len(map2))
	clear(map2)
	fmt.Printf("after clear: map2=%v, len(map2)=%d\n", map2, len(map2))
	// before clear: map2=map[a:{1} b:{2} c:{3}], len(map2)=3
	// after clear: map2=map[], len(map2)=0

	var map3 = tToClear3{"a": 1, "b": 2, "c": 3}
	fmt.Printf("before clear: map3=%v, len(map3)=%d\n", map3, len(map3))
	clear(map3)
	fmt.Printf("after clear: map3=%v, len(map3)=%d\n", map3, len(map3))
	// before clear: map3=map[a:1 b:2 c:3], len(map3)=3
	// after clear: map3=map[], len(map3)=0
}
