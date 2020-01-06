package main

import ("reflect")

// lowercase u - make it unexported..

// User Type
type user struct {
	name    string
	Age     int
	height  float32
	newProp string
}

type animal struct {
	name   string
	origin string
}

type bird struct {
	animal   // Another structure embedded, not = extenstion
	speedKPH float32
	canfly   bool
}

type loginType struct {
	username string `max:"100"`
	password string
}

func structs() {
	var user1 = user{
		name:   "Jaspreet",
		Age:    30,
		height: 202,
	}

	golog("user1", user1)
	gologv("user1.age", user1.Age)

	var user2 user
	user2.Age = 34

	gologv("user2", user2)
	golog("user2.height", user2.name)

	var bird1 = bird{
		speedKPH: 40,
		canfly:   true,
	}

	golog("bird1", bird1)
	bird1.animal = animal{name: "Crow", origin: "Austrailia"}
	golog("bird1", bird1)
	bird1.name = "Crow2"
	golog("bird1", bird1)
	bird1.origin = "India"
	golog("bird1", bird1)

	// Struct Tags - loginType
	var t = reflect.TypeOf(loginType{})
	var field, exist = t.FieldByName("username")
	gologv("field-exist", exist)
	gologv("tag", field.Tag)

	// var y = x // y is new copy of x; where
	// x is a struct...
}
