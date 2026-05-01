// package piscine
package main

// import "fmt"

// import "github.com/01-edu/z01"

func checkNumber(arg string) bool {
	for _, s := range arg{
		if s >= '0' && s <= '9'{
			return true
		}
	}
	return false
}

// func main(){
// 	fmt.Println(checkNumber("Hello1"))
// }