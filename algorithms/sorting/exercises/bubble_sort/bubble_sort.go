package main

import "fmt"

// O(n^2)
func BubbleSort(array []int) {
	length := len(array)
	key := 1

	for i := key; i < length; i++ {
		j := i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j--
		}
	}
	fmt.Print(array)
}

func main() {
	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	BubbleSort(numbers)
}
