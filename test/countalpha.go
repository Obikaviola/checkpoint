// package piscine
package main

import (
	"fmt"
)

func CountAlpha(s string) string {
	count := 0
	var str string

	// for _, t := range s {
	// 	if t >= 'a' && t <= 'z' || t >= 'A' && t <= 'Z' {
	// 		count++
	// 	}
	// }

	for _, t := range s {
		if t >= 'a' && t <= 'z' || t >= 'A' && t <= 'Z' {
			count++
		}
	}
	
	if count < 3 {
		str = "G\n"
	} else {
		str = "Invalid\n"
	}

	return str
}

// half := len(s)/2
// return str[:half]

// func main() {
// 	fmt.Print(CountAlpha("abcdef"))
// }
