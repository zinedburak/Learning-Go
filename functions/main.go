package main

import "fmt"

func doSomething() {

	fmt.Println("Hello function")

}

func addValuses(value1 int, value2 int) int {
	return value1 + value2
}

func addAllValuses(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}

func main() {

	doSomething()
	sum := addValuses(10, 5)
	fmt.Println(sum)
	sum, lena := addAllValuses(10, 48, 23094, 234)
	fmt.Println(sum)
	fmt.Println(lena)

}
