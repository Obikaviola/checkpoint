package main

import("fmt"
	"os"
	"bufio"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type your birth year: ")
	scanner.Scan()
	input := strconv.ParseInt(scanner.Text())
	fmt.Printf("You will be %d years old at the end of 2026\n", 2026 - input )
}