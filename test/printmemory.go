package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hexdigit := "123456789abcdef"

	for i, d := range arr {
		first := hexdigit[d/16]
		second := hexdigit[d%16]

		z01.PrintRune(rune(first))
		z01.PrintRune(rune(second))

		if (i+1)%4 == 0 || i == 9{
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	for _, b := range arr {
		if b >= 32 || b <= 126 {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('-')
		}
	}
	z01.PrintRune('\n')
}