package safer

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestKindOf(t *testing.T) {
	type testCase struct {
		exp reflect.Kind
		val interface{}
	}
	var tests = []testCase{
		{reflect.Invalid, nil},
		{reflect.Bool, true},
		{reflect.Int, int(1)},
		{reflect.Int8, int8(12)},
		{reflect.Int16, int16(123)},
		{reflect.Int32, int32(1234)},
		{reflect.Int64, int64(12345)},
		{reflect.Uint, uint(1)},
		{reflect.Uint8, uint8(12)},
		{reflect.Uint16, uint16(123)},
		{reflect.Uint32, uint32(1234)},
		{reflect.Uint64, uint64(12345)},
		{reflect.Uintptr, uintptr(0)},
		{reflect.Float32, float32(1.23)},
		{reflect.Float64, float64(1.23)},
		{reflect.Complex64, complex64(123)},
		{reflect.Complex128, complex128(123)},
		{reflect.Array, [3]int{1, 2, 3}},
		{reflect.Chan, make(chan string)},
		{reflect.Func, func() {}},
		// interfaces can't store a interface.
		{reflect.Invalid, (interface{})(nil)},
		{reflect.Map, make(map[int]string)},
		{reflect.Ptr, new(int)},
		{reflect.Ptr, (*interface{})(nil)},
		{reflect.Ptr, (*struct{})(nil)},
		{reflect.Ptr, (*func())(nil)},
		{reflect.Slice, []int{1, 2, 3}},
		{reflect.String, `strkind`},
		{reflect.Struct, struct{}{}},
		{reflect.UnsafePointer, unsafe.Pointer(new(int))},
	}
	for idx, test := range tests {
		t.Logf("test #%v exp %v kind from value of %#v", idx, test.exp, test.val)
		if kind := KindOf(test.val); kind != test.exp {
			t.Errorf(`exp kind %v; got %v`, test.exp, kind)
		}
	}
}

func TestPCForFunc(t *testing.T) {
	testFn := func() func() error {
		return func() error {
			return nil
		}
	}
	type testCase struct {
		fn interface{}
		ok bool
	}
	tests := []testCase{
		{testFn, true},
		{testFn(), true},
		{func() {}, true},
		{KindOf, true},
		{nil, false},
		{new(int), false},
		{0, false},
	}
	for _, test := range tests {
		got := PCForFunc(test.fn)
		if !test.ok {
			if got == 0 {
				continue
			}
			t.Fatalf(`exp zero value pc for invalid fn; got %v`, got)
		}
		if exp := reflect.ValueOf(test.fn).Pointer(); exp != got {
			t.Fatalf(`exp %v; got %v`, exp, got)
		}
	}
}
