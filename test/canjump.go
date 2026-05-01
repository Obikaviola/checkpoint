package main

func CanJump(steps []uint) bool {
	if len(steps) == 0{
		return false
	}

	lastIndex := len(steps) - 1
	currentIndex := 0

	for currentIndex < lastIndex{
		jumpSize := int(steps[currentIndex])

		if jumpSize == 0{
			return false
		}

		currentIndex = jumpSize
	}
	return currentIndex == lastIndex
}