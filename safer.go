// Package safer provides safer access to unsafe operations by providing simple
// functions with high test coverage that will never panic, instead returning
// zero values.
package safer

import (
	"reflect"
)

// KindOf like reflect.TypeOf(v).Kind() returns the reflect.Kind of the given v
// by using the unsafe package to directly access the underlying T. This is
// about 30x faster than reflect because no allocations are needed. If
// reflect.Value may live on the stack in the future most of the performance
// benefit will be lost.
func KindOf(v interface{}) reflect.Kind {
	if typ := typeOf(v); typ != nil {
		return reflectKindOf(typ.kind)
	}
	return 0
}

// PCForFunc compliments runtime.FuncForPC by returning the program counter for
// a given function value. Returns the zero value uintptr on failure.
func PCForFunc(fn interface{}) uintptr {
	return pcForFunc(fn)
}
