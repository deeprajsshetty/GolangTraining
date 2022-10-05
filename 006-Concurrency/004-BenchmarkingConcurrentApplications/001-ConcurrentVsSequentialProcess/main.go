package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	In this example we are comparing the time required to construct and sum 100 million floating-point
	numbers into a slice with and without concurrency.
*/

var (
	output1, output2, output3, output4 float64
	wg                                 sync.WaitGroup
)

// worker1 - appends a float64 value to construct an output slice of 100 million values while computing the sum in the slice
func worker1() {
	defer wg.Done()
	var output []float64
	sum := 0.0
	for index := 0; index < 100_000_000; index++ {
		output = append(output, 89.6)
		sum += 89.6
	}
	output1 = sum
}

// worker2 - appends a float64 value to construct an output slice of 100 million values while computing the sum in the slice
func worker2() {
	defer wg.Done()
	var output []float64
	sum := 0.0
	for index := 0; index < 100_000_000; index++ {
		output = append(output, 64.8)
		sum += 64.8
	}
	output1 = sum
}

// worker3 - appends a float64 value to construct an output slice of 100 million values while computing the sum in the slice
func worker3() {
	defer wg.Done()
	var output []float64
	sum := 0.0
	for index := 0; index < 100_000_000; index++ {
		output = append(output, 956.8)
		sum += 956.8
	}
	output1 = sum
}

// worker4 - appends a float64 value to construct an output slice of 100 million values while computing the sum in the slice
func worker4() {
	defer wg.Done()
	var output []float64
	sum := 0.0
	for index := 0; index < 100_000_000; index++ {
		output = append(output, 1235.8)
		sum += 1235.8
	}
	output4 = sum
}

/*
we compute and output the computation time if the worker functions are executed sequentially. Then we execute
the four worker functions concurrently using goroutines. We compare the computation time and verify the correctness
of the results by outputting the sums with and without concurrency.

The computation time running the four worker functions concurrently is 67 percent the time running the four functions
sequentially. As expected, the sums computed are the same.
*/
func main() {
	wg.Add(8)
	// Compute time with no concurrent processing
	start := time.Now()
	worker1()
	worker2()
	worker3()
	worker4()
	elapsed := time.Since(start)
	fmt.Println("\nTime for 4 workers in series: ", elapsed)
	fmt.Printf("Output1: %f \nOutput2: %f \nOutput3: %f \nOutput4: %f\n", output1, output2, output3, output4)

	// Compute time with concurrent processing
	start = time.Now()
	go worker1()
	go worker2()
	go worker3()
	go worker4()
	elapsed = time.Since(start)
	fmt.Println("\nTime for 4 workers in parallel: ", elapsed)
	fmt.Printf("Output1: %f \nOutput2: %f  \nOutput3: %f  \nOutput4: %f", output1, output2, output3, output4)
}

/* Output
Time for 4 workers in series:  3.9372479s
Output1: 95680000176.244049
Output2: 0.000000
Output3: 0.000000
Output4: 123580000205.352280

Time for 4 workers in parallel:  0s
Output1: 95680000176.244049
Output2: 0.000000
Output3: 0.000000
Output4: 123580000205.352280
*/
