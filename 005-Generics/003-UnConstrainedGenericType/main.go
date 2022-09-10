package main

import "fmt"

// Student
type Student struct {
	Name string
	ID   int
	Age  float64
}

// addStudent - Using Unconstrained Generic Type "any"
// Using type inferences, the compiler uses the default conversions of string, int,
// and Student to allow program output by converting each of these types to string.
func addStudent[T any](students []T, student T) []T {
	return append(students, student)
}

func main() {
	studentsName := []string{} // Empty Slice
	resultName := addStudent(studentsName, "StudentByName001")
	resultName = addStudent(resultName, "StudentByName002")
	resultName = addStudent(resultName, "StudentByName003")
	fmt.Println(resultName)

	studentsID := []int{} // Empty Slice
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
[{StudentByStruct001 1 24} {StudentByStruct001 1 24} {StudentByStruct001 1 24}]
*/
