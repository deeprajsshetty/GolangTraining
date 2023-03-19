// Sample program to show the basic concept of using a pointer to share data.
package main

import "fmt"

// user represents a user in the system.
type user struct {
	name   string
	email  string
	logins int
}

func main() {
	// Declare and initialize a variable named deep of type user.
	// We don't need to include all the fields when specifying field names with a struct literal.
	fmt.Println("*****Declare and initialize a variable named deep of type user and pass the \"address of\" the deep value using display()*******")
	deep := user{
		name:  "deep",
		email: "deep@gmail.com",
	}

	// Pass the "address of" the deep value.
	display(&deep)

	// Pass the "address of" the logins field from within the deep value.
	fmt.Println("*****Pass the \"address of\" the logins field from within the deep value and Prints &logins, logins and *logins*****************")
	increment(&deep.logins)

	// Pass the "address of" the bill value.
	fmt.Println("*****Display the refreshed value from deep type of user. Normally incremented values will be kept in heap.**********************")
	display(&deep)
}

// increment declares logins as a pointer variable whose value is always an address and points to values of type int.
func increment(logins *int) {
	*logins++
	fmt.Printf("&logins[%p] logins[%p] *logins[%d]\n\n", &logins, logins, *logins)
}

// display declares u as user pointer variable whose value is always an address and points to values of type user.
func display(u *user) {
	fmt.Println("******display() declares u as user pointer variable whose value is always an address and points to values of type user.*********")
	fmt.Printf("%p\t%+v\n", u, *u)
	fmt.Printf("Name: %q Email: %q Logins: %d\n\n", u.name, u.email, u.logins)
}

/* Output
*****Declare and initialize a variable named deep of type user and pass the "address of" the deep value using display()*******
******display() declares u as user pointer variable whose value is always an address and points to values of type user.*********
0xc00007c4b0    {name:deep email:deep@gmail.com logins:0}
Name: "deep" Email: "deep@gmail.com" Logins: 0

*****Pass the "address of" the logins field from within the deep value and Prints &logins, logins and *logins*****************
&logins[0xc00000a030] logins[0xc00007c4d0] *logins[1]

*****Display the refreshed value from deep type of user. Normally incremented values will be kept in heap.**********************
******display() declares u as user pointer variable whose value is always an address and points to values of type user.*********
0xc00007c4b0    {name:deep email:deep@gmail.com logins:1}
Name: "deep" Email: "deep@gmail.com" Logins: 1
*/
