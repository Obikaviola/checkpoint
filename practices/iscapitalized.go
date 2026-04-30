package piscine

func IsCapitalized(s string) bool {
	capitalize := false

	for i, r := range s {
		if i == 0 && (r >= 'A' && r <= 'Z'){
			capitalize = true
		} else if i != 0 && i != len(s)-1 && !((r >= 'a' && r <= 'z') || (r >= 'A' && r <='Z') || (r >= '0' && r <= '9')){
			if s[i+1] >= 'A' && s[i+1] <= 'Z' || s[i+1] >= '0' && s[i+1] <= '9'{
				capitalize = true
			} else {
				capitalize = false
			}
		}
	}
	return capitalize
}
