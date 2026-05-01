package main

import "fmt"

type SliceHeader struct{
	Data uintptr
	Len int 
	Cap int
}

func main() {
	makeSlice := make([]int, 0)

	if makeSlice == nil {
		fmt.Println("makeSlice is nil")
	} else {
		fmt.Println("makeSlice is not nil")
	}

	// fmt.Printf("%#v\n", makeSlice)
	fmt.Println(makeSlice)

	newSlice := new([]int)

	if *newSlice == nil{
		fmt.Println("newSlice is nil")
	} else {
		fmt.Println("newSlice is not nil")
	}

	// fmt.Printf("%#v\n", *newSlice)
	fmt.Println(*newSlice)

	var nilSlice []int // len, cap, append, copy, for...range works and it is equal to nil
	emptySlice := make([]int, 0, 10) // len, cap, append, copy, for...range works but it is not equal to nil
	nilSlice = append(nilSlice, 1)
	emptySlice = append(emptySlice, 2, 3, 4)
	fmt.Printf("capacity = %d\n", cap(nilSlice))
	fmt.Printf("capacity = %d\n", cap(emptySlice))

	// make returns the value, the new returns the pointer
	/**
	make - flexibility, control, optimisation make([]type, len, capacity)
	new - 
	*/
}
