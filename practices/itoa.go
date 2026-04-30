package piscine

func Itoa(n int) string {
	str := ""

	isNegative := false

	if n == 0{
		return "0"
	}

	if n < 0 {
		n = -n
		isNegative = true
	}

	for n > 0{
		str = string(rune(n%10) + '0') + str
		n /= 10 
	}

	if isNegative{
		return "-" + str
	}

	return str
}
