package main

import "fmt"

func main() {
	states := make(map[string]string)

	fmt.Println(states)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)
	// this will print different results every time it runs  since they are not added one by one or in order
	for k, v := range states {
		fmt.Printf("%v: %v \n", k, v)
	}
}
