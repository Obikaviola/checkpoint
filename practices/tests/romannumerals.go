package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2{
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	arg := os.Args[1]

	num := 0
	for _, r := range arg{
		if r < '0' || r > '9'{
			fmt.Println("ERROR: cannot convert to roman digit")
			return
		}
		num = num * 10 + int(r - '0')
	}

	if num == 0 || num >= 4000{
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	calcs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	symbols := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}

	calculation := ""
	result := ""

	for i := 0; i < len(values); i++{
		for num >= values[i]{
			if result != ""{
				result += "+"
			}

			result += symbols[i]
			calculation += calcs[i]
			num -= values[i]
		}
	}

	fmt.Println(result)
	fmt.Println(calculation)
}
