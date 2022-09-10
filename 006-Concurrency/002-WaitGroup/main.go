package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
This program does nothing useful except illustrating the WaitGroup construct and
shows three goroutines running concurrently.

A global variable wg of type sync.WaitGroup is declared.

In main, we invoke wg.Add(3). The last line of code in main is wg .Wait().
This causes main to pause until the value in wg is zero.
This assures us that all three goroutines complete their work before the program terminates.

In each of the goroutines, the first line of code, defer wg.Done(), causes the
value of the global variable wg to be decremented when the goroutine completes
its work. When wg reaches a value of zero, the function main is allowed to exit.
*/

// A global variable wg of type sync.WaitGroup is declared.
var wg sync.WaitGroup

// outputStrings - Output the Slice of String
func outputStrings() {
	defer wg.Done()
	strings := [5]string{"One", "Two", "Three", "Four", "Five"}
	for i := 0; i < 5; i++ {
		delay := 1 + rand.Intn(3)
		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Println(strings[i])
	}
}

// outputInts - Output the Slice of Ints
func outputInts() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		delay := 1 + rand.Intn(3)
		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Println(i)
	}
}

// outputFloats - Output the Slice of Floats
func outputFloats() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		delay := 1 + rand.Intn(3)
		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Println(float64(i*i) + 0.5)
	}
}

func main() {
	wg.Add(3)
	go outputStrings()
	go outputInts()
	go outputFloats()
	wg.Wait()
}

// The sequence of random numbers generated is the same each time the program is run,
// but the output sequence varies from run to run. This is because the time multiplexer
// allocates different chunks of execution time to each concurrent goroutine differently
// each time the program runs
// IF we run this Program Multiple time Output Sequence will vary
/* Output
0.5
0
One
Two
1.5
1
Three
4.5
2
Four
Five
3
9.5
16.5
4
*/
