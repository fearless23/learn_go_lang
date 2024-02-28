package basics

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

func EnumIota() {
	var c int
	fmt.Printf("ea %v, %T\n", ea, ea)
	fmt.Printf("eb %v, %T\n", eb, eb)
	fmt.Printf("ea==eb %v\n", ea == eb)
	fmt.Printf("c==ea%v\n", c == ea)

	fmt.Printf("ex %v, %T\n", ex, ex)
	fmt.Printf("ey %v, %T\n", ey, ey)

	fmt.Printf("KB %v\n", KB)
	fmt.Printf("MB is %d KB\n", MB/KB)
	fmt.Printf("GB is %d MB\n", GB/MB)
	fmt.Printf("TB is %d GB\n", TB/GB)
	const fileSize float32 = 4 * 1024 * 1024 * 1024
	fmt.Printf("%.2f GB\n", fileSize/GB)

}
