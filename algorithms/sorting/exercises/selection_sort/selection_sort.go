package main

import "fmt"

// O(n^2)
func SelectionSort(array []int) {
	length := len(array)

	for i := 0; i < length; i++ {
		// Set the current index as minimum
		min := i
		temp := array[i]
		for j := i + 1; j < length; j++ {
			//Update minimum if current is lower than what we had previously
			if array[j] < array[min] {
				min = j
			}
		}
		array[i] = array[min]
		array[min] = temp
	}
	fmt.Print(array)
}

func main() {
	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	SelectionSort(numbers)
}
