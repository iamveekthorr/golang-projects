// Package main provides - Entry point to the application
package main

import (
	"fmt"
)

/**
* Problem Statement ->
* Consider the problem of adding two n-bit binary integers,
* sorted in two n-element arrays A and B.
* The sum of the two integers should be sorted in binary form in an (n + 1)-element arrays C.
* State the problem formally and write pseudocode for adding the two integer
* **/

func main() {
	arr := []int{0, 1, 0, 1}
	arr2 := []int{1, 0, 1, 0}
	result := addTwoBinaryIntegers(arr, arr2)
	fmt.Println(result)
}

// To prove correctness of this algorithm, it abide by 3 Rules!
// 1. Initialization
// 2. Maintenance
// 3. Termination
func addTwoBinaryIntegers(a, b []int) []int {
	var c []int
	carry := 0

	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := a[i] + b[j] + carry

		if sum >= 2 {
			carry = 1
		} else {
			carry = 0
		}

		c = append(c, sum%2)

		// Move to the next element in the array
		i--
		j--
	}

	return c
}
