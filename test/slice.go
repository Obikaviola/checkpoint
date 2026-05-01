package main

import "fmt"

func main(){
	var numbers []int
	numbers = make([]int, 5, 10)

	// slice literals
	fruits := []string{"apple", "banana", "cherry"}

	arr := [5]int{3, 4, 5, 6, 5}
	slice := arr[1:4]

	fmt.Println(numbers)
	fmt.Println(fruits)
	fmt.Println(slice)
	
	slice = append(slice, 8)
	slice = append(slice, 6, 7, 9)

	other := []int{8, 9, 10}
	slice = append(slice, other...)
	fmt.Println(slice)
	fmt.Printf("slice capacity: %v\n", cap(slice))
	fmt.Printf("numbers capacity: %v\n", cap(numbers))

	// Copying
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))

	// Copy elements
	n := copy(dst, src)

	// Partial copy
	partial := make([]int, 3)
	copy(partial, src) // use copy inorder not to reference the source else any modfication in the destination, modifies the source

	fmt.Println(src)
	fmt.Println(dst)
	fmt.Println(n)
	fmt.Println(partial)

	s := []int{0, 1, 2, 3, 4, 5}

	fmt.Println(s[1:4])
	fmt.Println(s[:3])
	fmt.Println(s[2:])
	fmt.Println(s[:])
	fmt.Println(s[1:4:5]) //capacity is 4 

	//Overlapping Slice
	letters := []string{"a", "b", "c", "d"}
	letter2 := append(letters[:3], letters[1:]...)
	fmt.Printf("overlap: %v\n", letter2)

	// Remove last of a slice
	letter2 = letter2[:len(letter2)-1]
	fmt.Printf("Remove last of a slice: %v\n", letter2)

	// Remove the first element
	letter2 = letter2[1:]
	fmt.Printf("Remove first element: %v\n", letter2)

	//Add to start of slice
	letter2 = append([]string{"z"}, letter2...)
	fmt.Println(letter2)

	var a[][]int
	b := []int{1, 2}
	c := []int{3, 4, 5}
	a = append(a, b, c)
	fmt.Println(a)

	d := make([]int, 0, 1000000)

	for i := 0; i < 10; i++{
		d = append(d, i)
	}

	fmt.Println(d)

	e := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(removeFromSliceWithOrder(e, 1))

	num := mySlice{1, 2, 3, 4, 5, 6}
	num = num.Remove(1)
	fmt.Println(num)
}

func removeFromSliceWithOrder (slice []int, index int) []int{
	return append(slice[:index], slice[index+1:]...)
}

// func removeFromSlice(slice []int, index int) []int {
// 	slice[index] = slice[len(slice)-1]
// 	return slice[:len(slice)-1]
// }


type mySlice []int
func (s mySlice) Remove(index int) []int {
	return append(s[:index], s[index + 1:]...)
}