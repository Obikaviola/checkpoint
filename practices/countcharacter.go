package piscine

func CountChar(str string, c rune) int {
	if str == ""{
		return 0
	}

	count := 0

	for _, u := range str {
		if u == c{
			count++
		}
	}
	return count
}