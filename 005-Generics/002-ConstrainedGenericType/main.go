package main

import "fmt"

// Stringer - A type Stringer is defined as an interface containing a single method signature:
// String() string.
/*
Any entity that implements this type by having a well-defined String() definition
can be converted to a string for output purposes. Since we wish to be able to output
our various student collections (slices), we constrain the data type, T, in the
signature of our generic addStudent function to be of type Stringer.
*/
type Stringer = interface {
	String() string
}

/*
In Go, one implements an interface implicitly by implementing the function(s)
specified in the interface definition. In this case, any type that implements
a String() function can be considered to be of type Stringer.
*/

// Integer type implements a String() function can be considered to be of type Stringer.
type Integer int

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

// String type implements a String() function can be considered to be of type Stringer.
type String string

func (s String) String() string {
	return string(s)
}

// Student type implements a String() function can be considered to be of type Stringer.
type Student struct {
	Name string
	ID   int
	Age  float64
}

func (s Student) String() string {
	return fmt.Sprintf("%s %d %0.2f", s.Name, s.ID, s.Age)
}

// addStudent - The generic signature of our single addStudent function
/*
Types Integer, String, and Student are defined along with their definitions of
String() so that we can use generic function addStudent using each of these
Stringer types.
*/
func addStudent[T Stringer](students []T, student T) []T {
	return append(students, student)
}

func main() {
	studentsName := []String{} // Empty Slice
	resultName := addStudent(studentsName, "StudentByName001")
	resultName = addStudent(resultName, "StudentByName002")
	resultName = addStudent(resultName, "StudentByName003")
	fmt.Println(resultName)

	studentsID := []Integer{} // Empty Slice
	resultID := addStudent(studentsID, 1)
	resultID = addStudent(resultID, 2)
	resultID = addStudent(resultID, 3)
	fmt.Println(resultID)

	studentsStruct := []Student{} // Empty Slice
	resultStruct := addStudent(studentsStruct, Student{Name: "StudentByStruct001", ID: 1, Age: 24})
	resultStruct = addStudent(resultStruct, Student{Name: "StudentByStruct001", ID: 1, Age: 24})
	resultStruct = addStudent(resultStruct, Student{Name: "StudentByStruct001", ID: 1, Age: 24})
	fmt.Println(resultStruct)
}

/* Output
[StudentByName001 StudentByName002 StudentByName003]
[1 2 3]
[StudentByStruct001 1 24.00 StudentByStruct001 1 24.00 StudentByStruct001 1 24.00]
*/
