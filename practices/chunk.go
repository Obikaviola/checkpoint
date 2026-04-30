package piscine

import "fmt"

func Chunk(slice []int, size int) {
	chunk := make([][]int, 0)
	end := 0

	if len(slice) == 0{
		fmt.Println("[]")
		return
	}

	if size == 0{
		fmt.Println(slice)
		return
	}

	for i := 0; i < len(slice); i += size{
		end += size

		if end > len(slice){
			end = len(slice)
		}
		chunk = append(chunk, slice[i:end])
	}
	fmt.Println(chunk)
}