package main

import (
	"fmt"
)

/**
Arrays in golang like C++ are static, this means that you define a static length.
Slices solve this by letting us define arrays without lengths.
**/

/**
Values can be added to an array
O(1) - constant time
*/
func push(array *[]string) {
	*array = append(*array, "second") // O(1)
}

/**
Values can be added to the first array (unshift)
O(n) - linear time
*/
func unshift(array *[]string) {
	var x = make([]string, len(*array)+1) // We define an array x with the length of our slice plus an extra length to cater for the incoming string we intend to push to the first index.
	x[0] = "first"

	for i := 0; i < len(*array); i++ { // O(n)
		x[i+1] = (*array)[i]
	}

	*array = x
}

/**
Values can be removed to an array
O(1) - constant time
*/
func pop(array *[]string) {
	*array = (*array)[:len(*array)-1] // We remove the last item then truncate the slice => O(1)
	// This method changes the order of the array
}

func main() {
	array := []string{} // We define a slice, length is mutable by the compiler
	push(&array)
	fmt.Printf("push: %v", array)
	unshift(&array)
	fmt.Printf("\nunshift: %v", array)
	pop(&array)
	fmt.Printf("\npop: %v", array)
}
