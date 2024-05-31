package main

import (
	"fmt"
	"testingC/go_files"
)

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -Lgo_files/ -lhello
#include "go_files/hello.h"
*/
import "C"

func main() {
	a := 3
	b := 5
	sum := hello.Add(a, b)
	//sum := C.add(C.int(a), C.int(b))
	fmt.Printf("Sum: %d\n", int(sum))
	//typedb_driver.Connection_open_core("127.0.0.1")

}

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -Lgo_files/ -lhello
#include "go_files/hello.h"
*/
