package piscine

func SaveAndMiss(arg string, num int) string {
	str := ""
	save := true
	start := 0

	if arg == ""{
		return ""
	}

	if num <= 0{
		return arg
	}

	i := 0

	for i < len(arg){
		i += num
		if i < len(arg){
			if save{
				str += arg[start:i]
				save = false
				start = i
			} else {
				save = true
				start = i
			}
		} else {
			str += arg[start:]
		}
	}

	return str
}