package main

import (
	"fmt"
	"time"
)

// We can use a quit channel to block main from exiting and avoid the use of WaitGroup
var quit chan bool

// pingGenerator
func pingGenerator(c chan string) {
	for i := 0; i < 5; i++ {
		c <- "ping"
		time.Sleep(time.Second * 1)
	}
}

// output
func output(c chan string) {
	for {
		select {
		case value := <-c:
			fmt.Println(value)
		case <-time.After(3 * time.Second):
			fmt.Println("Program timed out.")
			quit <- true
		}
	}
}

/*
The quit channel is initialized as the first line of code in main. The last line of code in main, <- quit,
blocks main from ending until a Boolean value is assigned to quit. This occurs in the second case statement in goroutine output.
This mechanism for controlling the end of the program is simpler and less encumbered than using WaitGroup.
*/
func main() {
	quit = make(chan bool)
	c := make(chan string)
	go pingGenerator(c)
	go output(c)
	<-quit
}

/* Output
ping
ping
ping
ping
ping
Program timed out.
*/
