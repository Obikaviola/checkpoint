package piscine

func FirstWord(s string) string {
	str := ""
	n := 0

	if len(s) == 0{
		return "\n"
	}

	for n < len(s)-1 && s[n] == ' '{
		n++
	}

	for i, r := range s{
		if i != 0 && r == ' '{
			str += s[n:i]
			break
		}
		if i == len(s)-1{
			str += s[n:]
		}
	}
	
	return str + "\n"
}