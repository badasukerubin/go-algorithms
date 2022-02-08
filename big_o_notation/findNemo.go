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
func dropConstantsFindNemo() {
	var groupOfBystandingFishes = make([]string, 10) // O(1)
	groupOfBystandingFishes[0] = "Jack"              // O(1)
	groupOfBystandingFishes[1] = "Son"               // O(1)
	groupOfBystandingFishes[2] = "Mich"              // O(1)
	groupOfBystandingFishes[3] = "Ael"               // O(1)

	for i := 0; i < len(nemos); i++ {
		if nemos[i] == "Nemo" {
			fmt.Println("Found Nemo") // O(n)
			break
		}
	}

	fmt.Printf("The group of bystanding fishes we met on the way are %s %s %s and %s", groupOfBystandingFishes[0], groupOfBystandingFishes[1], groupOfBystandingFishes[2], groupOfBystandingFishes[3]) // O(1)
	// The Big Oh of this function should be O(6 + n) but 6 operations here doesn't matter when the funtion could likely do a thousand! 6 is a constant and has the likelihood of becoming infinitely smaller on a time scale hence should be removed.
	// The Big Oh noe becomes O(n)
}

/**
Constant Time O(1)
It takes exactly 1 operation hence 1 time to run this function. On all processors, this function will run only once
*/
func screamNemo() {
	fmt.Printf("Nemo!, Nemo!!, Nemo!!!")
}

func main() {
	dropConstantsFindNemo()
	screamNemo()
}
