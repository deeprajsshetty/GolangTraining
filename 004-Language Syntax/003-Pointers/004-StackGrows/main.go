// Sample program to show how stacks grow/change.
package main

import "fmt"

// Number of elements to grow each stack frame.
// Run with 1 and then with 1024.
const size = 1024

// main is the entry point for the application.
func main() {
	s := "HELLO"
	fmt.Println("******Example to demonstrate, how stack grows*************************************************************************")
	stackCopy(&s, 0, [size]int{})
}

// stackCopy recursively runs increasing the size of the stack.
// go:noinline
func stackCopy(s *string, c int, a [size]int) {
	fmt.Println(c, s, *s)

	c++
	if c == 100 {
		return
	}

	stackCopy(s, c, a)
}

/* Output - Run through Golang Playgroound
******Example to demonstrate, how stack grows*************************************************************************
0 0xc00007ff60 HELLO
1 0xc00007ff60 HELLO
2 0xc00010ff60 HELLO
3 0xc00010ff60 HELLO
4 0xc00010ff60 HELLO
5 0xc00010ff60 HELLO
6 0xc00019ff60 HELLO
7 0xc00019ff60 HELLO
8 0xc00019ff60 HELLO
9 0xc00019ff60 HELLO
*/
