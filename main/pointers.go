package main

import "fmt"

func printPointers() {
	fmt.Println("--- Pointers --- ")

	var firstName = new(string)

	*firstName = "Lucas"

	fmt.Println(firstName)
	fmt.Println(*firstName)
}
