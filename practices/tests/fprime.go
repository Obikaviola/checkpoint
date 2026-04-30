package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

func main(){
	if len(os.Args) != 2{
		return
	}

	arg := os.Args[1]

	numb := ToInt(arg)

	if numb <= 1{
		z01.PrintRune('0')
	}

	factor := 2
	first := true

	for numb > 1{
		if numb%factor == 0{
			if !first{
				z01.PrintRune('*')
			}
			z01.PrintRune(rune(factor + '0'))
			numb /= factor
			first = false
		} else {
			factor++
		}
	}
	z01.PrintRune('\n')
}


func ToInt(n string) int{
	var result int

	for _, r := range n{
		if !(r >= '0' && r <= '9'){
			return 0
		}

		result = result*10 + int(r - '0')
	}
	return result
}