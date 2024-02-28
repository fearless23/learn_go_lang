package basics

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
	Log("user2", grades["user2"])
	delete(grades, "toDel")
	Log("grades", grades)
	var user40, exists = grades["user40"]
	Log("user40", user40)
	Log("user40Exist", exists)

	var user3, ok2 = grades["user3"]
	Log("user3", user3)
	Log("user3Exist", ok2)

	divider()

	var gradesCopy = grades // copying map into another variable is copy-by-reference
	Log("gradesCopy", gradesCopy)
	Print("-> change original grade map with grades[user1] = 100")
	grades["user1"] = 100
	grades["user12"] = 112
	Log("grades", grades)
	Log("gradesCopy", gradesCopy)
	//
}
