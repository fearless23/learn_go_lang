package src

func Ifelse() {

	if true {
		log("$", "inside if block")
	}

	// With initializer
	if true {
		log("$", "init")
	}
	if a := 3; a > 2 {
		log("a inside if with init", a)
	}
	if a, b := 2, 3; b < 4 {
		log("", "inside if block with initializer")
		logWithType("a", a)
		logWithType("b", b)
	}
	// Other ops:   ==,!=, >=, <=
	// Logical ops: ||, &&

	var color = "1"
	switch color {
	case "red":
		{
			log("color is red", "")
		}
	case "blue":
		log("color is blue", "")
	case "1", "2", "3":
		log("color is num string", color)
	default:
		log("color is default", "")
	}

	// Type Switch
	var i interface{} = [1]int{1}
	switch i.(type) {
	case int:
		log("i is int", "")
	case string:
		log("i is string", "")
	case float64:
		log("i is float64", "")
	case []int:
		log("i in array", "")
	default:
		log("i is none", "")
	}

	// Break is default behavior in go
	// no need to explicitly mention it unlike JS, C++
	// if a case is matched, it will execute that case code block only and exit
	// If needed to exit inside case code block
	// we can use break...
}
