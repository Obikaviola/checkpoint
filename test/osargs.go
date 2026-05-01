package main

import (
	"fmt"
	"os"
)

var secretPassword = "FoodBar$"

func main(){
	// var s string

	// for i := 1; i < len(os.Args); i++ {
	// 	s += os.Args[i] + " "
	// }

	// fmt.Println(s)

	// argumentWithPath := os.Args
	// argumentSlice := os.Args[1:]
	// argumentAt2 := os.Args[2]

	// fmt.Println(argumentWithPath)
	// fmt.Println(argumentSlice)
	// fmt.Println(argumentAt2)

	// var s, arg string

	// for i := 1; i < len(os.Args); i++ {
	// 	s += arg + os.Args[i] + " "
	// }
	// fmt.Println(s)

	// fmt.Println(os.Args)

	// if len(os.Args) < 2{
	// 	fmt.Println("Error")
	// 	fmt.Println("usage: %s <password>\n", os.Args[0])

	// 	os.Exit(-1)
	// }

	// password := os.Args[1]
	// if (password == secretPassword) {
	// 	fmt.Println("Congrats")
	// } else {
	// 	fmt.Println("Wrong")
	// }

	// 	fmt.Println("Using os.Args")

	// 	for _, a := range os.Args {
	// 		t := getFileType(a)
	// 		fmt.Printf("%v = %v\n", a, t)
	// 	}
	// }

	// func getFileType(f string) (t string) {
	// 	fi, err := os.Stat(f)
	// 	if nil != err {
	// 		return "file or directory doesn't exist"
	// 	}

	// 	if fi.IsDir(){
	// 		return "Is a directory"
	// 	}

	// 	return "is a file"

	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("1st argument:", os.Args[2])
	fmt.Println("1st argument:", os.Args[3])
	fmt.Println("number of items inside os.Args", len(os.Args))
}