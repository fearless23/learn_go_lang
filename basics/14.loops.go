package basics

func Loops() {

	// for init i; cond for i; next val of i
	// i is scoped to for loop only...
	for i := 0; i < 5; i++ {
		Log("i", i)
	}

	// j init outside, incr inside
	// thus only cond is left
	j := 0
	for j < 5 {
		Log("j", j)
		j++
	}

	/*
		for loop requires 3 things
		1. init
		2. cond
		3. next val

		i:=0
		for ;i<5;i++{}

		i:0
		for i < 5 {
			i++
		}

		i := 0
		for {
			i++
			if i > 6 {
				break
			}
		}
		gologv("I", i) // i = 7
	*/
	// break and continue work same in JS
	// break out of a outer loop from inside
	// a nested loop, use labeled loop

Main:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 5; j++ {
			if i*j > 12 {
				Log("breaking at", i*j)
				break Main
			}
		}
	}

	s := [...]int{1, 1, 2, 3, 5, 8, 13, 21, 32, 55, 87}
	for index, value := range s {
		Log("index", index)
		Log("value", value)
	}

	// quick loop for 5 times
	for index := range [5]int{} {
		Log("index", index)
	}

	m := map[string]int{
		"user1": 1,
		"user2": 2,
	}
	for k, v := range m {
		Log("key", k)
		Log("val", v)
	}
}
