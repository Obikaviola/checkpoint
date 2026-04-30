package piscine

func WordFlip(str string) string {
	result := ""

	if str == ""{
		result = "Invalid Output"
	}

	var res []string
	var revres []string
	word := ""

	for i := 0; i < len(str); i++{
		if str[i] == ' '{
			if word != ""{
				res = append(res, word)
				word = ""
			}
		} else {
			word += string(str[i])
		}
	}

	if word != ""{
		res = append(res, word)
	}

	for i := len(res)-1; i >= 0; i--{
		revres = append(revres, res[i])
	}

	for i, r := range revres{
		for _, char := range r{
			result += string(char)
		}
		if i != len(revres)-1{
			result += " "
		}
	}

	return result + "\n"
}