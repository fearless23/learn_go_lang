package src

func Maps() {

	/*
		- Maps are like maps/object in Javascript
		- It is used for holding key-values of some key type and some value type like
			- student grades: name --> grade
			- population: state --> population
	*/

	var grades = map[string](int){
		"user1": 30,
		"user2": 23,
		"user3": 0,
		"toDel": 1,
	}
	logWithType("grades", grades)
	log("user2", grades["user2"])
	delete(grades, "toDel")
	log("grades", grades)
	var user40, exists = grades["user40"]
	log("user40", user40)
	log("user40Exist", exists)

	var user3, ok2 = grades["user3"]
	log("user3", user3)
	log("user3Exist", ok2)

	divider()

	var gradesCopy = grades // copying map into another variable is copy-by-reference
	log("gradesCopy", gradesCopy)
	print("-> change original grade map with grades[user1] = 100")
	grades["user1"] = 100
	grades["user12"] = 112
	log("grades", grades)
	log("gradesCopy", gradesCopy)
	//
}
