package main

/*
By summing half the numbers in each of two goroutines, a substantial improvement in execution
time occurs as is evident in the program output.
*/

import (
	"fmt"
	"time"
)

const (
	NumbersToSum = 100_000_00
)

func sum(s []float64, c chan<- float64) {
	// A generator that puts data into channel
	sum := 0.0
	for _, v := range s {
		sum += float64(v)
	}

	c <- sum // blocks until c is taken out of channnel
}

func plainSum(s []float64) float64 {
	sum := 0.0
	for _, v := range s {
		sum += float64(v)
	}
	return sum
}

func main() {
	s := []float64{}
	for i := 0; i < NumbersToSum; i++ {
		s = append(s, 1.0)
	}
	c := make(chan float64)
	// The two goroutines perform their computation in a for-loop concurrently and return
	//their values by assigning to the channel variable c
	start := time.Now()
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// the assignment of the two sums to first and second is blocked until both values are assigned to the channel.
	first, second := <-c, <-c // receive from each c
	elapsed := time.Since(start)
	fmt.Printf("first: %f  second: %f elapsed time: %v", first, second,
		elapsed)
	start = time.Now()
	answer := plainSum(s)
	elapsed = time.Since(start)
	fmt.Printf("\nplain sum: %f elapsed time: %v", answer, elapsed)
}

/*Output
first: 5000000.000000  second: 5000000.000000 elapsed time: 8.7998ms
plain sum: 10000000.000000 elapsed time: 16.8447ms
*/
