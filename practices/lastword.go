package piscine

func LastWord(s string) string{
	str := ""
	n := len(s) - 1

	for n >= 0 && s[n] == ' '{
		n--
	}

	for i := n; i >= 0; i--{
		if i == 0{
			str += s[:n+1]
			break
		}

		if s[i] == ' '{
			str += s[i+1:n+1]
			break
		}
	}
	return str + "\n"
}
