package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]
	var str []string
	var revstr []string
	word := ""

	for i := 0; i < len(arg); i++{
		if arg[i] == ' '{
			if word != ""{
				str = append(str, word)
				word = ""
			}
		} else {
			word += string(arg[i])
		}
	}

	if word != ""{
		str = append(str, word)
	}

	for i := len(str) - 1; i > 0; i-- {
		if i != 1{
			revstr = append(revstr, str[i], " ")
		} else {
			revstr = append(revstr, str[i])
		}
	}

	for _, r := range revstr {
		for _, char := range(r){
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
