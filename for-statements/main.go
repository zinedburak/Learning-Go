package main

import "fmt"

func main() {
	var colors_slice = []string{"Red", "Green", "Blue"}

	for i := 0; i < len(colors_slice); i++ {
		fmt.Println(colors_slice[i])
	}
	fmt.Println()
	for i := range colors_slice {
		fmt.Println(colors_slice[i])
	}
	fmt.Println()
	for _, color := range colors_slice {
		fmt.Println(color)
	}
	value := 1
	for value < 3 {
		fmt.Println("Value: ", value)
		value++
	}

}
