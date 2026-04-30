package piscine

func ItoaBase(value, base int) string {
	if base < 2 || base > 16{
		return "Invlaid input\n"
	}

	isNegative := false
	var n uint

	if value < 0{
		isNegative = true
		n = uint(-value)
	} else {
		n = uint(value)
	}

	equiv := "0123456789ABCDEF"
	res := ""
	reverseres := ""
	ubase := uint(base)

	for n > 0{
		res += string(equiv[n%ubase])
		n /= ubase
	}

	for i := len(res)-1; i >= 0; i--{
		reverseres += string(res[i])
	}

	if isNegative{
		return "-" + reverseres
	}
	return reverseres
}