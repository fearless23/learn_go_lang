package main

import "fmt"

func numbers() {
	var i = 42.
	fmt.Printf("%v, %T\n", i, i)
	i = 43.9
	fmt.Printf("%v, %T\n", i, i)

	var j float32 = 42
	fmt.Printf("%v, %T\n", j, j)
	j = 43.9
	fmt.Printf("%v, %T\n", j, j)

	var k float32
	fmt.Printf("%v, %T\n", k, k)
	k = 34
	fmt.Printf("%v, %T\n", k, k)

	l := 56.
	l = 34.6
	fmt.Printf("%v, %T", l, l)

	m := 2E6
	fmt.Printf("m: %v, %T", m, m)
}
