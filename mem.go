package cmem

/*
#include <stdlib.h>
*/
import "C"

import (
	"reflect"
	"runtime"
	"unsafe"
)

func alloc(size uintptr) *byte {
	return (*byte)(C.malloc(C.size_t(size)))
}

func AllocBytes(size int) []byte {
	buf := alloc(uintptr(size))
	out := (*[1 << 30]byte)(unsafe.Pointer(buf))[:size:size]
	runtime.SetFinalizer(&out, func(x *[]byte) {
		FreeBytes(*x)
	})
	return out
}

func FreeBytes(b []byte) {
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&b)))
	C.free(unsafe.Pointer(sliceHeader.Data))
}
