package main

import "fmt"

// MyFilter - any value in the input slice that is less than or equal to 10.0 is retained,
// and all other values are filtered out.
func MyFilter(input []float64, f func(num float64) bool) []float64 {
	var result []float64
	for _, value := range input {
		if f(value) == true {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	input := []float64{17.3, 11.1, 9.9, 4.3, 12.6}
	res := MyFilter(input, func(num float64) bool {
		if num <= 10.0 {
			return true
		}
		return false
	})
	fmt.Println(res)
}

/* Output
[9.9 4.3]
*/
