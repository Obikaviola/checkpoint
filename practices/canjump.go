package piscine

func CanJump(num []uint) bool{
	u := 0

	for i := 0; i < len(num);{
		u = int(num[i])

		if u > len(num){
			return false
		} else if u == len(num)-1{
			return true
		} else if u == 0{
			return false
		} else if u < len(num) - 1{
			i += u
		}
	}

	return false
}