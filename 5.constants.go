package main

import (
	"fmt"
)

func myConstnts() {
	const name string = "fearless"
	// Constants can only be made up of
	// hard code or combination of constants
	const a = "My name is " + name

	// Var cant be used with constant.
	var name2 = "Jaspreet"
	// const b = "My name is " + name2

	var c = "My name is " + name2
	fmt.Printf("%v, %T\n", c, c)

	const i = 42
	const j = 5.7
	const k = i + j
	// constant of diff number types
	// can be added.Because they are created
	// once.
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)

	var x = 42
	var y = 5.7
	// var z = i + j ERROR
	// variables of different number types
	// can`t be added
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	// fmt.Printf("%v, %T\n", k, k)
}
