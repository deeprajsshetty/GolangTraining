package main

import (
	"fmt"
	"sort"
)

// Ordered - Based on existing Datatypes.
// It can be common interface which includes all the Types in Golang.
type Ordered interface {
	~int | ~float64 | ~string
}

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

/*
We define OrderedSlice as follows and provide the required group of Len, Less, and Swap.
*/
// Group of functions that ensure that an OrderedSlice can be sorted
type OrderedSlice[T Ordered] []T // T must implement < and >
func (s OrderedSlice[T]) Len() int {
	return len(s)
}
func (s OrderedSlice[T]) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s OrderedSlice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// end group for OrderedSice

/*
We introduce another type, SortType, along with the required group of Len, Less, and Swap.
*/
// Group of functions that ensure that SortType can be sorted
type SortType[T any] struct {
	slice   []T
	compare func(T, T) bool
}

func (s SortType[T]) Len() int {
	return len(s.slice)
}
func (s SortType[T]) Less(i, j int) bool {
	return s.compare(s.slice[i], s.slice[j])
}
func (s SortType[T]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

// end of group for SortType

// PerformSort - The user of PerformSort must supply a function for comparing two elements of type T.
func PerformSort[T any](slice []T, compare func(T, T) bool) {
	sort.Sort(SortType[T]{slice, compare})
}

func main() {
	studentsName := []string{} // Empty Slice
	resultName := addStudent(studentsName, "StudentByName001")
	resultName = addStudent(resultName, "StudentByName002")
	resultName = addStudent(resultName, "StudentByName003")
	// Reference to OrderedSlice
	sort.Sort(OrderedSlice[string](resultName))
	fmt.Println(resultName)

	studentsID := []int{} // Empty Slice
	resultID := addStudent(studentsID, 1)
	resultID = addStudent(resultID, 2)
	resultID = addStudent(resultID, 3)
	// Reference to OrderedSlice
	sort.Sort(OrderedSlice[int](resultID))
	fmt.Println(resultID)

	studentsStruct := []Student{} // Empty Slice
	resultStruct := addStudent(studentsStruct, Student{Name: "StudentByStruct001", ID: 1, Age: 24})
	resultStruct = addStudent(resultStruct, Student{Name: "StudentByStruct001", ID: 1, Age: 24})
	resultStruct = addStudent(resultStruct, Student{Name: "StudentByStruct001", ID: 1, Age: 24})
	// Reference to SortType
	PerformSort(resultStruct, func(s1, s2 Student) bool {
		return s1.Age < s2.Age // comparing two Student values
	})
	fmt.Println(resultStruct)
}

/* Output
[StudentByName001 StudentByName002 StudentByName003]
[1 2 3]
[{StudentByStruct001 1 24} {StudentByStruct001 1 24} {StudentByStruct001 1 24}]
*/
