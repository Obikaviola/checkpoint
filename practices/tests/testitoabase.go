package main

import ("fmt"
		"piscine"
)

func main(){
	fmt.Println(piscine.ItoaBase(23, 4))
	fmt.Println(piscine.ItoaBase(233, 16))
	fmt.Println(piscine.ItoaBase(-23, 4))
	fmt.Println(piscine.ItoaBase(-233, 16))
}