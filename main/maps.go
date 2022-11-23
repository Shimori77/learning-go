package main

import "fmt"

func printMaps() {
	fmt.Println("--- Maps ---")

	m := map[string]int{"foo": 23, "bar": 98}
	fmt.Println(m)
	fmt.Println(m["bar"])

	m["foo"] = 54
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
}