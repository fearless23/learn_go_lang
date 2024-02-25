package src

import (
	"encoding/json"
	"reflect"
)

// User Type
type user struct {
	name   string
	Age    int
	height float32
	data   map[string](string)
}

type animal struct {
	name   string
	origin string
}

type bird struct {
	// Another structure embedded(similar to composing), not extenstion
	animal
	speed  float32
	canfly bool
}

type loginType struct {
	username string `max:"3"`
	password string
}

type userWithPascalCase struct {
	FullName string
	Age      float32
	City     string
}

type userWithJsonTags struct {
	FullName string  `json:"full--name"`
	Age      float32 `json:"age"`
	City     string  `json:"city"`
}

func Structs() {
	var user1 = user{
		name:   "Jaspreet",
		height: 202,
		Age:    30,
	}
	// All fields are optional by default in go struct and
	// are initialized by zero-value of each type like empty string for string; 0 for int etc...
	divider()
	logWithType("user1", user1)
	log("user1.age", user1.Age)
	log("user1.data", user1.data)

	var user2 user // or var user2 = user{} // empty initialization of all fields
	user2.Age = 34

	divider()
	log("user2", user2)
	logWithType("user2.name", user2.name)
	logWithType("user2.height", user2.height)

	var bird1 = bird{
		speed:  40,
		canfly: true,
	}

	divider()
	logWithType("bird1", bird1)
	bird1.animal = animal{name: "Tiger", origin: "India"}
	logWithType("bird1", bird1)
	bird1.name = "Crow2"
	logWithType("bird1", bird1)
	bird1.origin = "India"
	logWithType("bird1", bird1)

	divider()
	// Struct Tags - loginType
	var t = reflect.TypeOf(loginType{})
	var field, exist = t.FieldByName("username")
	log("field-exist", exist)
	log("tag", field.Tag)

	divider()
	// Struct Tags - loginType
	var l = loginType{username: "morethanthree", password: "Df"}
	log("username", l.username)
	log("password", l.password)

	divider()

	out, err := json.MarshalIndent(bird1, "", "  ")
	if err != nil {
		panic("cannot marshal bird1 to json")
	}
	log("bird1 json", string(out))
	// If field of struct are camelCase; they are not exported to outside world
	// since, json package is external to core go language, for json to work properly we need to
	// expose struct fields and thus we have to use PascalCase (this is weird decision from go)

	var jassi = userWithPascalCase{
		FullName: "Jassi",
		Age:      31,
		City:     "Barnala",
	}
	out2, err := json.MarshalIndent(jassi, "", "  ")
	if err != nil {
		panic("cannot marshal jassi1 to json")
	}
	log("jassi json", string(out2))
	// since struct userWithPascalCase field are PascalCase, all are exposed to outside world
	// and thus json was able to parse them

	var jassi2 = userWithJsonTags{
		FullName: "Jassi",
		Age:      31,
		City:     "Barnala",
	}
	out3, err := json.MarshalIndent(jassi2, "", "  ")
	if err != nil {
		panic("cannot marshal jassi2 to json")
	}
	log("jassi2 json", string(out3))
	// since struct userWithJsonTags is same as userWithPascalCase, all fields are parse-able
	// only difference is that json:xxx tags output the fields to be camelCase or any type of our choice

	divider()
	print("--> copying jassi struct instance")
	var jassiCopy = jassi
	jassi.FullName = "Happy"
	print("--> changing original jassi struct instance .fullName property to Happy")
	log("jassi", jassi)
	log("jassiCopy", jassiCopy)

	// jassiCopy and jassi are two different instances.
	// Thus, copying an struct is copy-by-value
}
