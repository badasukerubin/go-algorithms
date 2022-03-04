package main

import "fmt"

/**
Being lazy and too busy to write tests for this learning proc, badasukerubin shall move all exercises to individual directories from henceforth.
*/

// O(2^n)
func fibonacciRecursive(number int) int {
	if number < 2 {
		return number
	}

	return fibonacciRecursive(number-1) + fibonacciRecursive(number-2)
}

// O(n)
func fibonacciIterative(number int) int {
	array := []int{0, 1}

	for i := 2; i < number+1; i++ {
		array = append(array, array[i-1]+array[i-2])
	}

	return array[number]
}

func main() {
	fibonacciIterative := fibonacciIterative(8)
	fmt.Printf("Find Fibonacci iteratively: %v\n", fibonacciIterative)

	fibonacciRecursive := fibonacciRecursive(8)
	fmt.Printf("Find Fibonacci iteratively: %v\n", fibonacciRecursive)
}
