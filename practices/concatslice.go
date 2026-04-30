package piscine

func ConcatSlice(slice1, slice2 []int) []int {
	var slice []int

	slice = append(slice, slice1...)
	slice = append(slice, slice2...)

	return slice
}