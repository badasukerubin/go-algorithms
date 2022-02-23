package main

import "fmt"

/**
Write a function to reverse string(s)
Toluwani => inawuloT
I am engineer => reenigne ma I
**/
// [84 111 108 117 119 97 110 105]
func reverse(data string) {
	/**
	Simple implementation:
	data[0, len(data)-1] = data[len(data)-1, 0]
	...
	data[n, len(data)-1-n] = data[len(data)-1-n, n]
	*/
	dataToRune := []rune(data)            // Convert data of type string to rune O(1)
	lengthOfDataToRune := len(dataToRune) // Get length of rune O(1)
	midIndex := lengthOfDataToRune / 2    // Get the last index for the forward iteration O(1)

	for i := 0; i < lengthOfDataToRune; i++ { // O(n)
		if i < midIndex {
			dataToRune[i], dataToRune[lengthOfDataToRune-1-i] = dataToRune[lengthOfDataToRune-1-i], dataToRune[i]
		}
	}

	fmt.Print(string(dataToRune))
}

func cleanerReverse(data string) {
	dataToRune := []rune(data)            // O(1)
	lengthOfDataToRune := len(dataToRune) // O(1)

	for i, j := 0, lengthOfDataToRune-1; i < j; i, j = i+1, j-1 { // O(n)
		dataToRune[i], dataToRune[j] = dataToRune[j], dataToRune[i]
	}

	fmt.Print(string(dataToRune))
}

func main() {
	reverse("I am engineer")
	cleanerReverse("I am engineer")
}
