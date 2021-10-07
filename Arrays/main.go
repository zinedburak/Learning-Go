package main

import (
	"fmt"
)

func main() {
	// A way of declaring and populating an array
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	// Printing all the elements and the first element of the colors array
	fmt.Println(colors)
	fmt.Println(colors[0])

	// Another way of declaring and populating an array
	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)

	// Arrays are only good for storing different elements
}
