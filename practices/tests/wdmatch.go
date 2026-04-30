package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main(){
	if len(os.Args) != 3{
		return
	}

	word := ""

	first := os.Args[1]
	second := os.Args[2]

 	i, j := 0, 0

	for i < len(first) && j < len(second){
		if first[i] == second[j]{
			word += string(first[i])
			j++
			i++
		} else {
			j++
		}
	}

	if word == first{
		for _, r := range word{
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	} else {
		return
	}
}