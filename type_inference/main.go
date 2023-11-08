package main

import (
	"fmt"
	"net/netip"
)

func IsEven[T int | int8 | int16](n T) bool {
	return n%2 == 0
}

func Any(numbers []int, fn func(int) bool) bool {
	for _, v := range numbers {
		if fn(v) {
			return true
		}
	}
	return false
}

func AnyGeneric[T any](numbers []T, fn func(T) bool) bool {
	for _, v := range numbers {
		if fn(v) {
			return true
		}
	}
	return false
}

type RandomElement[T any] interface {
	// 返回一个随机数
	RandomElement() (T, bool)
}

func MustRandom[T any](collection RandomElement[T]) T {
	val, ok := collection.RandomElement()
	if !ok {
		panic("collection is empty")
	}
	return val
}

// MyList 是自定义的泛型类型
type MyList[T any] []T

func (l MyList[T]) RandomElement() (T, bool) {
	if len(l) == 0 {
		var zero T
		return zero, false
	}
	// assume the 0 is a random number
	return l[0], true
}

// IPSet是个具体类型
type IPSet map[netip.Addr]bool

func (s IPSet) RandomElement() (netip.Addr, bool) {
	for k := range s {
		// assume the first element is a random element
		return k, true
	}
	return netip.Addr{}, false
}

func Add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("Partially instantiated generic functions")
	numbers := []int{1, 2, 3}

	// This works in Go 1.21
	// In Go 1.20, you will get:
	//   "cannot use generic function IsEven without instantiation"
	anyEven := Any(numbers, IsEven)

	// Prior to Go 1.21 you had to write:
	// anyEven := Any(numbers, IsEven[int])

	fmt.Printf("anyEven: %v\n", anyEven)

	// This works in Go 1.21
	// In Go 1.20, you will get:
	//   "cannot use generic function IsEven without instantiation"
	anyEven = AnyGeneric(numbers, IsEven)
	fmt.Printf("anyEven: %v\n", anyEven)

	// Prior to Go 1.21 you had to write:
	anyEven = AnyGeneric(numbers, IsEven[int])
	fmt.Printf("anyEven: %v\n", anyEven)

	// 泛型函数还可以赋值给函数变量，从函数变量的类型推断出泛型函数的类型
	// This works in Go 1.21
	// In Go 1.20, you will get:
	//   "cannot use generic function IsEven without instantiation"
	// var isEven func(int) bool = IsEven
	// fmt.Println(isEven(2))

	fmt.Println("Interface assignment inference")

	// 作为调用方MustRandom函数，可以从参数类型推断出泛型函数的类型
	// In Go 1.20, you will get:
	//   "type MyList[int] of MyList[int]{…} does not match RandomElement[T] (cannot infer T)"
	randomInt := MustRandom(MyList[int]{1, 2, 3})
	var ipSet = make(IPSet)
	ipSet[netip.Addr{}] = true
	// // In Go 1.20, you will get:
	// //   "type IPSet of ipSet does not match RandomElement[T] (cannot infer T)"
	randomIP := MustRandom(ipSet)

	// In Go 1.20 you would have to do:
	// randomInt := MustRandom[int](MyList[int]{1, 2, 3})
	// randomIP := MustRandom[netip.Addr](ipSet)
	fmt.Printf("randomInt: %v\n", randomInt)
	fmt.Printf("randomIP: %v\n", randomIP)

	fmt.Println("Type inference for untyped constants")
	var x = 1 + 2
	var y = 1 + 2.5
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	// In Go 1.20, you will get:
	//  default type float64 of 2.5 does not match inferred type int for T
	var z = Add(1, 2.5)
	fmt.Printf("%v %T\n", z, z)
}
