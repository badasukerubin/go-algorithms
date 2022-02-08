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
1. Linear Time O(n)
It takes n time to find Nemo depending on the number of elements in the array
*/
func findNemo() {
	for i := 0; i < len(nemos); i++ {
		if nemos[i] == "Nemo" {
			fmt.Println("Found Nemo")
		}
	}
}

// 1.1 Rule 1:: Worst Csae Scenario: We are interested, more like worried about the worst case scenario
func optimizeFindNemo() {
	for i := 0; i < len(nemos); i++ {
		if nemos[i] == "Nemo" {
			fmt.Println("Found Nemo")
			break
		}
	}
	// Sadly, we care about the worst case scenarios i.e. We are wary that Nemo could be the last element in the array and our array still gets to run O(n) times. Big Oh's a pain.
}

// 1.2 Rule 2:: Drop constants: We drop constants because mathematically, they barely influence time complexities linearly
func dropConstantsAndFindNemo() {
	var groupOfBystandingFishes = make([]string, 10) // O(1)
	groupOfBystandingFishes[0] = "Jack"              // O(1)
	groupOfBystandingFishes[1] = "Son"               // O(1)
	groupOfBystandingFishes[2] = "Mich"              // O(1)
	groupOfBystandingFishes[3] = "Ael"               // O(1)

	for i := 0; i < len(nemos); i++ {
		if nemos[i] == "Nemo" {
			fmt.Println("Found Nemo") // O(n)
		}
	}

	fmt.Printf("The group of bystanding fishes we met on the way are %s %s %s and %s", groupOfBystandingFishes[0], groupOfBystandingFishes[1], groupOfBystandingFishes[2], groupOfBystandingFishes[3]) // O(1)
	// The Big Oh of this function should be O(6 + n) but 6 operations here doesn't matter when the funtion could likely do a thousand! 6 is a constant and has the likelihood of becoming infinitely smaller on a time scale hence should be removed.
	// The Big Oh noe becomes O(n)
}

// 1.3 Rule 3:: Different terms for different inputs: due to the fact that different inputs will produce different outputs, it is wise to split the big ohs for each operations depending on these inputs into different terms.
func findNemoInDifferentPlaces() {
	for i := 0; i < len(nemos); i++ {
		if nemos[i] == "Nemo" {
			fmt.Println("Found Nemo in nemos")
		}
	}
	// Here, it could take 1000 operations to find Nemo, hence the big oh is O(a)

	for i := 0; i < len(friendsOfNemo); i++ {
		if nemos[i] == "Nemo" {
			fmt.Println("Found Nemo in friendsOfNemo")
		}
	}
	// Here, it could take 10 operations to find Nemo, hence the big oh is O(b)

	// The big oh here now becomes O(a + b), could be O(n + m) just different terms for different inputs.
}

func findNemoTwice() {
	for i := 0; i < len(friendsOfNemo); i++ {
		for j := 0; j < len(nemos); j++ {
			if nemos[j] == "Nemo" {
				fmt.Println("Found Nemo in nemos")
			}
		}
	}
	// Don't get confused but the big oh notation for this block is O(a * b). When it comes to nested loops or operations, we multiply the terms

	for i := 0; i < len(friendsOfNemo); i++ {
		for j := 0; j < len(friendsOfNemo); j++ {
			if nemos[j] == "Nemo" {
				fmt.Println("Found Nemo in nemos")
			}
		}
	}
	// This one is O(n * n) => O(n ^ 2), because the inputs are the same.
}

// 1.4 Rule 4:: Drop non dominant pairs: We drop non dominant pairs/term because mathematically, they barely influence time complexities linearly
func findNemoOnceThenTwice() {
	for j := 0; j < len(nemos); j++ { // O(n)
		if nemos[j] == "Nemo" {
			fmt.Println("Found Nemo in nemos")
		}
	}

	for i := 0; i < len(friendsOfNemo); i++ { // O(n ^ 2)
		for j := 0; j < len(friendsOfNemo); j++ {
			if nemos[j] == "Nemo" {
				fmt.Println("Found Nemo in nemos")
			}
		}
	}

	// When we have a big Oh with O(n) and another with O(n ^ 2), we drop the least pair which would be O(n). if it were O(n ^ 2) and O(n ^ 3), we drop O(n ^ 2).

	// The big oh of the above even though it should be O(n + n ^ 2) becomes O(n ^ 2) simply because O(n) really wouldn't matter in the scale of things considering how big O(n ^ 2) could get.
}

/**
2. Constant Time O(1)
It takes exactly 1 operation hence 1 time to run this function. On all processors, this function will run only once
*/
func screamNemo() {
	fmt.Printf("Nemo!, Nemo!!, Nemo!!!")
}

func main() {
	findNemoOnceThenTwice()
	screamNemo()
}
