package piscine

func FifthAndSkip(str string) string {
	s := ""

	if str == ""{
		return "\n"
	}
	
	if len(str) < 5 {
		return "Invalid Input\n"
	}

	var container []rune

	for _, r := range str {
		if r != ' '{
			container = append(container, r)
		}
	}

	count := 0

	for i := 0; i < len(container); i++{
		if count < 5{
			s += string(container[i])
		}

		if count == 5{
			s += " "
			count = 0
			continue
		} else {
			count++
		}
	}

	return s + "\n"
}