package main

import "fmt"

func RepeatAlpha(s string) string {
	var str []rune

	for _, r := range s {
		count := 1
		if r >= 'a' && r <= 'z' {
			count = int(r - 'a' + 1)
		}else if r >= 'A' && r <= 'Z'{
			count = int(r - 'A' + 1)
		}
		for j := 0; j < count; j++ {
			str = append(str, r)
		}
	}
	return string(str)
}


func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}