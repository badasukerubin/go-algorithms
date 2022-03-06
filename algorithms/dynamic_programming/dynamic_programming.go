package main

import (
	"fmt"
)

// var cache map[int]int

func addTo80(n int) int {
	fmt.Println("long time")
	return n + 80
}

func memoizedAddTo80() func(int) int {
	cache := make(map[int]int)

	return func(n int) int {
		if val, ok := cache[n]; ok {
			// runs for all remaining static values after one dynamic value
			return val
		} else {
			fmt.Println("long time - runs only once")
			cache[n] = n + 80
			return cache[n]
		}
	}

}

// O(2^n)
// From algorithms/recursion/exercises/fibonacci/fibonacci.go
var calculations int

func fibonacciRecursive(number int) int {
	calculations++
	if number < 2 {
		return number
	}

	return fibonacciRecursive(number-1) + fibonacciRecursive(number-2)
}

var calculationsMem int

func fibonacciMemoized(n int) int {
	cache := make(map[int]int)

	return fib(n, cache)
}

func fib(n int, cache map[int]int) int {
	calculationsMem++
	if val, ok := cache[n]; ok {
		return val
	} else {
		if n < 2 {
			return n
		} else {
			cache[n] = fib(n-1, cache) + fib(n-2, cache)
			return cache[n]
		}
	}
}

func main() {
	// Calling addTo80 multiple times consumes time for each process; especially terrible if it's a long process
	addTo80(5)
	addTo80(5)
	addTo80(5)

	// Calling memoizedAddTo80 multiple times only takes time for the first dynamic process then consecutive static processes per dynamic process; it is way efficient
	memoized := memoizedAddTo80()
	memoized(5)
	memoized(5)
	memoized(5)

	fibonacciRecursive(20)
	// From the result of calculations, we can see that it drastically increases in relation to the fibonacci number - try 100 💀. This is rather inefficient considering the fact that we could cache the results
	// fib(10) => 177 operations, f(20) => 21891 operations
	fmt.Println(calculations)

	_ = fibonacciMemoized(20)
	// fib(10) => 19 operation, f(20) => 39 operations 💀
	fmt.Println(calculationsMem)
}
