package main

import "fmt"

func updateName(x string) {
	x = "wedge"
}

func updateNameWithPtr(x *string) {
	*x = "dereferenced"
}

func main() {
	name := "tifa"
	updateName(name)

	// memory address
	fmt.Println(&name)

	namePtr := &name

	// dereference
	fmt.Println(*namePtr)

	updateNameWithPtr(namePtr)

	fmt.Println(name)
}
