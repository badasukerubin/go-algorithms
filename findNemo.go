package main

import (
	"fmt"
)

func factoryNemo(numnerOfNemos int) []string {
	s := make([]string, numnerOfNemos)
	for i := range s {
		s[i] = "Nemo"
	}

	return s
}

/**
1. Linear Time O(n)
It takes n time to find Nemo depending on the number of elements in the array
*/
// nemo := [1]string{"Nemo"}
// friendsOfNemo := [10]string{"Bruce", "Jef", "Bird", "Nemo", "Finder", "Nemo's Brother", "Kingston", "Still Typing", "Still Finding", "Not Ten Yet?"}
var nemos = factoryNemo(100000)

func findNemo() {
	for i := 0; i < len(nemos); i++ {
		if nemos[i] == "Nemo" {
			fmt.Println("Found Nemo")
		}
	}
}

func optimizeFindNemo() {
	for i := 0; i < len(nemos); i++ {
		if nemos[i] == "Nemo" {
			fmt.Println("Found Nemo")
			break
		}
	}
	// Sadly, we care about the worst case scenarios i.e. Nemo is the last element in the array and our array still gets to run O(n) times. Big Oh's a pain.
}

/**
Constant Time O(1)
It takes exactly 1 operation hence 1 time to run this function. On all processors, this function will run only once
*/
func screamNemo() {
	fmt.Printf("Nemo!, Nemo!!, Nemo!!!")
}

func main() {
	optimizeFindNemo()
	screamNemo()
}
