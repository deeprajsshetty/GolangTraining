package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// pingGenerator - The pingGenerator goroutine assigns the string “ping” to the channel variable
// c every second and does this five times. The left arrow from the value “ping” to the variable
// c represents the assignment of the “ping” value to c.
// The channel must be empty for this assignment to be made. In the output goroutine , the assignment to value,
// using value := <- c, gobbles up the channel variable c as soon as it is assigned in the pingGenerator.
// This occurs every second. During the time between “ping” assignments from the pingGenerator, the value assignment is blocked.
func pingGenerator(c chan string) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		c <- "ping"
		time.Sleep(time.Second * 1)
	}
	close(c)
}

// output - Execution is halted within the output goroutine until there is information in the channel assigning to value.
// So the two goroutines are being affected by the channel variable c, common to both.

// There is a problem. When the pingGenerator has emitted its five “ping” assignments, each displayed on the console through
// the output goroutine, it blocks while waiting for a sixth channel assignment. This never occurs. The program crashes with
// the error message shown earlier. A deadlock has occurred. The program cannot terminate.
func output(c chan string) {
	defer wg.Done()
	for {
		value := <-c
		fmt.Println(value)
	}
}

func main() {
	// The first line of code in main initializes a channel, c. Channels must be initialized
	// with a make statement before they can be used.
	c := make(chan string)

	// As in the previous listing, we create a WaitGroup variable, wg, with the initial value of 2.
	wg.Add(2)

	// We next launch the two goroutines, pingGenerator and output, passing the channel variable c to each.
	go pingGenerator(c)
	go output(c)

	wg.Wait()
}

/* Output
ping
ping
ping
ping
ping
fatal error: all goroutines are asleep - deadlock!
*/
