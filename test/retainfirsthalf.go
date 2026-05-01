package main

import "fmt"

func RetainFirstHalf(s string) string {
	if len(s) <= 1 {
		return ""
	}

	half := len(s) / 2

	if half % 2 != 0 {
		half++
	}

	return s[:half]

}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
}
