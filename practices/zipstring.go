package piscine

func ZipString(s string) string {
	str := ""
	index := 0

	for i := 0; i < len(s);{
		count := 0
		char := s[i]

		j := i
		for j < len(s) && s[j] == char{
			count++
			index++
			j++
		}

		str += string(rune(count + '0')) + string(char)
		i = index 
	}

	return str 
}
