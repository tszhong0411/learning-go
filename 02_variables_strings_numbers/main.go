package main

import "fmt"

// Can't use short declaration outside of a function
// some := "hello"

func main() {
	// strings
	var one string = "mario"
	var two = "luigi"
	var three string

	fmt.Println(one, two, three)

	one = "peach"
	three = "bowser"

	fmt.Println(one, two, three)

	four := "yoshi"

	fmt.Println(four)

	// integers
	var age int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(age, ageTwo, ageThree)

	// bits & memory
	var num int8 = 25
	var numTwo int16 = -256
	var numThree uint8 = 255

	fmt.Println(num, numTwo, numThree)

	// floats
	var score float32 = -3.123
	var scoreTwo float64 = 12091283901345678983.123

	// inferred as float64
	scoreThree := 1.5

	fmt.Println(score, scoreTwo, scoreThree)
}
