package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) != 2{
		return
	}

	arg := os.Args[1]
	str := ""

	for i := 0; i < len(arg); i++{
		if arg[i] == 'a' || arg[i] == 'e' || arg[i] == 'i' || arg[i] == 'o' || arg[i] == 'u' || arg[i] == 'A' || arg[i] == 'E' || arg[i] == 'I' || arg[i] == 'O' || arg[i] == 'U'{
			str += arg[i:] + arg[:i] + "ay"
			break
		} 
	}

	if str == ""{
		fmt.Printf("No vowel\n")
	} else {
		fmt.Printf(str + "\n")
	}
}