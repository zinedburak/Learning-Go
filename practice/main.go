package main

import (
	"fmt"
	"math"
)

func main() {
	var aString string = "Hello World"
	fmt.Println("Oh shit here we go again  :) ")
	fmt.Println(aString)
	// To print the variable type you can use the print with format function of fmt which is Printf and use %T for the type
	fmt.Printf("The variable's type is %T \n", aString)
	// Another way of declaring variables
	var anotherString = "Hello world"
	fmt.Printf("The variable's type is %T \n", anotherString)

	// Another way of declaring variables. This type of declaration only works inside of functions and they can not be const

	differentString := "Hello Burak"
	fmt.Printf("The variable's type is %T \n", differentString)

	// Declaring multiple variables at the same line
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum: ", intSum)

	// Addition is done in binary and the precision of adding two floating numbers can not be fully trusted
	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum: ", floatSum)
	// This method should be used if you are adding two floats
	floatSum = math.Round(floatSum*100) / 100
	fmt.Println(floatSum)

}
