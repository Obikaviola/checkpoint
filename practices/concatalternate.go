package piscine

func ConcatAlternate(slice1,slice2 []int) []int {
	var slice []int
	var first []int
	var second []int

	if len(slice1) == 0{
		return slice2
	}

	if len(slice2) == 0{
		return slice1
	}

	if len(slice1) == len(slice2){
		first = slice1
		second = slice2
	}

	if len(slice1) > len(slice2){
		first = slice1
		second = slice2
	} else {
		first = slice2
		second = slice1
	}
	minlen := len(second)

	for i := 0; i < minlen; i++{
		slice = append(slice, first[i], second[i])
	}

	slice = append(slice, first[minlen:]...)

	return slice
}