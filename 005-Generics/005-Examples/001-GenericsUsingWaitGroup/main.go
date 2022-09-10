package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// PrintOutput - Generic function capable for printing any type Slice instead of creating different functions.
func printOutput[T any](slice []T) {
	defer wg.Done()
	for _, value := range slice {
		fmt.Println(value)
	}
}

func main() {

	wg.Add(3)

	// Strings - Output the Slice of String
	go printOutput([]string{"One", "Two", "Three", "Four", "Five"})

	// ints - Output the Slice of Ints
	go printOutput([]int{5, 10, 78, 20, 45})

	// floats - Ouput the Slice of Floats
	go printOutput([]float32{23.7, 87.0, 43.8, 1.5, 2.8})

	wg.Wait()

}

/* Output
23.7
87
43.8
1.5
2.8
One
Two
Three
Four
Five
5
10
78
20
45
*/
