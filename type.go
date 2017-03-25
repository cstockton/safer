package safer

import (
	"reflect"
	"unsafe"
)

const (

	// Size of ptr
	ptrSize = unsafe.Sizeof(uintptr(0))

	// From: runtime/typekind.go
	kindBool = 1 + iota
	kindInt
	kindInt8
	kindInt16
	kindInt32
	kindInt64
	kindUint
	kindUint8
	kindUint16
	kindUint32
	kindUint64
	kindUintptr
	kindFloat32
	kindFloat64
	kindComplex64
	kindComplex128
	kindArray
	kindChan
	kindFunc
	kindInterface
	kindMap
	kindPtr
	kindSlice
	kindString
	kindStruct
	kindUnsafePointer

	kindDirectIface = 1 << 5
	kindGCProg      = 1 << 6
	kindNoPointers  = 1 << 7
	kindMask        = (1 << 5) - 1
)

func reflectKindOf(k uint8) reflect.Kind {
	return reflect.Kind(k & kindMask)
}

// typeOf returns the runtimes representation of a Go type.
func typeOf(v interface{}) *_type {
	return *(**_type)(unsafe.Pointer(&v))
}

// pcForFunc returns a uintptr to the given funcs pc
func pcForFunc(fn interface{}) uintptr {
	if KindOf(fn) != reflect.Func {
		return 0
	}
	return **(**uintptr)(add(unsafe.Pointer(&fn), ptrSize))
}

// From: runtime/stubs.go
// --
// Should be a built-in for unsafe.Pointer?
//go:nosplit
func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

// From: runtime/type.go
// --
// tflag is documented in reflect/type.go.
//
// tflag values must be kept in sync with copies in:
//	cmd/compile/internal/gc/reflect.go
//	cmd/link/internal/ld/decodesym.go
//	reflect/type.go
type tflag uint8

const (
	tflagUncommon  tflag = 1 << 0
	tflagExtraStar tflag = 1 << 1
	tflagNamed     tflag = 1 << 2
)

// From: runtime/alg.go
// --
// typeAlg is also copied/used in reflect/type.go.
// keep them in sync.
type typeAlg struct {
	// function for hashing objects of this type
	// (ptr to object, seed) -> hash
	hash func(unsafe.Pointer, uintptr) uintptr
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
}

// Type is a Go type.
//
// From: runtime/type.go
// --
// Needs to be in sync with ../cmd/compile/internal/ld/decodesym.go:/^func.commonsize,
// ../cmd/compile/internal/gc/reflect.go:/^func.dcommontype and
// ../reflect/type.go:/^type.rtype.
type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      uint8
	align      uint8
	fieldalign uint8
	kind       uint8
	alg        *typeAlg
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       int32 // nameOff
	ptrToThis int32
}

type nameOff int32
type typeOff int32
type textOff int32

// From: runtime/typekind.go
// --
// isDirectIface reports whether t is stored directly in an interface value.
// func isDirectIface(k uint8) bool {
// 	return k&kindDirectIface != 0
// }
