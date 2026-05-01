package main

import "fmt"

func PrintIf(str string) string {
	if len(str) >= 3 {
		return "G\n"
	} else {
		return "Invalid input\n"
	}
}

func main() {
	fmt.Print(PrintIf("Hello"))
}
