package main

import (
	"fmt"
)

func bitwiseOps() {

	var a = 10 // Binary 2^3 + 2^1 = 1010
	var b = 3  // Binary 2^1 + 2^0 = 11 or 0011

	fmt.Printf("a: %v %T\n", a, a)
	fmt.Printf("b: %v %T\n", b, b)
	// Bitwise operators
	// Compare each byte
	var x = a & b  // 1010 AND 0011 = 0010 => 2
	var y = a | b  // 1010 OR  0011 = 1011 => 11
	var z = a ^ b  // 1010 XOR 0011 = 1001 => 9
	var s = a &^ b // 1010 AND NOT 0011 = 1010 AND 1100 = 1000 => 8
	var t = a << 3 // 1010 shift 3 left = 1010000 = 80
	var u = a >> 3 // 1010 shift 3 right = 1 = 1

	fmt.Printf("x: %v %T\n", x, x)
	fmt.Printf("y: %v %T\n", y, y)
	fmt.Printf("z: %v %T\n", z, z)
	fmt.Printf("s: %v %T\n", s, s)
	fmt.Printf("t: %v %T\n", t, t)
	fmt.Printf("u: %v %T\n", u, u)
}
