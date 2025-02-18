package main

import (
	"fmt"
	"strings"
)

func getInitialis(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fn, sn := getInitialis("tifa lockhart")
	fmt.Println(fn, sn)

	fn2, sn2 := getInitialis("cloud strife")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitialis("barret")
	fmt.Println(fn3, sn3)
}
