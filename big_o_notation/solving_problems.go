package main

import "fmt"

/**
Check two arrays for common elements. Return true if common element is found else false

Case 1: ['a', 'b', 'c', 'x'], ['z', 'y', 'i'] returns false
Case 2: ['a', 'b', 'c', 'x'], ['z', 'y', 'x'] returns true
*/

func naiveApproach(array1 []string, array2 []string) bool {
	// We iterate over array1, then over array2 and check if the values in array1 exist in array2
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[j] {
				fmt.Print(true)
				return true
			}
		}
	}
	fmt.Print(false)
	return false

	// The Big Oh here is O(a * b); you know, nested loops which is time consuming, we gotta optimize this.
}

func optimzedApproach(array1 []string, array2 []string) bool {
	m := make(map[string]bool)

	// O(a)
	// We iterate over array1, create a map off the array with structure {value: true} e.g. {b: true}
	for i := 0; i < len(array1); i++ {
		_, exist := m[array1[i]]
		if !exist {
			item := array1[i]
			m[item] = true
		}
	}
	// At the end of this iteration, m has the structure {a: true, b: true ...}

	// O(b)
	// We iterate over array2 and check if the values in array2 exist in the map
	for j := 0; j < len(array2); j++ {
		if m[array2[j]] {
			fmt.Print(true)
			return true
		}
	}
	fmt.Print(false)
	return false

	// The Big Oh here is O(a + b); this is more optimised.
}

func main() {
	var array1 = []string{"a", "b", "c", "g"}
	var array2 = []string{"z", "y", "x"}

	// naiveApproach(array1, array2)
	optimzedApproach(array1, array2)
}
