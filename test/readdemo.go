package main

import ("fmt"
		"os"
)

func main(){
	data, err := os.ReadFile("myFile") // reads data from file, id=f any errors found return back

	// if an error occurs, raise a panic else print data
	if err != nil {
		panic(err)
	}
	// fmt.Println(data) //returns byte data
	fmt.Println(string(data))
}