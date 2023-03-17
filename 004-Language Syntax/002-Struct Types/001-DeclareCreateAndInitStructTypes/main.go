// Sample program to show how to declare and initialize struct types.
package main

import "fmt"

// example represents a type with different fields.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// Declare a variable of type example set to its zero value.
	fmt.Println("******Declare a variable of type example set to its zero value.*************************************************************")
	var e1 example

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and init using a struct literal.
	fmt.Println("******Declare a variable of type example and init using a struct literal.***************************************************")
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("pi", e2.pi)
}

/* Output
******Declare a variable of type example set to its zero value.*************************************************************
{flag:false counter:0 pi:0}
******Declare a variable of type example and init using a struct literal.***************************************************
Flag true
Counter 10
pi 3.141592
*/
