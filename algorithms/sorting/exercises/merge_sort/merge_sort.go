package main

import "fmt"

// var result = make([]int, 0)

func merge(left []int, right []int) []int {
	leftIndex := 0
	rightIndex := 0

	result := []int{}

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	for ; leftIndex < len(left); leftIndex++ {
		result = append(result, left[leftIndex])
	}

	for ; rightIndex < len(right); rightIndex++ {
		result = append(result, right[rightIndex])
	}

	return result
}

// Olog(n)
func MergeSort(array []int) []int {
	length := len(array)

	if length == 1 {
		return array
	}

	middle := length / 2
	left := array[:middle]
	right := array[middle:]

	// fmt.Print(merge(MergeSort(left), MergeSort(right)))
	// return left
	return merge(MergeSort(left), MergeSort(right))
}

func main() {
	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	mergeSort := MergeSort(numbers)
	fmt.Print(mergeSort)
}
