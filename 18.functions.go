package main

import (
	"fmt"
)

func functionsGO() {
	sayMessage("Hello")
	s := sum2(1, 2, 3, 4, 5, 6)
	gologv("Sum", s)

	d, err := divide(4.0, 0.0)
	gologv("err", err)
	gologv("d", d)

	// Anonymous Function
	fn := func() {
		gologv("$", "anonymous fn")
	}
	fn()

	// Anonymous Function - IIFE
	func() {
		gologv("Hello GO", "")
	}()

	// Declare anonymous fn
	// afn not available in main, before this
	//   var =   let in JS
	// const = const in JS
	var afn func(i int, j int) int
	afn = func(i, j int) int {
		return i + j
	}

	afn(2, 3)

	user1 := myUser{
		firstName: "Jaspreet",
		lastName:  "Singh",
	}

	gologv("user1", user1)

	fullName := user1.getFullName()
	gologv("user1:fullName", fullName)

	user1.setFullName()
	gologv("user1", user1)

}

func sayMessage(msg string) string {
	/*
		changing value of msg inside
		do not effect msg value in
		global or main or other
		functions.
		if msg is pointer, changing its
		value will change its value.
	*/
	return msg
}

/*Sum Function*/
func sum(values ...int) int {
	/*
		Variatic params or rest params as in JS
		Only Last value can be variatic
	*/
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

/*Named Return Function*/
func sum2(values ...int) (result int) {
	/*
		Named return = result and type int
		int is always initialized as 0
		we can change value of result
	*/
	for _, v := range values {
		result += v
	}
	return
}

// Return multiple types
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

type myUser struct {
	firstName string
	lastName  string
	fullName  string
	age       int
}

// Method on myUser struct - Immutable
func (u myUser) getFullName() string {
	// Copy is created here of the struct,
	// changes in u here dont change u outside
	// of this function
	// To change values outside, pass pointer...
	// see setFullNameG
	// use these methods for returning a value
	u.fullName = u.firstName + " " + u.lastName
	// above line will not change instance of u.
	return u.fullName
}

// Method on myUser struct - Modifies
func (u *myUser) setFullName() {
	// Use pointer to
	u.fullName = u.firstName + u.lastName
}

// Functions can receive and return pointers.
// Returning pointers of variables created inside func
// when a variable is created inside func, it is
// destroyed/deleted after func executes, but if we
// return a pointer, GO moves memory from heap to shared for
// that pointer.
func getNums(n int) *[5]int {
	var roles = [5]int{1,2,3,4,5}
	return &roles
}
