package hello

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lhello
#include "hello.h"
*/
import "C"

// Add is a wrapper around the C add function.
func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}
