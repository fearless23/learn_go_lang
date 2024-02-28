package basics

import (
	"fmt"
	"strconv"
)

var strToInt = strconv.Atoi
var intToStr = strconv.Itoa

func Conversions() {

	// way1
	// var x int
	// x = 42
	// way2.1
	// var y1 int = 32
	// way2.2
	// var y2 = 32
	// way3
	// z := 32

	//
	// start with var y2 = 32
	// add type if you feel it is complex or necessary (other than inferred)
	// if no starting value; use var x int
	// in for-loop; function output use z := 32

	var x int = 12e1
	y := intToStr(x)
	z, e := strToInt("1")
	fmt.Printf("x is %v (%T)\n", x, x)
	fmt.Printf("y is %v (%T)\n", y, y)
	fmt.Printf("z is %v (%T)\n", z, z)

	if e != nil {
		var error = e.Error()
		fmt.Printf("error type is %T\n", error)
		fmt.Printf("error is %v\n", error)
	}
}
