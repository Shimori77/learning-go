package main

import "fmt"

func printArrays() {
	fmt.Println("--- Arrays and Slices ---")
	var arr [3]int
	arr[0] = 3
	arr[1] = 2
	arr[2] = 1

	fmt.Println(arr)

	arr0 := [3]int{4, 5, 6}
	fmt.Println(arr0)

}
