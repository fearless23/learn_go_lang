package src

import (
	"fmt"
)

func ArrayPart1() {
	var threeNums = [3]int{1, 2, 3}
	fmt.Printf("threeNums: %v, %T\n", threeNums, threeNums)

	var idx2 = threeNums[2]
	fmt.Printf("idx2: %v, %T\n\n", idx2, idx2)

	nums := [7]int{1, 2, 3, 4, 5, 6}
	// if length is less than items - error
	// if length is more than items; extra items initialized to 0
	fmt.Printf("nums: %v, %T\n", nums, nums)
	nums[0] = 42
	fmt.Printf("nums: %v, %T\n\n", nums, nums)

	var students [3]string // 3 items initialized to empty string
	students[1] = "John"
	fmt.Printf("students: %v, %T\n", students, students)
	fmt.Printf("Length: %v\n", len(students))
	fmt.Printf("Students[0]: %v %T\n\n", students[0], students[0])

	var matrix [2][2]int
	matrix[0] = [2]int{1, 0}
	matrix[1] = [2]int{0, 1}
	fmt.Printf("matrix: %v, %T\n\n", matrix, matrix)

	var a = [...]int{1, 2, 3, 4}
	fmt.Printf("a start: %v, %T\n", a, a)
	var b = a // Copies array (deep copy aka clone of top level)
	b[2] = 42
	a[2] = 72
	fmt.Printf("a after: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n\n", b, b)

	var c = &a // Pointer to a,
	// change in c changes a, eventually c
	c[2] = 21
	fmt.Printf("a after &a: %v, %T\n", a, a)
	fmt.Printf("c: %v, %T\n\n", c, c)

	// Slice, empty [] before declaration
	var x = []int{1, 2, 3, 4, 5}
	var y = x
	// copying slice is copy-by-reference
	x[1] = 21
	y[2] = 42
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)
}

// var a = [4]int{1, 2, 3, 4} // ARRAY
// var a = [...]int{1, 2, 3, 4} // ARRAY
// var a = []int{1, 2, 3, 4} // SLICE

// b = a // ARRAY -> clone i.e copy by value; SLICE -> copy by reference
