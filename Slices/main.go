package main

import "fmt"

func main() {
	var colors_array = [3]string{"Red", "Green", "Blue"}
	var colors_slice = []string{"Red", "Green", "Blue"}

	fmt.Println(colors_array)
	fmt.Println(colors_slice)

	// Append new item to slice
	colors_slice = append(colors_slice, "Purple")

	fmt.Println(colors_slice)

	// delete
}
