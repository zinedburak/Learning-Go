package main

import (
	"fmt"
)

func main() {
	// Declaring a pointer

	var pointer *float64

	// A pointer that has only been initialized will have the nill value
	// Since it is not pointing at any thing
	fmt.Println("Value of p: ", pointer)

	anInt := 42.13
	pointer = &anInt
	// This will give us the memory address of the anInt variable
	fmt.Println("Memory address of anInt variable: ", pointer)

	// This will give us the value of the variable that is in the above address
	fmt.Println("True Value of the anInt variable ", *pointer)

	// You can change the value of the variable via pointers
	pointer1 := &anInt
	*pointer1 = *pointer1 / 31
	fmt.Println("The value of pointer 1: ", *pointer1) // 1.35
	fmt.Println("The value of pointer: ", *pointer)    // 1.35
	fmt.Println("The value of variable ", anInt)       // 1.35

}
