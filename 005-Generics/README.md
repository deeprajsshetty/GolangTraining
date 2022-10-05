## Non Generic

We present a set of related problems of adding a new student to an existing slice of students. First, we add just the name of the student to our existing slice. Next, we add the student’s ID number to a slice containing ID numbers. Next, we add a struct containing name, ID, and age to an existing slice of structs

Having to add a new function each time we wish to add a new underlying data type to our various student collections is tedious and a major downside to earlier versions of Go.

We can solve this problem using Generics.

## Introducing Generic

### Enter Go, Version 1.18, that introduces support for generics.

**Constrained Generic Type**  
* 005-Generics -> 002-ConstrainedGenericType explained the details with Example ***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/blob/master/005-Generics/002-ConstrainedGenericType)***.
* Using ***Interface*** Type we can achieve this.

**Unconstrained Generic Type any**  
* 005-Generics -> 003-UnConstrainedGenericType explained the details with Example ***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/tree/master/005-Generics/003-UnConstrainedGenericType)***.
* Using ***Any*** we can achieve this.
* The Go compiler can do type inferencing if we replace the constrained generic type ***[T Stringer]*** with the unconstrained type ***any***.

## Benefits of Generics

**Using Go’s Sort Package**  
* The Sort function in Go’s sort package, sort.Sort, requires that the type in the slice being sorted must implement three methods:***Len, Less, Swap***
* We show how a generic collection implemented as a slice can be sorted.
* 005-Generics -> 004-BenefitsOfGenerics -> 001-UsingGoSortPackage explained the details with Example ***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/tree/master/005-Generics/004-BenefitsOfGenerics/001-UsingGoSortPackage)***.

**Map Functions**  

***Non Generic***
* Map functions in Go are commonplace and perform a transformation in a slice to produce a new slice with the transformed results.
* 005-Generics -> 004-BenefitsOfGenerics -> 002-MapFunctions -> 01-NonGeneric explained the details with Example ***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/tree/master/005-Generics/004-BenefitsOfGenerics/002-MapFunctions/01-NonGeneric)***.

***Generic***
* MyMap can be made genric.
* 005-Generics -> 004-BenefitsOfGenerics -> 002-MapFunctions -> 02-Generic explained the details with Example ***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/tree/master/005-Generics/004-BenefitsOfGenerics/002-MapFunctions/02-Generic)***.

**Filter Functions**  

***Non Generic***
* Filter functions in Go are also commonplace and perform a filtering operation on an input slice 
based on a function passed in.
* 005-Generics -> 004-BenefitsOfGenerics -> 003-FilterFunctions -> 01-NonGeneric explained the details with Example ***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/tree/master/005-Generics/004-BenefitsOfGenerics/003-FilterFunctions/01-NonGeneric)***.

***Generic***
* MyFilter can be made generic.
* 005-Generics -> 004-BenefitsOfGenerics -> 003-FilterFunctions -> 02-Generic explained the details with Example ***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/tree/master/005-Generics/004-BenefitsOfGenerics/003-FilterFunctions/02-Generic)***.

## Examples

**Generics Using WaitGroup**  
* 005-Generics -> 005-Examples -> 001-GenericsUsingWaitGroup ***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/tree/master/005-Generics/005-Examples/001-GenericsUsingWaitGroup)***.

