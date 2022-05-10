package furtherpractice

import (
	"fmt"
	"sort"
)

// Write a function:

// func Solution(A []int) int

// that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

// For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

// Given A = [1, 2, 3], the function should return 4.

// Given A = [−1, −3], the function should return 1.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [1..100,000];
// each element of array A is an integer within the range [−1,000,000..1,000,000].

func Solution() {
	A := []int{1, 3, 6, 4, 1, 2}
	// Assign the lowest possible (+ve) value to result
	result := 1

	// Sort A asc
	sort.Ints(A)

	// Iterate over A
	for i := 0; i < len(A); i++ {
		// Check if the consecutive min values = result and increment if true
		if A[i] == result {
			result++
		}
	}

	fmt.Println(result)
}
