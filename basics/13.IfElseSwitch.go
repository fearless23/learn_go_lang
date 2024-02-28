package basics

func Ifelse() {

	if true {
		Log("$", "inside if block")
	}

	// With initializer
	if true {
		Log("$", "init")
	}
	if a := 3; a > 2 {
		Log("a inside if with init", a)
	}
	if a, b := 2, 3; b < 4 {
		Log("", "inside if block with initializer")
		logWithType("a", a)
		logWithType("b", b)
	}
	// Other ops:   ==,!=, >=, <=
	// Logical ops: ||, &&

	var color = "1"
	switch color {
	case "red":
		{
			Log("color is red", "")
		}
	case "blue":
		Log("color is blue", "")
	case "1", "2", "3":
		Log("color is num string", color)
	default:
		Log("color is default", "")
	}

	// Type Switch
	var i interface{} = [1]int{1}
	switch i.(type) {
	case int:
		Log("i is int", "")
	case string:
		Log("i is string", "")
	case float64:
		Log("i is float64", "")
	case []int:
		Log("i in array", "")
	default:
		Log("i is none", "")
	}

	// Break is default behavior in go
	// no need to explicitly mention it unlike JS, C++
	// if a case is matched, it will execute that case code block only and exit
	// If needed to exit inside case code block
	// we can use break...
}
