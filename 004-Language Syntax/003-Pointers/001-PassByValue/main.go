// Sample program to show the basic concept of pass by value.
package main

import "fmt"

func main() {
	// Declare variable of type int with a value of 10
	fmt.Println("*******Declare variable of type int with a value of 10. Display the \"value of\" and \"address of\" count.**********")
	count := 10

	// Display the "value of" and "address of" count.
	fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" the count to "increment" function.
	fmt.Println("*******Pass the \"value of\" the count to \"increment\" function. Will be showing diffrent Addr as it is copy*******")
	increment(count)

	fmt.Println("******Same value and addr will be printed as it is pass by value.***************************************************")
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

// increment declares count as a pointer variable whose value is
// always an address and points to values of type int.
//go:noinline
func increment(inc int) {

	// Increment the "value of" inc.
	inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

/* Output
*******Declare variable of type int with a value of 10. Display the "value of" and "address of" count.**********
count:  Value Of[ 10 ]  Addr Of[ 0xc00001a0b8 ]
*******Pass the "value of" the count to "increment" function. Will be showing diffrent Addr as it is copy*******
inc:    Value Of[ 11 ]  Addr Of[ 0xc000115ec0 ]
******Same value and addr will be printed as it is pass by value.***************************************************
count:  Value Of[ 10 ]  Addr Of[ 0xc00001a0b8 ]
*/
