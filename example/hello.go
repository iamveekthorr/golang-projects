package main

import (
	"fmt"
)

func main() {

	// variables and types
	var name string = "victor"
	var age int = 24
	isMarried := false

	fmt.Println(age, name, isMarried)

	// Functions
	var sum int = addTwoNumbers(2, 3)
	fmt.Println("Sum is: ", sum)

	var product, message = multiplyTwoNumbersAndReturnString(2, 3)
	fmt.Println(message, product)
}

// if a function should return a datatype,
// there is a need to name the type being returned
func addTwoNumbers(a, b int) int {
	return a + b
}

// if a function should return multiple datatypes,
// there is a need to name the type being returned
func multiplyTwoNumbersAndReturnString(a, b int) (int, string) {
	return a + b, "Product is: "
}
