package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"

	for i := 0; i < len(arr); i++{
		first := hex[arr[i]/16]
		second := hex[arr[i]%16]

		z01.PrintRune(rune(first))
		z01.PrintRune(rune(second))

		if (i+1)%4 == 0 || i == 9{
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	for _, r := range arr{
		if r >= 32 && r < 127{
			z01.PrintRune(rune(r))
		} else{
			z01.PrintRune('.')
		}
	}
	
	z01.PrintRune('\n')
}