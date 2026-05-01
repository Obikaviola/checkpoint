package main

import "fmt"

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	var result string
	current := from

	for {
		result += string(rune(current/10 + '0'))
		result += string(rune(current%10 + '0'))

		if current == to {
			break
		}

		result += ", "

		if from < to {
			current++
		} else {
			current--
		}
	}
	return result + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}