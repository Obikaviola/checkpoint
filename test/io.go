package main

import ("fmt"
		"io"
		"os"
		"strings"
		"net"
)

func main() {
	fileReader()
	stringReader()
	connReader()
}

func connReader(){
	conn, err := net.Dial("top", "google.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buf := make([]byte, 5)

	for {
		n, err := conn.Read(buf)

		if err == io.EOF{
			break
		}

		if err != nil{
			fmt.Println(err)
			break
		}

		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}

func stringReader(){
	s := strings.NewReader("very short but interesting string")
	buf := make ([]byte, 2)

	for {
		n, err := s.Read(buf)
		if err == io.EOF{
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}

func fileReader(){
	f, err := os.Open("myFile")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := make ([]byte, 5)//reads up to 10 bytes of data
	for{
		n, err := f.Read(buf)
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println(err)
			break
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}