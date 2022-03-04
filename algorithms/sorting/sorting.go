package main

import (
	"fmt"
	"sort"
)

// Go's built-in sort
func main() {
	letters := []string{"a", "d", "z", "e", "r", "b"}
	basket := []int{2, 65, 34, 2, 1, 7, 8}

	sort.Strings(letters)
	sort.Ints(basket)
	fmt.Println(letters)
	fmt.Println(basket)
}
