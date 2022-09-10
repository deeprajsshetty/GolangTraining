package main

import (
	"fmt"
	"time"
)

/*
If an attempt is made to send information to the channel when it is specified as a consumer,
or if an attempt is made to access information from the channel in the case that it is specified as a generator,
a compiler error will occur.
*/

var quit chan bool

// pingGenerator
func pingGenerator(c chan<- string) {
	// The channel can only be sent to - a generator
	for i := 0; i < 5; i++ {
		c <- "ping"
	}
}

// pongGenerator
func pongGenerator(c chan<- string) {
	// Information can only be sent to the channel - a generator
	for i := 0; i < 5; i++ {
		c <- "pong"
	}
}

// We have moved the one-second time delay into the output goroutine. This allows the ping and pong generators
// to alternate since each assignment to the channel blocks alternately until the channel is read by the consuming output goroutine.
func output(c <-chan string) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case value := <-c:
			fmt.Println(value)
		case <-time.After(3 * time.Second):
			fmt.Println("Program timed out.")
			quit <- true
		}
	}
}

func main() {
	quit = make(chan bool)
	c := make(chan string)
	go pingGenerator(c)
	go pongGenerator(c)
	go output(c)
	<-quit
}

/* Output
ping
pong
ping
pong
ping
pong
ping
pong
ping
pong
Program timed out.
*/
