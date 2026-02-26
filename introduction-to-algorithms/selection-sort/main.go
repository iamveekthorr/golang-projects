// Package main provides TODO: add description
package main

import "fmt"

func main() {
	a := []int{29, 10, 14, 37, 13}

	selectionSort(a)

}

func selectionSort(a []int) {
	length := len(a)
	for i := 0; i < length-1; i++ {
		minIndex := i

		for j := i + 1; j < length; j++ {
			if a[j] < a[minIndex] {
				minIndex = j // found smaller element
			}
		}

		// move element to the correct position
		temp := a[i]
		// swap elements
		a[i] = a[minIndex]
		a[minIndex] = temp
	}

	fmt.Printf("%v", a)
}
