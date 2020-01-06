package main

import (
	"fmt"
)

// Enumarated constants
const (
	ea = iota
	eb = iota
)

// Ignore First value by assigning to blank identifier
const (
	_ = iota
	ex
	ey
)

// Exported Constant = First Character Capital

// Bytes in kilo, giga, tera
const (
	_  = iota
	KB = 1 << (10 * iota) // Get multiplied by 1024
	MB
	GB
	TB
)

func enumIota() {
	var c int
	fmt.Printf("%v, %T\n", ea, ea)
	fmt.Printf("%v, %T\n", eb, eb)
	fmt.Printf("%v\n", ea == eb)
	fmt.Printf("%v\n", c == ea)

	fmt.Printf("%v, %T\n", ex, ex)
	fmt.Printf("%v, %T\n", ey, ey)

	const fileSize float32 = 4 * 1024 * 1024 * 1024
	fmt.Printf("%.2f GB\n", fileSize/GB)

}
