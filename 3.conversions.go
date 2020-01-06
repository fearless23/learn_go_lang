package main

import (
	"fmt"
	"strconv"
)

var strToInt = strconv.Atoi

func conversion() {

	var i int = 42

	// j := string(i)
	j := strconv.Itoa(i)
	// string to number fn atoi stored in ff
	fmt.Printf("%T", strToInt)
	g, e := strToInt("45")

	fmt.Println(i, j, g, e)
}
