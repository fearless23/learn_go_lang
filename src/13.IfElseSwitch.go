package src

func Ifelse() {

	if true {
		gologv("$", "inside if block")
	}

	// With initializer
	if true {
		gologv("$", "init")
	}
	if a := 3; a > 2 {
		gologv("a inside if with init", a)
	}
	if a, b := 2, 3; b < 4 {
		gologv("", "inside if block with initializer")
		golog("a", a)
		golog("b", b)
	}
	// Other ops:   ==,!=, >=, <=
	// Logical ops: ||, &&

	var color = "1"
	switch color {
	case "red":
		gologv("color is red", "")
	case "blue":
		gologv("color is blue", "")
	case "1", "2", "3":
		gologv("color is num string", color)
	default:
		gologv("color is default", "")
	}

	// Type Switch
	var i interface{} = [1]int{1}
	switch i.(type) {
	case int:
		gologv("i is int", "")
	case string:
		gologv("i is string", "")
	case float64:
		gologv("i is float64", "")
	case []int:
		gologv("i in array", "")
	default:
		gologv("i is none", "")
	}

	// Break is default behavior in go
	// no need to explicitly mention it unlike JS, C++
	// if a case is matched, it will execute that case
	// code block and exit
	// If needed to exit inside case code block
	// we can use break...
}
