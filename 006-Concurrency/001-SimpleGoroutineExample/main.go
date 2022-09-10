package main

import (
	"fmt"
	"time"
)

// regularFunction
func regularFunction() {
	fmt.Println("Just entered regularFunction()")
	time.Sleep(10 * time.Second)
}

// goroutineFunction
func goroutineFunction() {
	fmt.Println("Just entered goroutineFunction()")
	time.Sleep(5 * time.Second)
	fmt.Println("goroutineFunction finished its work")
}

/*
When the goroutineFunction is launched as a goroutine using go goroutineFunction(),
it runs concurrently with the main function, which is a goroutine.
The first line of output occurs immediately even though the goroutineFunction requires
ten seconds to complete its work. When the regularFunction() is invoked next, five seconds elapses
before the line of output. “In main, one line below regularFunction()” is emitted.
Function main terminates immediately after this output is emitted, which ends the program before the goroutineFunction can complete its work. It gets interrupted and terminates when the program ends.
*/
func main() {
	go goroutineFunction()
	fmt.Println("In main one line below goroutineFunction()")
	regularFunction()
	fmt.Println("In main one line below regularFunction()")
}

/* Output
In main one line below goroutineFunction()
Just entered regularFunction()
Just entered goroutineFunction()
In main one line below regularFunction()
*/

// If we swap the time delays so that the goroutineFunction has a time delay of five seconds and
// the regularFunction has a time delay of ten seconds, the output becomes
/* Output
In main one line below goroutineFunction()
Just entered regularFunction()
Just entered goroutineFunction()
goroutineFunction finished its work
In main one line below regularFunction()
*/
