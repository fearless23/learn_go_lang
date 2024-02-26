package src

func ArrayPart2() {
	var originalArray = [...]int{0, 1, 2, 3, 4, 5}
	var originalSlice = []int{0, 1, 2, 3, 4, 5}
	// Slicing can be done to t(array) or x(slice)
	// slicing is copy-by-reference (for originalArray as well originalSlice)
	logWithType("originalArray", originalArray)
	logWithType("originalSlice", originalSlice)

	var sliceOfArray = originalArray[3:]
	logWithType("slice of array (created)", sliceOfArray)
	var sliceOfSlice = originalSlice[3:]
	logWithType("slice of slice (created)", sliceOfSlice)

	originalArray[4] = 40
	logWithType("sliceOfArray - after originalArray changed", sliceOfArray)
	// arraySlice changes with change to originalArray; since slicing is copy-by-reference

	originalSlice[4] = 40
	logWithType("sliceOfSlice - after originalSlice changed", sliceOfSlice)
	// sliceOfSlice changes with change to originalSlice; since slicing is copy-by-reference

	sliceOfArray[0] = 30
	logWithType("sliceOfArray - after sliceOfArray changed", sliceOfArray)
	logWithType("originalArray - after sliceOfArray changed", originalArray)
	// originalArray changes with change to sliceOfArray; since slicing is copy-by-reference

	// TLDR; slicing is copy-by-reference; thus slice points to original data-structure;
	// So, change in original changes slice and vice-versa...
	Print("------------------------------------------------")

	var original = [...]int{0, 1, 2, 3, 4, 5, 6}
	var slice = original[:4]
	logWithType("original", original)
	logWithType("slice", slice)

	// b[i] dont change b behaviour that
	// is pointing to a, if append is used
	// b stops pointing to a
	slice[1] = 15
	Print("--> slice[1] = 15")
	logWithType("slice", slice)
	logWithType("original", original)
	original[2] = 20
	Print("--> original[2] = 20")
	logWithType("original", original)
	logWithType("slice", slice)

	Print("------------------------------------------------")

	Print("--> append(slice, 7, 8, 9, 10)")
	slice = append(slice, 7, 8, 9, 10) // []int
	// --> append creates new slice in memory, now `slice` do not point to original
	// since we can append any number of int, it cannot properly point to original; thus it creates
	// a completely new slice
	logWithType("new slice", slice)
	logWithType("original", original)

	Print("--> slice[3] = 121 and original[3] = 786")
	slice[3] = 121
	original[3] = 786
	logWithType("slice", slice)
	logWithType("original", original)

	Print("---------------- SPREAD operator --------------------------------")
	var p = []int{1, 2, 3}
	var q = []int{4, 5, 6}
	// append`s 1st arg should be slice only
	// remaining arguments is spread of same type i.e int in above case
	// since; we need to spread int as arguments to append function we use spread operator
	// spread operator i.e `...` can be used on slice data-structure only
	// we can write ==> var z = append(p, 4, 5, 6)
	// or we can use spread operator
	var z = append(p, q...)
	logWithType("p", p)
	logWithType("q", q)
	logWithType("z", z)

}
