package main

import (
	"context"
	"fmt"
	"reflect"
)

// This is how some assertions libraries check for equality
// This particular implementation was taken from:
// https://github.com/golang/mock/blob/0cdccf5f55d777b12c1ac5a93f607cdd1dbf5296/gomock/matchers.go#L104
func equal(x, y any) bool {
	// In case, some value is nil
	if x == nil || y == nil {
		return reflect.DeepEqual(x, y)
	}

	xVal := reflect.ValueOf(x)
	yVal := reflect.ValueOf(y)

	if xVal.Type().AssignableTo(yVal.Type()) {
		x1ValConverted := xVal.Convert(yVal.Type())
		return reflect.DeepEqual(x1ValConverted.Interface(), yVal.Interface())
	}

	return false
}

func main() {
	background := context.Background()
	todo := context.TODO()

	// This probably should have never worked in <1.21
	// Mostly, this occurs when tests were carelessly written by declaring the
	// mock to expect context.TODO() but the call uses context.Background()
	// This is a trivial fix to make both use the same call either TODO or Background
	if !equal(background, todo) {
		fmt.Printf("background: %T(%#v) is equel to TODO: %T(%#v)", background, background, todo, todo)
	}
	// 1.21是不相等的
	// 1.20是相等的
}
