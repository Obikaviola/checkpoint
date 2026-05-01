package main

func main() {
	sentence := []byte("Hello, World!")

	counter := 0

	for index, char := range sentence {
		counter++
		println(index, char)
	}
	println("Total characters:", counter)
}