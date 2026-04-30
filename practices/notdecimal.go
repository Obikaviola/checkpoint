package piscine

func NotDecimal(dec string) string {
	if dec == ""{
		return "\n"
	}

	decimalP := -1

	for i := 0; i < len(dec); i++{
		if dec[i] == '.'{
			decimalP = i
			break
		}
	}

	if decimalP == -1{
		return dec + "\n"
	}

	nonZeros := false

	for i := decimalP + 1; i < len(dec); i++{
		if dec[i] != '0'{
			nonZeros = true
			break
		}
	}

	if !nonZeros {
		return dec + "\n"
	}

	str := ""
	isNegative := false

	for i := 0; i < len(dec); i++{
		if dec[i] == '-'{
			isNegative = true
			continue
		} else if dec[i] == '.'{
			continue
		} else if dec[i] >= '0' && dec[i] <= '9'{
			str += string(dec[i])
		} else {
			return dec + "\n"
		}
	}

	for i, r := range str{
		if r != '0'{
			if isNegative{
				return "-" + str[i:] + "\n"
			} else {
				return str[i:] + "\n"
			}
		}
	}
	return str
}