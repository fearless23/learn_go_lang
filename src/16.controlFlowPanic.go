package src

func CfPanic() {
	/*
		Error Handling - Panic
		Open non-existence file is not error or exception in GO
	*/

	/*
		a, b := 1, 0
		ans := a / b // division by zero; go will panic
		gologv("ans", ans)
		panic("Something wrong!")
	*/

	/*
		all defer before panic will execute, but any
		defers after panic will not. Once panic is
		detected, go reverse its direction and execute defer.
		If no panic is detected, go executes until end of main
		and then reverses.
	*/

	/*
		a, b := 1, 0
		gologv("a", a)
		panic("Something wrong!")
		gologv("b", b)
	*/

	/*
		panic in other func in main...

		panic inside other func will be	same as main. Once other func has
		panic main stops as well and reverses till 1st defer in main.
		However, after panic inside other func, it first reverses
		in other func, if that defer in other func can recover the
		error, main continues as shown with willRecover
	*/

	a, b := 1, "end"
	Log("a", a)
	willRecover()
	Log("", b)
}

func willPanic() {
	a := 1
	Log("willPanic a", a)
	panic("Error from willPanic Func")
	// gologv("End of insideFn, never reached", "")
}

func willRecover() {
	a := 1
	Log("inside willRecover a", a)

	defer recoverErr()

	panic("Error from willRecover Func")
	// gologv("End of insideFn, never reached", "")
}

func cantRecover() {
	a := 1
	Log("cantRecover a", a)

	defer Repanic()

	panic("Error from cantRecover Func")
	// gologv("End of insideFn, never reached", "")
}

func recoverErr() {
	err := recover()
	if err != nil {
		Log("$", "Error Recovered")
	}
}

func Repanic() {
	err := recover()
	if err != nil {
		Log("$", "Error Recovered")
		Log("$", "Repanicking...")
		panic(err)
	}
}
