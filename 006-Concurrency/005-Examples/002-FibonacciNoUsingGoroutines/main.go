package main

import "fmt"

func main() {
	// The first parameter, c, in function fibonacci puts information into the channel,
	// and the second parameter, quit, takes information out of the channel.
	c := make(chan int)
	quit := make(chan bool)

	// The goroutine is launched within main as an internal function. The Println(<-c) statement
	// blocks until the fibonacci function puts the value x into the integer channel c.
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}
		quit <- true
	}()
	fibonacci(c, quit)
}

// The select statement in function fibonacci either takes the next value of x into channel c
// or ends the program as soon as quit becomes true. The actual fibonacci sequence is computed
// as the second line within the case c <-x statement.
func fibonacci(c chan<- int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y // Generate the sequence
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}

/* Output
0
1
1
2
3
5
8
13
21
34
55
89
144
233
377
610
987
1597
2584
4181
quit
*/
