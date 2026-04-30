package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main(){
	if len(os.Args) != 3{
		return
	}

	first := os.Args[1]
	second := os.Args[2]

	if first == ""{
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}

	j := 0

	for i := 0; i < len(second); i++{
		if first[j] == second[i]{
			j++
		}

		if j == len(first)-1{
			break
		}
	}

	if j == len(first){
		z01.PrintRune('1')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	}
}