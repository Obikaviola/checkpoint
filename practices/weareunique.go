package piscine

func WeAreUnique(str1 , str2 string) int {
	if str1 == "" && str2 == ""{
		return -1
	}

	count := 0

	mp1 := make(map[string]bool)
	mp2 := make(map[string]bool)

	for _, r := range str1{
		mp1[string(r)] = true
	}

	for _, r := range str2{
		mp2[string(r)] = true
	}

	for _, r := range str1{
		if !mp2[string(r)]{
			count++
		}
	}

	for _, r := range str2{
		if !mp1[string(r)]{
			count++
		}
	}

	return count
}