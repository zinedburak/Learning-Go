package main

import "fmt"

func main() {

	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "Less than zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Print(result)

}
