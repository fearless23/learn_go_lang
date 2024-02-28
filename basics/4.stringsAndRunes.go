package basics

import (
	"fmt"
)

// func log(int j){
// 	fmt.Printf("%v,%T\n", j, j)
// }

func StringsAndRunes() {

	var i string = "ab"
	j := []byte(i) // or []uint8(i)

	var r = 'a'

	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v,%T\n", j, j)
	fmt.Printf("%v,%T\n", r, r)

}
