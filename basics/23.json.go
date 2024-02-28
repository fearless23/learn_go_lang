package basics

import (
	"encoding/json"
)

// ### ---------- 1 --------------###
func JsonArray() {
	// 1. start with a literal string to capture valid json
	jsonString := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	Log("jsonString", jsonString)

	// 2. unmarshall into string[]
	var array []string
	err := json.Unmarshal([]byte(jsonString), &array)
	if err != nil {
		panic("error unmarshalling json")
	}
	Log("array", array)

	jsonString2, _ := json.Marshal(array)
	Log("jsonString", string(jsonString2))
}

// ### ---------- 2 --------------###
type JsonUserStruct struct {
	Name    string   `json:"name"`
	Age     uint8    `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func JsonObject() {
	// 1. capture valid json representing a user with name and age
	jsonUserString := `{ "name": "jassi", "age": 34, "hobbies": ["a","b","c"] }`
	Log("jsonUserString - start", jsonUserString)

	// 2. unmarshall into JsonUserStruct
	var user JsonUserStruct
	// json.Unmarshall is like JSON.parse from javascript
	// json string can be converted to a struct({}) or array([]), since json top level can be
	// an array or object
	err := json.Unmarshal([]byte(jsonUserString), &user)
	if err != nil {
		Log("error unmarshalling user json", err.Error())
		panic("error unmarshalling user json")
	}
	logWithType("user", user)

	// Now, lets try to do the opposite
	// json.Marshal is like JSON.stringify from javascript
	// struct or array can be converted to a JSON string
	jsonUserString2, _ := json.Marshal(user)
	Log("jsonUserString - end", string(jsonUserString2))
}
