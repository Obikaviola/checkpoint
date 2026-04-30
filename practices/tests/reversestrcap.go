package main

import ("github.com/01-edu/z01"
		"os"
)

func main() {
	if len(os.Args) == 1{
		return
	}

	var arg []string

	arg = os.Args[1:]
	str := ""

	for _, r := range arg{
		for i := 0; i < len(r);{
			if i == 0 && r[i] == ' '{
				if r[i+1] >= 'A' && r[i+1] <= 'Z'{
					str += " " + string(r[i+1] + 32)
					i+=2
					continue
				} else{
					str += string(r[i])
				}
			} else if i == 0 && r[i] != ' '{
				if r[i] >= 'A' && r[i] <= 'Z'{
					str += string(r[i] + 32)
				} else {
					str += string(r[i])
				}
			} else if i == len(r)-1{
				if r[i] >= 'a' && r[i] <= 'z'{
					str += string(r[i] - 32)
					break
				} else{
					str += string(r[i])
					break
				}
			} else if i != 0 && r[i+1] == ' '{
				if r[i] >= 'a' && r[i] <= 'z'{
					str += string(r[i] - 32) 
					str += " "
					i++
					continue
				} else{
					str += string(r[i])
				}
			} else {
				if r[i] >= 'A' && r[i] <= 'Z'{
					str += string(r[i] + 32)
				} else{
					str += string(r[i])
				}
			}
			i++
		}
		str += "\n"
	}
	
	for _, r := range str{
		z01.PrintRune(r)
	}
}
