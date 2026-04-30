package piscine

func RepeatAlpha(s string) string {
	count := 0
	str := ""

	for _, r := range s{
		if r >= 'a' && r <= 'z'{
			count = int(r - 'a' + 1)
		} else if r >= 'A' && r <= 'Z'{
			count = int(r - 'A' + 1)
		} else{
			str += string(r)
		}

		for count > 0{
			str += string(r)
			count--
		} 
	}
	return str
}
