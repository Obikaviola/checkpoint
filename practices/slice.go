package piscine

func Slice(a []string, nbrs... int) []string{
	first := nbrs[0]
	second := 0
	two := false

	if len(nbrs) == 2{
		second = nbrs[1]
		two = true
	}

	if first < 0{
		first = len(a) + first
	}

	if second < 0{
		second = len(a) + second
	}

	if two{
		if second > 0 {
			return a[first:second]
		}else if second == 0{
			return nil
		}
	}
	return a[first:]
}