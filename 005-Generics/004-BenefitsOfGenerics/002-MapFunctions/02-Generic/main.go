package main

import "fmt"

// MyMap - Produces an output slice containing the square of the integers contained in the input slice.
// After declaring result to be a slice of len(slice) integers, it iterates over the range of values in input,
// transforming each value based on the function f passed in to MyMap.
/*func MyMap(input []int, f func(int) int) []int {
	result := make([]int, len(input))
	for index, value := range input {
		result[index] = f(value)
	}
	return result
}*/

// GenericMap - MyMap function handled using Generic Way.
/*
Function GenericMap takes two generic parameters, T1 and T2.
Using the function f that is passed in, it transforms the data in the input slice to type T2.
Here, T1 and T2 are not constrained. They are of type any.
*/
func GenericMap[T1, T2 any](input []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(input))
	for index, value := range input {
		result[index] = f(value)
	}
	return result
}

func main() {
	slice := []int{10, 50, 22, 7, 48, 66}
	result := GenericMap(slice, func(i int) int {
		return i * i
	})
	fmt.Println(result)
}

/* Output
[100 2500 484 49 2304 4356]
*/
