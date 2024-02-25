package src

import (
	"fmt"
)

// Print message with any value
func log(a string, b interface{}) {
	fmt.Printf(a+": %v\n", b)
}

// Print message with any value and its type
func logWithType(message string, value interface{}) {
	fmt.Printf(message+"(%T): %v\n", value, value)
}

// Print value with Println
func print(value any) {
	fmt.Println(value)
}

func divider() {
	print("----------------------------")
}
