package main

import "fmt"

/**
Google Question
Given an array = [2,5,1,2,3,5,1,2,4]:
It should return 2

Given an array = [2,1,1,2,3,5,1,2,4]:
It should return 1

Given an array = [2,3,4,5]:
It should return 0
*/

/**
O(n)
*/
func naiveApproach(array []int) (bool, int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				return true, array[i]
			}
		}
	}
	return false, 0
}

/**
O(1)
Note: This doesn't sort orderly hance reuslt will differ with naiveApproach
*/
func optimzedApproach(array []int) (bool, int) {
	bucket := make(map[int]bool)

	for i := 0; i < len(array); i++ {
		if bucket[array[i]] {
			return true, array[i]
		}
		bucket[array[i]] = true
	}

	return false, 0
}

func main() {
	array := []int{2, 5, 1, 2, 3, 5, 1, 2, 4}
	// found, value := naiveApproach(array)
	found, value := optimzedApproach(array)
	fmt.Printf("%v %v", found, value)
}
