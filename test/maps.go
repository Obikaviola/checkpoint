package main

import "fmt"

func main() {
	score := make(map[string]int)
	score["hitesh"] = 89
	score["josh"] = 88
	score["hope"] = 88
	fmt.Println(score)
	fmt.Println(score["hope"])

	delete(score, "hope")
	k := "kate"
	score[k] = 45
	fmt.Println(score)

	for k, v := range score{
		fmt.Printf("score of %v: %v\n", k, v) // never rely on the order
	}

	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])
	
	for key, value := range menu{
		fmt.Println(key, "-", value)
	}

	phonebook := map[int]string{
		2343545446: "john",
		3465767998: "ada",
		1234567678: "mario",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[1234567678])

	phonebook[3465767998] = "bowser"
	fmt.Println(phonebook)
	fmt.Println(len(phonebook))

	mp := map[string]int {
		"apple": 7,
		"orange": 8,
		"pear": 10,
	}

	_, ok := score["kate"]
	if  ok {
		fmt.Println("yes")
		// fmt.Println(v)
	}

	val, ok := mp["apple"]
	fmt.Println(val, ok)
	fmt.Println(mp)
	fmt.Println(len(mp))

	if v, ok := mp["app"]; ok {
		fmt.Println("true")
		fmt.Println(v)
	} else {
		fmt.Println("false")
	}

	mp["mango"] = 9
	fmt.Println(mp)

	if v, ok := mp["mango"]; ok {
		fmt.Println("The value is =>", v)
		delete(mp, "mango")
	}
	fmt.Println(mp)

	language := make(map[string]string)
	language["JS"] = "JavaScript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("list of languages:", language)
	fmt.Println("JS short for:", language["JS"])

	delete(language, "RB")
	fmt.Println(language)

	for k, v := range language {
		fmt.Printf("For key %v, value is %v\n", k, v)
	}

}