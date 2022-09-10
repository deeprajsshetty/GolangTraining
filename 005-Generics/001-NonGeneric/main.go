package main

import "fmt"

// Student Struct
type Student struct {
	Name string
	ID   int
	age  float64
}

func main() {
	studentsName := []string{} // Empty Slice
	resultName := addStudent(studentsName, "StudentByName001")
	resultName = addStudent(resultName, "StudentByName002")
	resultName = addStudent(resultName, "StudentByName003")
	fmt.Println(resultName)

	studentsID := []int{} // Empty Slice
	resultID := addStudentID(studentsID, 1)
	resultID = addStudentID(resultID, 2)
	resultID = addStudentID(resultID, 3)
	fmt.Println(resultID)

	studentsStruct := []Student{} // Empty Slice
	resultStruct := addStudentStruct(studentsStruct, Student{Name: "StudentByStruct001", ID: 1, age: 24})
	resultStruct = addStudentStruct(resultStruct, Student{Name: "StudentByStruct001", ID: 1, age: 24})
	resultStruct = addStudentStruct(resultStruct, Student{Name: "StudentByStruct001", ID: 1, age: 24})
	fmt.Println(resultStruct)
}

// Adding New Student By Name
func addStudent(students []string, student string) []string {
	return append(students, student)
}

// Adding New Student By ID Number
func addStudentID(students []int, student int) []int {
	return append(students, student)
}

// Adding New Student by Student Struct
func addStudentStruct(students []Student, student Student) []Student {
	return append(students, student)
}

/* Output
[StudentByName001 StudentByName002 StudentByName003]
[1 2 3]
[{StudentByStruct001 1 24} {StudentByStruct001 1 24} {StudentByStruct001 1 24}]
*/
