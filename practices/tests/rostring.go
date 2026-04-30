package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]
	var res []string
	var str []string
	word := ""

	for i := 0; i < len(arg); i++ {
		if arg[i] == ' ' {
			if word != "" {
				res = append(res, word)
				word = ""
			}
		} else {
			word += string(arg[i])
		}
	}

	if word != "" {
		res = append(res, word)
	}

	str = append(str, res[1:len(res)]...)
	str = append(str, res[0])

	for i, r := range str{
		for _, char := range r{
			z01.PrintRune(char)
		}

		if i != len(str)-1{
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}
