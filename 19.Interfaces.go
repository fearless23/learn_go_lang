package main

import (
	"fmt"
)

type writer interface {
	write([]byte) (int, error)
}

type consoleWriter struct{}

func (cw consoleWriter) write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

/*Incrementer Increase value by 1*/
type Incrementer interface {
	Increment() int
}

type myIntType int

func (i *myIntType) Increment() int {
	*i++
	return int(*i)
}

type myArrayType []int

func (a myArrayType) Sum() int {
	// var str string = *s
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

// Main Function
func mainInterfaces() {
	/*
		goInt := 1
		// No, Increment method	on GO int type

		goSlice := []int{1,2,3}
		// No, Sum method on GO []int type
	*/

	myInt := myIntType(1)
	gologl(myInt)
	myInt.Increment()
	gologl(myInt)

	myArray := myArrayType([]int{1, 2, 3})
	sum := myArray.Sum()
	gologl(sum)

	// u := [4]int{1,2,3,4}
	// var z Incrementer = &u
	// [4]int type dont implement Incrementor
	// Interfaces can be set on any Type...
	var z Incrementer = &myInt
	z.Increment()
	// Same as myInt.Increment(), increases value
	// of myInt by 1
	// Increment method is present on myInt pointer, as
	// increment method expects pointer of myIntType. So
	// myInt.Increment() and anyPointer of myInt of type
	// Incrementor has .Increment() method.

	// z can be type converted to another interface,
	// as long as those methods are present on new
	// interface.

	// Methods implemented receive data, pointer or value
	// on type it is implemented. Interfaces related to
	// those methods will work with same input type.
}

/*
Struct Methods

type user struct{
	name string
	age int
}

func (u user) ageInMonths() float64{
	return u.age*12
}

Whereas interface is set of common functions
available on any variable. Example below

type length interface {
	Length() int
}

Above interface has one prop or func called
Length; this interface provide blueprint of
Length i.e always return int.

type Stringer string
type Slicee []int
type user struct{
	name string
	age string
}


func (s Stringer) Length() int{
	return len(s)
}

func (s Slicee) Length() int {
	return len(s)
}

func (u user) Length() int {
	return 2
}


So, now interface length is used on
three types Stringer, Slicee, user
but depending upon type the methods
behaves differently.
The common thing is the input and output type

Earlier, we have to define a struct to create
new methods. But with interface we can
extend existing type without need to create
new structs.

Any pointer to Stringer, Slicee or user is
of valid type length interface.

Readonly interface -- normal name
Write something -- add er at end like Writer


*/
