package main

import "fmt"

var counter int

func inception() string {
	// Base case
	if counter > 3 {
		fmt.Println(counter)
		fmt.Print("done!")
		return "done"
	}
	counter++
	// Recursive case
	return inception()
}

func main() {
	inception()
}
