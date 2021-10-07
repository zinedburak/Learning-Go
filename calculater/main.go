package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

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

	fmt.Printf("The sum of %v and %v is %v", val, val1, (val + val1))

}
