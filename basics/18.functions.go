package basics

import (
	"fmt"
)

func Functions() {
	a := sayMessage("Hello")
	Log("a", a)

	divider()
	s1 := sum(1, 2, 3, 4, 5, 6)
	Log("Sum s1", s1)

	divider()
	s2 := sum2(1, 2, 3, 4, 5, 6)
	Log("Sum s2", s2)

	divider()
	d, err := divide(4.0, 0.0)
	Log("err", err)
	Log("d", d)

	divider()
	k := getNums()
	Log("k", k)

	// Anonymous Function
	fn := func() {
		Log("$", "anonymous fn")
	}
	fn()

	// Anonymous Function - IIFE
	func() {
		Log("Hello GO", "")
	}()

	// Declare anonymous fn
	// afn not available in main, before this
	//   var =   let in JS
	// const = const in JS

	// var afn func(i int, j int) int
	afn := func(i, j int) int {
		return i + j
	}
	afn(2, 3)

	// -----------------------

	divider()
	user1 := myUser{
		firstName: "Jaspreet",
		lastName:  "Singh",
		fullName:  "original-full-name",
		age:       35,
	}

	Log("user1 (created)", user1)

	fullName := user1.getFullName()
	Log("user1:fullName", fullName)
	Log("user1 (user1 passed to function and modified)", user1)
	// getFullName function changes fullName on user1 passed to function
	// but it do not modify user1 instance

	user1.setFullName()
	Log("user1 (user1 reference passed and modified)", user1)
	// setFullName function changes fullName on user1 reference passed to function
	// and it does modify user1 instance

	divider()
	// let do same thing with map; as we did with user1
	myMap := map[string](int){
		"a": 1,
		"b": 2,
	}
	Log("myMap (created)", myMap)

	returnedMap := intakeSomeMap(myMap)
	Log("myMap (after modification in func)", myMap)
	Log("returnedMap", returnedMap)

	// passing map to function which internally changes it, behaves as passed-by-reference

	// So, conclusion is
	// any type passed to function which is internally modified in function,
	// behaves similar to their copy traits
	// maps are copied-by-reference and  same behaviour is seen when map passed to func
	// strcut are copeid-by-value and same behaviour is seen when passed to func
	// Note that; if we pass pointer (&value); internal modification always modify passed value obviously

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
		return 0.0, fmt.Errorf("cannot divide by zero")
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
	// changes in u here dont change u outside of this function
	// To change values outside, pass pointer...
	// see setFullName function below
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
func getNums() *[5]int {
	var roles = [5]int{1, 2, 3, 4, 5}
	return &roles
}

func intakeSomeMap(m map[string](int)) map[string](int) {
	m["c"] = 3
	return m
}

func NonMain() {

	i2 := [3][3]int{{1, 2, 3}, {10, 20, 30}, {7, 8, 9}}
	j2 := i2 // this is a deep-copy
	i2[0][0] = 1000
	Log("j2", j2)
	Log("i2", i2)
}

// Function with receivers
/*
normal function

func add(a int, b int) int {
	return a + b
}


---> function with receiver (defining class methods with value of receiver)
func (receiver int) sum(b int) int {
	return receiver + int
}

---> function with recevier as pointer (defining class methods with pointer to receiver)
// similar to this in javascript
func (receiver *int) sum(b int) int {
	return receiver + int
}
*/
