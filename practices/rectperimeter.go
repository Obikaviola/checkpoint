// package main

// import "fmt"
package piscine

func RectPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}

	return 2*(w + h)
}

// str := "12345"
// break := rune(str)
// break[0] = '6'
// string(break)

// for_, char := range str{
// 	int(char - '0')
// }

// func main() {
// 	n := "123"
// 	var digit int

// 	// for i := 0; n > 0; i++ {
// 	// 	num := n%10
// 	// 	digit = string(rune(num) + '0') + digit
// 	// 	n /= 10
// 	// }
// 	for _, r := range n {
// 		digit = digit*10 + int(r - '0')
// 	}
// 	fmt.Println(digit)
// }

/**
To convert from string to int
1. Change to rune - using for _, r := range str
2. create a rune slice variable var runes []rune
3. 
int(char - 'a') - to convert to small letter

*/