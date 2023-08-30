# Control Structures
Control structures are essential programming constructs that help you control the flow of execution in your code. They allow you to make decisions, select specific paths, and repeat tasks as needed. In this section, we'll explore three fundamental control structures: if statements, switch statements, and various forms of loops.

## If Statements and Conditional Logic
- If statements allow you to execute certain blocks of code based on whether a given condition is true or false. They are vital for implementing decision-making in your programs.

### Syntax of If Statement
```
if condition {
    // Code to execute if condition is true
} else {
    // Code to execute if condition is false
}
```
- The if keyword marks the start of the statement.
- The condition is a boolean expression that determines whether the code inside the if block should execute.
- The else block is optional and contains code that executes when the condition is false.

### Example of If Statement
```
func main() {
    age := 20

    if age >= 18 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a minor.")
    }
}
```
## Switch Statements
Switch statements provide a concise way to handle multiple cases for a single value. They are particularly useful when you want to compare a single value against different possibilities.

### Syntax of Switch Statement
```
switch expression {
case value1:
    // Code to execute if expression matches value1
case value2:
    // Code to execute if expression matches value2
default:
    // Code to execute if no cases match
}
```
- The switch keyword begins the statement.
- The expression is the value you want to compare.
- Each case represents a possible value that the expression can match.
- The default case is optional and is executed when none of the cases match.
### Example of Switch Statement
```
func main() {
    day := "Monday"

    switch day {
    case "Monday":
        fmt.Println("It's the start of the week.")
    case "Friday":
        fmt.Println("It's the end of the week.")
    default:
        fmt.Println("It's a regular day.")
    }
}
```
## Loops:
- Loops are used to repeat a set of statements multiple times.
- GoLang offers two primary types of loops: for loops and range loops.

### >>For Loops
For loops are used when you know the number of iterations you want to perform.

### Syntax of For Loop
```
for initialization; condition; post {
    // Code to execute in each iteration
}
```
- The for keyword initiates the loop.
- The initialization is a statement executed before the loop starts (often used for variable initialization).
- The condition is checked before each iteration.
- The post statement is executed after each iteration (used for updating variables).
### Example of For Loop
```
func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```
### >>Range Loops
Range loops are used to iterate over elements in a collection, like arrays, slices, maps, and strings.

### Syntax of Range Loop
```
for index, value := range collection {
    // Code to execute in each iteration
}
```
- The range keyword is used to iterate over elements in the collection.
- index represents the index or key of the current element.
- value represents the value of the current element.
### Example of Range Loop
```
func main() {
    numbers := []int{1, 2, 3, 4, 5}

    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
```
