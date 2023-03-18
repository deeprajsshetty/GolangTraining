// Sample program to show how variables of an unnamed type can
// be assigned to variables of a named type, when they are identical.
package main

import "fmt"

// example represents a type with different fields.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// Declare a variable of an anonymous type and init using a struct literal.
	fmt.Println("**********Assign the value of unnamed struct type to the named struct type.*****************************************")
	e := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Create a value of type example.
	var ex example

	// Assign the value of unnamed struct type to the named struct type.
	ex = e

	// Display the values.
	fmt.Printf("%+v\n", ex)
	fmt.Printf("%+v\n", e)
	fmt.Println("Flag", e.flag)
	fmt.Println("Counter", e.counter)
	fmt.Println("Pi", e.pi)
}

/* Output
**********Assign the value of unnamed struct type to the named struct type.*****************************************
{flag:true counter:10 pi:3.141592}
{flag:true counter:10 pi:3.141592}
Flag true
Counter 10
Pi 3.141592
*/
