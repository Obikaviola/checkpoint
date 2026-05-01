package main

import "fmt"

func HashCode(dec string) string {
	n := len(dec)
	result := make([]rune, n)

	for i, r := range dec {
		hash := (int(r) + n) % 127

		if hash < 32 || hash == 127{
			hash += 33
		}
		result[i] = rune(hash)
	}
	return string(result)

}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}