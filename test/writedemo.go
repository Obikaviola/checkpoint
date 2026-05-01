package main

import (
		"os"
)

func main(){
	data := []byte("Hello world\n")
	os.WriteFile("myFile",data, 0644) //the name of the file yuo waant to create, the byte array, the file permission
	// Create or replace the myFile with the given data
}