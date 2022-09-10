package main

import "fmt"

// MyMap - Produces an output slice containing the square of the integers contained in the input slice.
// After declaring result to be a slice of len(slice) integers, it iterates over the range of values in input,
// transforming each value based on the function f passed in to MyMap.
func MyMap(input []int, f func(int) int) []int {
	result := make([]int, len(input))
	for index, value := range input {
		result[index] = f(value)
	}
	return result
}

func main() {
	slice := []int{10, 50, 22, 7, 48, 66}
	result := MyMap(slice, func(i int) int {
		return i * i
	})
	fmt.Println(result)
}

/* Output
[100 2500 484 49 2304 4356]
*/
