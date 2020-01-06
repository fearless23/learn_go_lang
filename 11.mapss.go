package main

func mapss() {

	/*
		Maps are like maps in JS or for holding values
		of same type like student grades name-->grade,
		population state --> population
	*/

	var grades = map[string](int){
		"user1": 30,
		"user2": 23,
		"user3": 0,
		"toDel": 1,
	}
	golog("grades", grades)
	gologv("user2Grades", grades["user2"])
	delete(grades, "toDel")
	gologv("grades", grades)
	var user40, ok = grades["user40"]

	gologv("user40", user40)
	gologv("user40Exist", ok)

	var user3, ok2 = grades["user3"]
	gologv("user3", user3)
	gologv("user3Exist", ok2)

	// var y = x // y is pointer to x

}
