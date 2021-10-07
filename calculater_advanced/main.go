package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readValue() (float64, float64, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value 1 = ")
	text, _ := reader.ReadString('\n')
	val, err := strconv.ParseFloat(strings.TrimSpace(text), 64)

	if err != nil {
		panic(err)
	}

	fmt.Print("Value 2 = ")
	text, _ = reader.ReadString('\n')
	val1, err1 := strconv.ParseFloat(strings.TrimSpace(text), 64)

	if err1 != nil {
		panic(err1)
	}

	fmt.Print("Select an operation  (+, *, -, /) ")
	text, _ = reader.ReadString('\n')
	operation := strings.TrimSpace(text)

	return val, val1, operation
}

func add(value1 float64, value2 float64) float64 {
	return math.Round((value1 + value2))
}
func multiply(value1 float64, value2 float64) float64 {
	return math.Round((value1 * value2))
}
func divide(value1 float64, value2 float64) float64 {
	return math.Round((value1 / value2))
}
func subtract(value1 float64, value2 float64) float64 {
	return math.Round((value1 - value2))
}
func main() {
	value1, value2, operation := readValue()
	var result float64
	switch operation {
	case "*":
		result = multiply(value1, value2)
	case "+":
		result = add(value1, value2)
	case "-":
		result = subtract(value1, value2)
	case "/":
		result = divide(value1, value2)
	}
	fmt.Println("The result is : ", result)

}
