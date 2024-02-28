package basics

import "fmt"

func Numbers() {
	var i = 42.
	fmt.Printf("i: %v, %T\n", i, i)
	i = 43.9
	fmt.Printf("i: %v, %T\n", i, i)

	var j float32 = 42
	fmt.Printf("j: %v, %T\n", j, j)
	j = 43.9
	fmt.Printf("j: %v, %T\n", j, j)

	var k float32
	fmt.Printf("k: %v, %T\n", k, k)
	k = 34
	fmt.Printf("k: %v, %T\n", k, k)
	fmt.Printf("k: %b, %T\n", k, k)

	var v int8 = 127
	fmt.Printf("v: %v, %T\n", v, v)
	l := 56.23
	// l = 34.6
	fmt.Printf("l: %v, %T\n", l, l)

	m := 2e6
	fmt.Printf("m: %v, %T\n", m, m)

	var t = 60 >> 7
	fmt.Printf("t: %b\n", 60)
	fmt.Printf("t: %v %b\n", t, t)

}
