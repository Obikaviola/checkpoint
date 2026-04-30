package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3{
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	printed := make(map[string]bool)

	for _, r1 := range s1{
		if !printed[string(r1)]{
			for _, r2 := range s2{
				if r1 == r2{
					z01.PrintRune(r1)
					printed[string(r1)] = true
					break
				}
			}
		}
	}
	z01.PrintRune('\n')
}