# Concurrency and Goroutines
Concurrency is the ability of a program to handle multiple tasks concurrently. GoLang offers built-in support for concurrency through Goroutines and Channels.

## Introduction to Concurrency and Parallelism
### Concurrency vs. Parallelism
- Concurrency is the concept of handling multiple tasks simultaneously, where tasks may not necessarily execute at the exact same time.
- Parallelism refers to the actual execution of multiple tasks simultaneously by utilizing multiple CPUs or cores.
- GoLang focuses on making it easy to write concurrent programs, which can then be executed in parallel if the underlying hardware supports it.
- Goroutines: Creating and Managing Concurrent Tasks
  
### Goroutines
A Goroutine is a lightweight, independently executing function that can be created and scheduled by the Go runtime. Goroutines allow you to perform concurrent tasks without creating multiple threads, making them efficient and manageable.

### Creating Goroutines
Goroutines are created using the go keyword followed by a function call.
```
func main() {
    go printNumbers()
    go printLetters()
    // ...
}
```

### Synchronization
When using Goroutines, it's important to ensure proper synchronization to prevent race conditions and ensure consistent behavior.
```
import(
 "sync"
)
var wg sync.WaitGroup

func main() {
    wg.Add(2)
    go printNumbers()
    go printLetters()
    wg.Wait()
}

func printNumbers() {
    defer wg.Done()
    // ...
}

func printLetters() {
    defer wg.Done()
    // ...
}
```
The sync.WaitGroup is used to wait for all Goroutines to complete before the program exits.

# Channels for Communication Between Goroutines
## Channels
Channels are used for communication and synchronization between Goroutines. They provide a way for one Goroutine to send data to another.

- ### Creating Channels
Channels are created using the make function.
```
ch := make(chan int) // Create an unbuffered channel of integers
```
- ### Sending and Receiving
Channels support sending and receiving data using the <- operator.
```
ch <- value // Send value to the channel
receivedValue := <-ch // Receive value from the channel
```
- ### Buffered Channels
Buffered channels allow a certain number of values to be stored before sending blocks.
```
ch := make(chan int, 5) // Create a buffered channel with a capacity of 5
```
- ### Closing Channels
Closing a channel signals that no more values will be sent.
```
close(ch)
```
- ### Select Statement
The select statement is used to wait for multiple channel operations.
```
select {
case value := <-ch1:
    // Process value from ch1
case value := <-ch2:
    // Process value from ch2
case ch3 <- newValue:
    // Send newValue to ch3
}
```

- Concurrency and Goroutines are powerful features of GoLang that allow you to create efficient and concurrent programs. By understanding Goroutines, synchronization, and communication using channels, you can take full advantage of Go's concurrency model.
