package main

import ("fmt"
		"os"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	num := os.Args[1]
	sum := 0
	intnum := Atoi(num)

	for i := intnum; i > 0; i--{
		if IsPrime(i){
			sum += i
		}
	}

	fmt.Println(sum)
}

func Atoi(s string) int{
	var n int

	for _, r := range s{
		n = n*10 + int(r - '0')
	}
	return n
}

func IsPrime(n int) bool{
	if n < 2{
		return false
	}

	for i := 2; i*i <= n; i++{
		if n%i == 0{
			return false
		}
	}
	return true
}