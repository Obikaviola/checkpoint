package piscine

func RevConcatAlternate(slice1,slice2 []int) []int {
	var first []int
	var second []int
	var result []int

	if len(slice1) == len(slice2){
		first = slice1
		second = slice2
	}

	if len(slice1) > len(slice2) {
		first = slice1
		second = slice2
	} else {
		first = slice2
		second = slice1
	}

	minLen := len(second)

	if len(first) == len(second){
		for i := len(second)-1; i >= 0; i--{
			result = append(result, slice1[i])
			result = append(result, slice2[i])
		}
	}

	if len(first) > len(second){
		for i := len(first)-1; i >= minLen; i--{
			result = append(result, first[i])
		}

		for i := minLen-1; i >= 0; i--{
			result = append(result, slice1[i])
			result = append(result, slice2[i])
		}
	}
	
	return result
}