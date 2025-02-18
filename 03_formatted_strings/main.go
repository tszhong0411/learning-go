package main

import "fmt"

func main() {
	age := 35
	name := "mario"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello learners!")
	fmt.Println("goodbye learners!")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	// %q: a double-quoted string
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	// %T: shows the type of an item
	fmt.Printf("age is of type %T \n", age)
	// %f: decimal point
	fmt.Printf("you scored %0.2f points! \n", 225.55)

	// Sprintf (save formatted strings)
	str := fmt.Sprintf("my age is %v and my name is %v", age, name)
	fmt.Println(str)
}
