package main

// solve this by either have the dylib next to the .go files and compile with just filename
// or make sure the relative path to the .go file is the fiel name its compiled with

// so in the cgo we can put either the absolute path to hello2
// or when generating we can generate from main directory, so that it is encoded in the dylib, and macos isnt confused.
// so we specify go_files/lib and go_files/.cxxfile so macos is happy, rather than copy into the main. ok cool

/*
#cgo LDFLAGS: -Lc_files -ltypedb_driver_go_native -framework CoreFoundation
*/
import "C"
import (
	"fmt"
	typedb_driver "testingC/c_files"
)

func main() {
	//a := 3
	//b := 5
	//sum := hello.Add(a, b)
	//fmt.Printf("Sum: %d\n", int(sum))

	typedb_driver.Connection_open_core("127.0.0.1")

	value := 667
	driver_value := typedb_driver.Value_new_long(int64(value))
	fmt.Println(typedb_driver.Value_get_long(driver_value))

	// get the .a file, use it as the rust cbind file, we have the header file as well.
	// then we compile it to link against the shared library, just need 1 dylib

	// the cxx and the header and the .a become a dylib
}

/*
#cgo LDFLAGS: -Lgo_files -lhello2
*/
