package main

func main() {

	a := 42
	b := &a
	a = 12
	*b = 14
	golog("a", a)
	golog("b", b)
	golog("*b", *b)

}
