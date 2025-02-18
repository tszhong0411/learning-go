package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	// looping maps
	for k, v := range menu {
		fmt.Printf("key: %v, value: %v \n", k, v)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[984759373])

	phonebook[984759373] = "bowser"
	fmt.Println(phonebook)

	phonebook[845775485] = "yoshi"
	fmt.Println(phonebook)
}
