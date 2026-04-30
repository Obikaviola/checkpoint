package main

import (
	// "debug/plan9obj"
	"fmt"
	"os"
)

func main(){
	if len(os.Args) != 3{
		return
	}

	numstr := os.Args[1]

	for _, r := range os.Args[2]{
		if r == ' '{
			fmt.Println("Invalid target sum")
			return
		}
	}
	

	target := os.Args[2]
	targetSum := 0
	
	for _, r := range target{
		targetSum = targetSum * 10 + int(r - '0')
	}
	

	var numslice []int
	var slice [][]int
	n := 0

	for i := 0; i < len(numstr);{
		if i == len(numstr)-1 && string(numstr[i]) != "]"{
			fmt.Println("Invalid input")
			return
		} else if i == len(numstr)-1 && string(numstr[i]) == "]"{
			break
		}else if i == 0 && string(numstr[i]) != "["{
			fmt.Println("Invalid input")
			return
		} else if i == 0 && string(numstr[i]) == "["{
			 i++
			continue
		} else if string(numstr[i]) == " " || string(numstr[i]) == "," {
			i++
			continue
		} else if numstr[i] == '-'{
			if i+1 < len(numstr) && numstr[i+1] >= '0' && numstr[i+1] <= '9'{
				n = int(numstr[i+1] - '0')
				numslice = append(numslice, -n)
				i+= 2
				continue
			}
		} else if i != 0 && !(numstr[i] >= '0' && numstr[i] <= '9'){
			fmt.Printf("Invalid input: %v \n", string(numstr[i]))
			return
		} else {
			numslice = append(numslice, int(numstr[i] - '0'))
		}
		 i++
	}


	for i := 0; i < len(numslice); i++{
		for j := i+1; j < len(numslice); j++{
			if numslice[i] + numslice[j] == targetSum{
				s := []int{i, j}
				slice = append(slice, s)
			}
		}
	}

	fmt.Println(slice)
}