package safer_test

import (
	"fmt"
	"runtime"

	safer "github.com/cstockton/safer"
)

func ExampleKindOf() {
	val := []int{1, 2, 3}
	kind := safer.KindOf(val)
	fmt.Printf("KindOf(%#v) -> %v\n", val, kind)

	// Output:
	// KindOf([]int{1, 2, 3}) -> slice
}

func ExamplePCForFunc() {
	// Obtain the PC from a function, this could be any func value.
	pc := safer.PCForFunc(func() {})

	// Print the functions name using runtime.FuncForPC
	fmt.Println(runtime.FuncForPC(pc).Name())

	// Output:
	// github.com/cstockton/safer_test.ExamplePCForFunc.func1
}
