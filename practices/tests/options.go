package main

import (
	"fmt"
	"os"

	// "github.com/01-edu/z01"
)

func main(){
	if len(os.Args) < 2{
		fmt.Println("Invalid Option") 
		return
	}

	opt := "options: abcdefghijklmnopqrstuvwxyz"

	arg := os.Args[1:]
	index := 0
	bit := [32]int{}

	for _, r := range arg{
		if len(r) < 2{
			fmt.Printf("Invalid Option\n")
			return
		} else if r[0] != '-' || r[1] == 'h'{
			fmt.Printf(opt + "\n")
			return
		}
		for _, char := range r[1:]{
			index = 31 - int(char - 'a')
			bit[index] = 1
		}
	}

	for i := 0; i < 32; i++{
		fmt.Print(bit[i])

		if (i+1)%8 == 0 && i != 31{
			fmt.Print(" ")
		}
	}
	fmt.Println()
}