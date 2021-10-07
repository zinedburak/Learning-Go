package main

import "fmt"

// if you use upper case it is exported (public)
// if you use lower case it is not exported (private)
type Dog struct {
	Breed  string
	Weight int
}

func main() {
	poodle := Dog{Breed: "Poodle", Weight: 10}
	fmt.Println(poodle)
}
