package main

import (
	"fmt"
)

// func log(int j){
// 	fmt.Printf("%v,%T\n", j, j)
// }

func stringsAndRunes() {

	var i string = "a"
	j := []byte(i)

	var r = 'a'

	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v,%T\n", j, j)
	fmt.Printf("%v,%T\n", r, r)

}
