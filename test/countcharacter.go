package main

import "fmt"

func CountChar(str string, c rune) int {
	count := 0

	for _, s := range str {
		if s == c {
			count++
		}
	}
	return count
}

// func main(){
// 	fmt.Println(CountChar("    ", ' '))
// }