package main

func pointterss() {

	a := 42
	b := &a
	a = 12
	*b = 14
	golog("a", a)
	golog("b", b)
	golog("*b", *b)

	x := [3]int{1, 2, 3}
	y := &x[0]
	z := &x[1]
	*y = 56
	golog("x", x)
	golog("y", y)
	golog("z", z)

	var k *int // Nil pointer
	golog("k", k)

	// Assign pointer to some val

}
