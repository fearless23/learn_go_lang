package src

func CfDefer() {
	/*
		Go executes left to right and
		top to bottom.
		Async behaviour is done in few wats
	*/
	a := 1
	gologv("a", a)

	b := 2
	defer gologv("b", b)

	c := 3
	gologv("c", c)

	/*
		Defer take away the func from main flow
		and keeps it somewhere.
		Once main func is complete(just), before main
		func exit, defer start executing func added
		by using defer keyword in LIFO fashion.
		Visually if all defers at taken out, they will
		be executed bottom to top.
	*/
	/*
		url := "http://www.google.com/robots.txt"
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		robots, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		gologs("robots", robots)
	*/

	x := "start"
	defer gologv("x", x)
	x = "end"
	// When defer was declared, value of x
	// is eagerly used. Thus, it will print
	// "start" not "end"

}
