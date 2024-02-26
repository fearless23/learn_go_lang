package src

import (
	"fmt"
)

type consoleWriter struct{}

func (cw consoleWriter) write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type writer interface {
	write([]byte) (int, error)
}

// -----------------------------------------
type myIntType int

func (i *myIntType) increment(by myIntType) int {
	*i = *i + by
	return int(*i)
}

func (i *myIntType) decrement(by myIntType) int {
	*i = *i - by
	return int(*i)
}

type myIntMethods interface {
	increment(by myIntType) int
	decrement(by myIntType) int
}

// -------------------------------------------

type myArrayType []int

func (a myArrayType) sum() int {
	var total int
	for _, v := range a {
		total += v
	}
	return total
}

type myArrayMethods interface {
	sum() int
}

/*
- interface is like blueprint for set of methods/functions
- Same as typescript abstract type for class methods
- func defined in interface must be implemented in some concrete
*/
func Interfaces() {

	/*
		goInt := 1
		// goInt.increment method do not exist

		goSlice := []int{1,2,3}
		// goSlice.sum method do not exist

		// to add methods to GO int type
		// we first alias int to myIntType and then add methods on that
	*/

	myInt := myIntType(1)
	Log("myInt (1)", myInt)
	myInt.increment(10)
	Log("myInt +10", myInt)
	myInt.decrement(5)
	Log("myInt -5", myInt)

	var z myIntMethods = &myInt
	// if we declare z as type of myIntMethods; then must z must have 2 methods present
	// to do that; we set z to pointer to int type
	// since int (myIntType) has those 2 methods; z will also have those methods
	z.increment(1)
	Log("myInt +1", myInt)
	z.decrement(4)
	Log("myInt -4", myInt)

	myArray := myArrayType([]int{1, 2, 3})
	sum := myArray.sum()
	Log("sum", sum)

	var k myArrayMethods = &myArray
	Log("sum", k.sum())

	myStruct := consoleWriter{}
	writeResponse, err := myStruct.write([]byte{1, 2, 3})
	if err != nil {
		panic("something happen wrong")
	}
	Log("writeResponse", writeResponse)

	var writeC writer = &myStruct
	writeResponse2, err := writeC.write([]byte{30, 40, 50})
	if err != nil {
		panic("something happen wrong")
	}
	Log("writeResponse2", writeResponse2)

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

// func with receiver
// receiver is user struct instance
func (u user) ageInMonths() float64{
	return u.age*12
}

Whereas interface is set of common functions
available on any variable. Example below

type length interface {
	Length() int
}

Above interface has one prop or func called Length;
this interface provide blueprint of Length i.e always return int.

type MyString string
type MySlice []int
type user struct{
	name string
	age string
}


func (s MyString) Length() int{
	return len(s)
}

func (s MySlice) Length() int {
	return len(s)
}

func (u user) Length() int {
	return len(u.name)
}


So, now interface with length method is used on three types MyString, MySlice and user
but different implementation based on type as shown in 3 func above

Earlier, we have to define a struct to create new methods.
But with interface we can extend existing type without need to create new structs.

Any pointer to MyString, MySlice or user is of valid type length interface.

Readonly interface -- normal name
Write something -- add er at end like Writer


DS=Data structure (struct,myInt, myArray, map etc...)
DS1 methods = A, B, C, D
DS2 methods = A, X, Y, Z
DS3 methods = A, B, Y, Z

interface1 = A
interface2 = B, C, D
interface3 = X, Y, Z

DS1 has interface1 + interface2
Thus, DS1 can be used where interface1 or interface2 is needed
DS2 has interface 1 and interface3
Thus, DS2 can be used where interface1 or interface3 is needed


Without interface, type of function input will be DS1 | DS2 | or any custom DS
With interface, type of  function input can be interface1.
So, any DS which has methods mentioned in interface1 is valid...,
since function only going to use methods mentioned in interface


example
sum(input interface1)
sum function needs an input which has method A available;
instead of saying input type can be DS1|DS2|DS3 and then updating this will addition of new types
we can simply declare input to be of type interface1 and then any DS which has method A available is valid


- 1. Think of interface as set of methods whose input|output is defined
- 2. Your code can have so many data-structures according to business logic; but interfaces can group them by use-case
- 3. if data-structures are compared to users; then interfaces are user-groups or user-roles.
- 4. EXAMPLE: DS1, DS20 etc.. are part of interface1 and interface4
- 5. when designing function or business logic (specially dependency injection or generics); func can expect the input
			to have interface blueprint
- 6. We want input to func to have .len method present; so all data-structures which have .len method are valid; so
			if we create an interface called interface1 with len method; those data-structures are valid interface1

*/
