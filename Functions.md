# Functions

## Defining Functions and Their Syntax:
In GoLang, a function is a self-contained block of code that performs a specific task. Functions help in modularizing your code and making it more organized and readable.

#### Syntax of Function Declaration
```go
func <functionName>(parameters) returnType {
    // Function body
    return result
}
```
- func: Keyword to define a function.
- functionName: Name of the function (followed by parentheses).
- parameters: Input values that the function receives (if any).
- returnType: Data type of the value the function returns.
- return: Keyword used to send a value back from the function (optional).

## Function Arguments and Return Values
### Function Arguments:
- Functions can accept zero or more arguments (parameters) to work with.
- Each parameter specifies its name and data type.
- Arguments are passed by value, meaning the function receives a copy of the argument's value.
```go
func add(a int, b int) int {
    return a + b
}
```
- GoLang allows you to omit the type name for consecutive parameters of the same type.
```go
func multiply(x, y int) int {
    return x * y
}
```

### Return Values:
- Functions can return one or more values.
- The return statement is used to send values back to the caller.
- Multiple return values are enclosed in parentheses.
```go
func divideAndRemainder(dividend, divisor int) (int, int) {
    quotient := dividend / divisor
    remainder := dividend % divisor
    return quotient, remainder
}
```

### Named return values:
- You can specify variable names for return values in the function signature. These variables act as placeholders and are returned automatically.
```go
func divideAndRemainder(dividend, divisor int) (quotient, remainder int) {
    quotient = dividend / divisor
    remainder = dividend % divisor
    return // No explicit return values needed
}
```

## Anonymous Functions and Closures
### Anonymous Functions:
- Anonymous functions, also known as lambda functions, are functions without a name.
- They can be assigned to variables and invoked later.
```go
func main() {
    greet := func() {
        fmt.Println("Hello from an anonymous function!")
    }
    greet()
}
```
### Closures:
- Closures are anonymous functions that can access variables defined outside their body.
- They "close over" variables from their surrounding scope.
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    increment := counter()
    fmt.Println(increment()) // 1
    fmt.Println(increment()) // 2
}
```
- Closures can be used to create stateful functions that retain their internal state between calls.
