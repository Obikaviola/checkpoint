package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.CamelToSnakeCase("HelloWorld"))
	fmt.Println(piscine.CamelToSnakeCase("helloWorld"))
	fmt.Println(piscine.CamelToSnakeCase("camelCase"))
	fmt.Println(piscine.CamelToSnakeCase("CAMELtoSnakeCASE"))
	fmt.Println(piscine.CamelToSnakeCase("CameltoSnakeCase"))
	fmt.Println(piscine.CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(piscine.CamelToSnakeCase("hey2"))
	fmt.Println(piscine.CamelToSnakeCase("helloWOrld"))
}