package cmem

//include <stdlib.h>
import "C"
import "unsafe"

func alloc(size uintptr) *byte {
	return (*byte)(C.malloc(C.size_t(size)))
}

func free(ptr *byte, size uintptr) { //size not used
	C.free(unsafe.Pointer(ptr))
}

func AllocBytes(size int) []byte {
	alloc := alloc(uintprt(size))
	buf := (*[1 << 30]byte)(unsafe.Pointer(alloc))[:count]
	return buf
}

func FreeBytes(ptr *byte) {
	free(ptr, uintptr(0))
}

/*
	alloc := memory.Alloc(uintptr(count))
	buf := (*[1 << 30]byte)(unsafe.Pointer(alloc))[:count]

	//free
	memory.Free(alloc, uintptr(count))
*/
