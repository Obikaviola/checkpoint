package piscine

func CamelToSnakeCase(s string) string{
	if s == "" {
		return s
	}

	snakecase := ""

	for i, r := range s {
		if r >= '0' && r <= '9' {
			return s
		}
		if  i != 0 && (r >= 'A' && r <= 'Z') && (s[i-1] >= 'A' && s[i-1] <= 'Z') {
			return s
		} else if i != 0 && (r >= 'A' && r <= 'Z'){
			snakecase = snakecase + "_" + string(r)
		} else {
			snakecase = snakecase + string(r)
		}
	}
	return snakecase
}