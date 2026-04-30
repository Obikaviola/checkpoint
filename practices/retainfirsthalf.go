package piscine

func RetainFirstHalf(str string) string {
	if len(str) == 0 {
		return ""
	}

	if len(str) == 1 {
		return str
	}

	len := len(str)

	if len%2 != 0 {
		len--
	}

	return str[:len/2]
}
