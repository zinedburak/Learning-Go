package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// This is how you create functions for types
// The object you pass in the type functions is the copy of the object not the originale object
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func main() {
	poodle := Dog{Breed: "poodle", Weight: 10, Sound: "Woof!"}

	poodle.Speak()

}
