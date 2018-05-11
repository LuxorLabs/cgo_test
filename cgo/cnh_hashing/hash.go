package cnh_hashing

// #cgo CFLAGS: -std=c11 -D_GNU_SOURCE
// #cgo LDFLAGS: -L${SRCDIR} -lcnh_hashing -Wl,-rpath,${SRCDIR} -lstdc++
// #include <stdlib.h>
// #include <stdint.h>
// #include "src/hash.h"
import "C"
import "unsafe"

func CNHHash(blob []byte) []byte {
	output := make([]byte, 32)
	C.cryptonight_hash((*C.char)(unsafe.Pointer(&blob[0])), (*C.char)(unsafe.Pointer(&output[0])), (C.uint32_t)(len(blob)))
	return output
}
