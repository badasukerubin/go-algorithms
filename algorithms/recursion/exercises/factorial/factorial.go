package main

import "fmt"

/**
Being lazy and too busy to write tests for this learning proc, badasukerubin shall move all exercises to individual directories from henceforth.
*/

// O(n)
func findFactorialRecursive(number int) int {
	if number <= 2 {
		return number
	}

	return number * findFactorialRecursive(number-1)
}

// O(n)
func findFactorialIterative(number int) int {
	answer := 1

	if number <= 0 {
		return number
	}

	for i := 1; i <= number; i++ {
		answer = answer * i
	}

	return answer
}

func main() {
	findFactorialIterative := findFactorialIterative(3)
	fmt.Printf("Find factorial iteratively: %v\n", findFactorialIterative)

	findFactorialRecursive := findFactorialRecursive(3)
	fmt.Printf("Find factorial iteratively: %v\n", findFactorialRecursive)
}
