package piscine

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "invalid\n"
	}

	str := ""

	if from > to {
		for i := from; i >= to; i--{
			str += string(rune(i/10) + '0')
			str += string(rune(i%10) + '0')

			if i != to {
				str += ", "
			}
		}
	}

	if from <= to{
		for i := from; i <= to; i++{
			str += string(rune(i/10) + '0')
			str += string(rune(i%10) + '0')

			if i != to{
				str += ", "
			}
		}
	}

	return str + "\n"
}