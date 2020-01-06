package main

func arrayPart2() {
	// var t = [...]int{0, 1, 2, 3, 4, 5}
	var x = []int{0, 1, 2, 3, 4, 5}
	// Slicing any of t(array) or x(slice)
	// can be assigned to y
	golog("x", x)

	var y = x[3:] // t[3:]
	golog("y", y)

	x[4] = 40
	golog("y", y)

	y[0] = 30
	golog("y", y)
	golog("x", x)

	golog("Hello", y)

	var a = [...]int{0, 1, 2, 3, 4, 5, 6}
	var b = a[:4]
	golog("a", a)
	golog("b", b)

	// b[i] dont change b behaviour that
	// is pointing to a, if append is used
	// b stops pointing to a
	b[1] = 15
	golog("b, b#1=15", b)
	golog("a, b#1=15", a)
	a[2] = 20
	golog("a, a#3=30", a)
	golog("b, a#3=30", b)

	b = append(b, 7, 8, 9, 10)
	// CREATES NEW SLICE OR ARRAY, now b dont point to a
	// only b changees, not a
	golog("b, append", b)
	golog("a after append to b", a)

	b[3] = 121
	a[3] = 786
	// since b is now not pointing to a,
	// b, a change independently
	golog("b", b)
	golog("a", a)

	// Spread operator
	var p = []int{1, 2, 3}
	var q = []int{4, 5, 6}
	// append 1st arg should be slices only
	// spread ... can be used only on slice
	var z = append(p, q...)
	// var z = append(p,4,5,6)
	golog("p", p)
	golog("q", q)
	golog("z", z)

}
