package piscine

func HashCode(dec string) string {
	hash := 0
	str := ""

	for _, r := range dec{
		hash = (int(r) + len(dec)) % 127

		if hash < 32 || hash == 127{
			hash += 33
		}
		
		str += string(rune(hash))
	}
	return str
}