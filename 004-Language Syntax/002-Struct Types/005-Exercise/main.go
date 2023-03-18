// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import "fmt"

// user represents a user in the system.
type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	fmt.Println("********Declare variable of type user and init using a struct literal.***********************************************")
	deep := user{
		name:  "deep",
		email: "deep@gmail.com",
		age:   36,
	}

	// Display the field values.
	fmt.Println("Name", deep.name)
	fmt.Println("Email", deep.email)
	fmt.Println("Age", deep.age)

	// Declare a variable using an anonymous struct.
	fmt.Println("********Declare a variable using an anonymous struct.****************************************************************")
	ed := struct {
		name  string
		email string
		age   int
	}{
		name:  "ed",
		email: "ed@gmail.com",
		age:   40,
	}

	// Display the field values.
	fmt.Println("Name", ed.name)
	fmt.Println("Email", ed.email)
	fmt.Println("Age", ed.age)

}

/* Output
********Declare variable of type user and init using a struct literal.***********************************************
Name deep
Email deep@gmail.com
Age 36
********Declare a variable using an anonymous struct.****************************************************************
Name ed
Email ed@gmail.com
Age 40
*/
