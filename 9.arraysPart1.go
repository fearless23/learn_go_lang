package main

import (
	"fmt"
)

func arrPart1() {
	threeNums := [3]int{1, 2, 3}
	fmt.Printf("threeNums: %v, %T\n", threeNums, threeNums)

	var idx2 = threeNums[2]
	fmt.Printf("idx2: %v, %T\n\n", idx2, idx2)

	nums := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("nums: %v, %T\n", nums, nums)
	nums[0] = 42
	fmt.Printf("nums: %v, %T\n\n", nums, nums)

	var students [3]string
	students[1] = "John"
	fmt.Printf("students: %v, %T\n", students, students)
	fmt.Printf("Length: %v\n\n", len(students))

	var matrix [2][2]int
	matrix[0] = [2]int{1, 0}
	matrix[1] = [2]int{0, 1}
	fmt.Printf("matrix: %v, %T\n\n", matrix, matrix)

	var a = [...]int{1, 2, 3, 4}
	fmt.Printf("a start: %v, %T\n", a, a)
	var b = a // Copies array
	b[2] = 42
	a[2] = 72
	fmt.Printf("a after: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n\n", b, b)

	var c = &a // Pointer to a,
	// change in c changes a, eventually c
	c[2] = 21
	fmt.Printf("a after &a: %v, %T\n", a, a)
	fmt.Printf("c: %v, %T\n\n", c, c)

	// Slice, empty [] before dec
	var x = []int{1, 2, 3, 4, 5}
	var y = x
	// Any variable equals a slice
	// is actually a pointer, unlike array is does
	// not create copy
	x[1] = 21
	y[2] = 42
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)
}
