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
RESOURCES:

https://www.bigocheatsheet.com/
**/

// nemo := [1]string{"Nemo"}
var friendsOfNemo = [10]string{"Bruce", "Jef", "Bird", "Nemo", "Finder", "Nemo's Brother", "Kingston", "Still Typing", "Still Finding", "Not Ten Yet?"}
var nemos = factoryNemo(100000)

/**
1. Constant Space O(1)
It takes exactly 1 assignment hence 1 `space` to run this function.
*/
func findNemo() {
	for i := 0; i < len(nemos); i++ {
		if nemos[i] == "Nemo" {
			fmt.Println("Found Nemo")
		}
	}
	// Here, the space complexity to finding nemo is O(1). This is because we assign a value to the variable i only once throughout the entire runtime in i := 0
}

/**
2. Linear Space O(n)
It takes n assignments to find Nemo depending on the number of elements in the array
*/
func findAndAssignBadNemo() {
	for i := 0; i < len(nemos); i++ {
		nemos[i] = "Bad Nemo"
		if nemos[i] == "Nemo" {
			fmt.Println("Found Nemo")
		}
	}
	// Here, the space complexity to finding nemo is O(n). This is because we assign a value to the variable nemos[i] n times throughout the entire runtime in nemos[i] = "Bad Nemo"
}

func main() {
	findAndAssignBadNemo()
}
