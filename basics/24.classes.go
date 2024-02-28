// ### ---------- 3 --------------###
package basics

type User struct {
	Name       string
	City       string
	Attributes map[string]string
}

// >> way to create new struct; same as constructor(params?)
// createNewUser, create new struct instance with any initial values set as in javascript constructor
func createNewUser() *User {
	return &User{
		Attributes: map[string]string{"foo": "bar"},
	}
}

func (o *User) getSomeAttribute(prop string) (string, bool) {
	value, found := o.Attributes[prop]
	return value, found
}

// https://github.com/miguelmota/golang-for-nodejs-developers?tab=readme-ov-file#go-17
// basically, it is structs part2 (part1=12.structs.go)
func Classes() {
	// Typical way to create struct instance
	// user := User{
	// 	Name:       "",
	// 	City:       "",
	// 	Attributes: map[string]string{"foo": "bar"},
	// }
	// call createNewUser func
	user := createNewUser()
	user.City = "Delhi"

	city := user.City
	Log("city", city)

	foot, found2 := user.getSomeAttribute("foot")
	Log("found", found2)
	Log("foot is empty string", foot == "")
}
