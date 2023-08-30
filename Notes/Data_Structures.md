# Data Structures
Data structures are fundamental building blocks in programming that allow you to organize and manipulate data efficiently. In this section, we'll explore various data structures available in GoLang.

## Arrays, Slices, and Maps
### Arrays
Arrays in GoLang are fixed-size collections of elements, all of the same data type. The size of an array is determined at the time of declaration and cannot be changed during program execution.

- Declaration
```
var numbers [5]int // An array of 5 integers
```

- Initialization
```
numbers := [5]int{1, 2, 3, 4, 5}
```

- Usage:
Arrays are accessed using zero-based indexing.
- For example, to access the first element of the array:
```
firstElement := numbers[0] // Value: 1
```

- Arrays are useful when you need a fixed number of elements with a known size. However, their fixed size can be limiting in situations where you want to dynamically adjust the number of elements.

### Slices
Slices are dynamic and flexible data structures built on top of arrays. They allow you to work with sequences of elements and adjust their size as needed. Slices are commonly used in GoLang for various data manipulation tasks.


- Declaration
```
var numbers []int // A slice of integers
```


- Initialization: Slices can be initialized from an existing array, another slice, or using the built-in make function:
```
numbers := []int{1, 2, 3, 4, 5} // Initialized from a literal
```

### Slicing: 
Slicing is a powerful feature of slices that allows you to create new slices from existing ones.
```
sliced := numbers[1:4] // Creates a new slice from index 1 to 3 (inclusive)
```

- Dynamic Sizing: One of the key advantages of slices is their ability to grow or shrink in size dynamically.
```
numbers = append(numbers, 6) // Adds an element to the end of the slice
```

- Underlying Array:Slices are built on top of arrays. When you create a slice, GoLang implicitly creates an underlying array to hold the data.
```
s1 := numbers[1:4]
s2 := numbers[2:5]
```
Both s1 and s2 share the same underlying array. Modifying elements in s1 will affect numbers and elements in s2.

- Length and Capacity: Slices have a length (number of elements) and a capacity (number of elements in the underlying array from the starting index of the slice). The len and cap functions provide the length and capacity of a slice, respectively.
```
s := []int{1, 2, 3, 4, 5}
length := len(s) // 5
capacity := cap(s) // 5 (as the entire array is used)
```

- Copying Slices: To create an independent copy of a slice, you can use the copy function.
```
copyOfS := make([]int, len(s))
copy(copyOfS, s)
```
### Slices with Strings
Strings in GoLang are sequences of characters, and you can use slices to work with different portions of a string.

- Creating a Slice from a String: Let's say we have a string and we want to create a slice to extract a portion of it:
```
originalString := "Hello, World!"
substring := originalString[7:12]
```
In this example, substring will contain the characters "World" from the original string. The slice expression [7:12] specifies the range of indices to extract.

- Modifying a Slice: Slices allow you to modify the underlying string as well.
- For instance:
```
stringSlice := []byte("Hello, World!")
stringSlice[0] = 'J'
modifiedString := string(stringSlice)
```
Here, we're converting the original string to a byte slice, modifying the first character to 'J', and then converting it back to a string.


- Dynamic Sizing with Slices: Slices also work well for dynamic sizing, like appending characters to a string:
```
stringSlice := []byte("Hello")
stringSlice = append(stringSlice, ' ')
stringSlice = append(stringSlice, []byte("World!")...)
finalString := string(stringSlice)
```
In this example, we're appending a space and the "World!" string to the stringSlice.


- Slices are a versatile tool in GoLang, providing dynamic sizing and a convenient interface for working with collections of data. They are widely used for various tasks, including iterating over data, managing collections, and handling sequences of elements efficiently.

## Maps
Maps are collections of key-value pairs, where each key is unique and used to access its associated value.


- Declaring a Map
```
var ages map[string]int // A map with string keys and integer values
```

- Initializing a Map
```
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
}
```

## Structs and User-Defined Types
### Structs
Structs are user-defined composite types that group fields of different types under a single name.


- Declaring a Struct
```
type Person struct {
    FirstName string
    LastName  string
    Age       int
}
```
- Initializing a Struct
```
person := Person{
    FirstName: "Alice",
    LastName:  "Smith",
    Age:       28,
}
```
### User-Defined Types
GoLang allows you to create your own types by aliasing existing types.

### Aliasing Types
```
type Celsius float64
type Fahrenheit float64

func main() {
    var c Celsius = 25.0
    var f Fahrenheit = 77.0

    c = Celsius(f) // Conversion between aliased types
}
```
## Pointers and References
### Pointers
Pointers hold the memory address of a value instead of the value itself. They are used to indirectly access and modify values.

- Declaring a Pointer
```
var x int
var ptr *int // A pointer to an integer
ptr = &x    // Assign the address of x to ptr
```
- Dereferencing a Pointer
```
*x = 42 // Assign the value 42 to the memory location pointed by ptr
```
### References
In GoLang, all function arguments are passed by value. However, slices and maps are reference types, meaning the underlying data is shared between different references.
```
func modifySlice(s []int) {
    s[0] = 100
}

func main() {
    numbers := []int{1, 2, 3}
    modifySlice(numbers)
    fmt.Println(numbers) // Output: [100, 2, 3]
}
```
