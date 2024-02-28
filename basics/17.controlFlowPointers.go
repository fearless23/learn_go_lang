package basics

func Pointers() {

	a := 42
	b := &a
	a = 12
	*b = 14
	logWithType("a", a)
	logWithType("b", b)
	logWithType("*b", *b)

	x := [3]int{1, 2, 3}
	y := &x[0]
	z := &x[1]
	*y = 56
	logWithType("x", x)
	logWithType("y", y)
	logWithType("z", z)

	var k *int // Nil pointer
	logWithType("k", k)

	// Assign pointer to some val

}
