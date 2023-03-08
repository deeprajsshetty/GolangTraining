## Variables

Variables are at the heart of the language and provide the ability to read from and write to memory. In Go, access to memory is type safe. This means the compiler takes type seriously and will not allow us to use variables outside the scope of how they are declared.

## Notes

- The purpose of all programs and all parts of those programs is to transform data from one form to the other.
- Code primarily allocates, reads and writes to memory.
- Understanding type is crucial to writing good code and understanding code.
- If you don't understand the data, you don't understand the problem.
- You understand the problem better by understanding the data.
- When variables are being declared to their zero value, use the keyword var.
- When variables are being declared and initialized, use the short variable declaration operator.

## Links

[Built-In Types](http://golang.org/ref/spec#Boolean_types)  
[Variables](https://golang.org/doc/effective_go.html#variables)  
[Gustavo's IEEE-754 Brain Teaser](https://www.ardanlabs.com/blog/2013/08/gustavos-ieee-754-brain-teaser.html) - William Kennedy  
[What's in a name](https://www.youtube.com/watch?v=sFUSP8Au_PE)  
[A brief history of “type”](http://arcanesentiment.blogspot.com/2015/01/a-brief-history-of-type.html) - Arcane Sentiment

## Code Review

### Zero Values - Type Initialized Value

    - Boolean false
    - Integer 0
    - Floating Point 0
    - Complex 0i
    - String "" (empty string)
    - Pointer nil

[Declare and initialize variables](https://github.com/deeprajsshetty/GolangTraining/blob/master/004-Language%20Syntax/004.001-Variables/001-DeclareAndInitializeVariables/main.go)

## Exercises

### Exercise 1

**Part A:** Declare three variables that are initialized to their zero value and three declared with a literal value. Declare variables of type string, int and bool. Display the values of those variables.

**Part B:** Declare a new variable of type float32 and initialize the variable by converting the literal value of Pi (3.14).

[Solution](https://github.com/deeprajsshetty/GolangTraining/blob/master/004-Language%20Syntax/004.001-Variables/002-Exercise/main.go) 

---

