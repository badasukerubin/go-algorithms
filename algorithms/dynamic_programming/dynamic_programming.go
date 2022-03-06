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
}
