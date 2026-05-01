package main

import "fmt"

func FirstWord(s string) string {

	var str string 

	for i, t := range s {
		if t == ' '{
			str = s[:i]
			break
		} else if i == len(s) - 1 {
			str = s[:i]
		}
	}
	return str + "\n"
}

func main() {
	fmt.Print(FirstWord("Hello me to the world"))
	fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}
