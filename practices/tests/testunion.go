package main

import ("os"
		"fmt"
)

func main(){
	if len(os.Args) != 3{
		return
	}

	arg := os.Args[1:]

	str := ""
	mp1 := make(map[string]bool)

	for _, r := range arg{
		for i := 0; i < len(r); i++{
			if !(mp1[string(r[i])]){
				mp1[string(r[i])] = true
			}
		}
	}

	for _, r := range arg{
		for _, char := range r{
			if mp1[string(char)]{
				str += string(char)
				delete(mp1, string(char))
			}
		} 
	}
	fmt.Println(str)
}
