package cmem

/*
#include <stdlib.h>
*/
import "C"

import (
	"reflect"
	"unsafe"
)

func alloc(size uintptr) *byte {
	return (*byte)(C.malloc(C.size_t(size)))
}

func AllocBytes(size int) []byte {
	alloc := alloc(uintptr(size))
	buf := (*[1 << 30]byte)(unsafe.Pointer(alloc))[:size:size]
	return buf
}

func FreeBytes(b []byte) {
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&b)))
	C.free(unsafe.Pointer(sliceHeader.Data))
}
