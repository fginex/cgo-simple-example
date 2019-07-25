package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "api.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {

	buf := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(buf))

	sz := C.helloFromC((*C.char)(buf))

	b := C.GoBytes(buf, sz)
	msg := string(b)

	fmt.Printf("Len: %d\n", sz)
	fmt.Printf("Msg: %s\n", msg)
}
