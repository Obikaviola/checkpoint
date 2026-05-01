package main

import (
	// "github.com/01-edu/z01"
	"fmt"
)

func CamelToSnakeCase(s string) string {

	str := []rune(s)
	var result []rune
	var previous rune

	if s == "" || s == " " {
		return " "
	}

	for i, t := range str {
		if !(t >= 'a' && t <= 'z' || t >= 'A' && t <= 'Z'){
			return s
		}
		if i == 0 && t >= 'A' && t <= 'Z'{
			return s
		} else if i == len(s)-1 && t >= 'A' && t <= 'Z' {
			return s
		} else if (t >= 'A' && t <= 'Z') && (previous >= 'A' && previous <= 'Z') {
			return s
		} else if t >= 'A' && t <= 'Z'{
			result = append(result, '_')
			result = append(result, str[i])
		} else {
			result = append(result, str[i])
		}
		previous = t
	}
	return string(result)
}

func main() {
	fmt.Println(CamelToSnakeCase("hellOWorld"))
}
