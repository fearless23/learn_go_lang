package main

import (
	"fmt"
)

func golog(a string, b interface{}) {
	fmt.Printf(a+": %v, %T\n", b, b)
}

func gologl(b interface{}) {
	fmt.Println(b)
}

func gologv(a string, b interface{}) {
	fmt.Printf(a+": %v\n", b)
}

func gologs(a string, b interface{}) {
	fmt.Printf(a+": %s\n", b)
}
