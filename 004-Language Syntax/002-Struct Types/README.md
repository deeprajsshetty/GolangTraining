## Struct Types

Struct types are a way of creating complex types that group fields of data together. They are a great way of organizing and sharing the different aspects of the data your program consumes.

A computer architecture’s potential performance is determined predominantly by its word length (the number of bits that can be processed per access) and, more importantly, memory size, or the number of words that it can access. 

## Notes

* We can use the struct literal form to initialize a value from a struct type.
* The dot (.) operator allows us to access individual field values.
* We can create anonymous structs.

## Quotes

_"Implicit conversion of types is the Halloween special of coding. Whoever thought of them deserves their own special hell." - Martin Thompson_

## Links

[Understanding Type in Go](https://www.ardanlabs.com/blog/2013/07/understanding-type-in-go.html) - William Kennedy    
[Object Oriented Programming in Go](https://www.ardanlabs.com/blog/2013/07/object-oriented-programming-in-go.html) - William Kennedy    
[Padding is hard](https://dave.cheney.net/2015/10/09/padding-is-hard) - Dave Cheney    
[Structure Member Alignment, Padding and Data Packing](https://www.geeksforgeeks.org/structure-member-alignment-padding-and-data-packing/)    
[The Lost Art of Structure Packing](http://www.catb.org/esr/structure-packing) - Eric S. Raymond    

## Code Review

[Declare, create and initialize struct types](https://github.com/deeprajsshetty/GolangTraining/blob/master/004-Language%20Syntax/002-Struct%20Types/001-DeclareCreateAndInitStructTypes/main.go) | 
[Anonymous struct types](https://github.com/deeprajsshetty/GolangTraining/blob/master/004-Language%20Syntax/002-Struct%20Types/002-AnonymousStructTypes/main.go) | 
[Named vs Unnamed types](https://github.com/deeprajsshetty/GolangTraining/blob/master/004-Language%20Syntax/002-Struct%20Types/003-NamedVsUnnamedTypes/main.go) 

## Advanced Code Review

[Struct type alignments](https://github.com/deeprajsshetty/GolangTraining/blob/master/004-Language%20Syntax/002-Struct%20Types/004-StructTypeAlignment/main.go)

## Exercises

### Exercise 1

**Part A:** Declare a struct type to maintain information about a user (name, email and age). Create a value of this type, initialize with values and display each field.

**Part B:** Declare and initialize an anonymous struct type with the same three fields. Display the value.

[Solution](https://github.com/deeprajsshetty/GolangTraining/blob/master/004-Language%20Syntax/002-Struct%20Types/005-Exercise/main.go)
___
