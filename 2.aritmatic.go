package main

import (
	"fmt"
)

func aritmatics() {

	var i float32 = 42
	var j float32 = 4
	var k float32 = i / j

	var z complex64 = 1 + 2i
	var z2 = complex(4.3, 6.7)

	fmt.Println(k, z, real(z), imag(z), z2)
}
