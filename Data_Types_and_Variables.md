# Data Types and Variables

## Data Types:-
In GoLang, data types define the kind of values that variables can hold. Here are some common data types:

```Integers```
Integers represent whole numbers. GoLang provides various sizes of integers, including int8, int16, int32, and int64.
Unsigned integers, like uint8, uint16, uint32, and uint64, represent positive whole numbers only.
The int and uint types are machine-dependent and have the size of the architecture's word size (32 or 64 bits).

```Floats```
Floats represent decimal numbers. GoLang has two primary floating-point types: float32 and float64.
float32 is a single-precision floating-point type, while float64 is a double-precision floating-point type with higher precision.

```Strings```
Strings represent sequences of characters. They are created using double quotes or backticks.
GoLang treats strings as immutable, meaning you can't modify individual characters directly.

```Booleans```
Booleans represent true or false values. The bool type is used for boolean variables.
Logical operators like && (AND), || (OR), and ! (NOT) are used with boolean values.

```Other Data Types```
GoLang also has data types like byte (alias for uint8), rune (alias for int32, used to represent Unicode code points), and complex64 and complex128 (for complex numbers). 

## Variable Declaration, Initialization, and Scoping:-
### Variable Declaration:
Variables are declared using the var keyword, followed by the variable name and type.
var age int
var name string

You can also declare multiple variables of the same type in a single statement.
var x, y, z int

### Variable Initialization:
Variables can be initialized with values at the time of declaration.
var count int = 10
var message string = "Hello, Go!"

### GoLang can infer the variable type based on the provided value.
var score = 95 // score is inferred as int

### Short variable declaration := is used to declare and initialize variables without explicitly stating the type.
speed := 50 // Inferred as int

### Variable Scoping:
In GoLang, variables have block scope. They are only accessible within the block of code where they are declared.
Function parameters and local variables are block-scoped.

## Constants and iota:-
### Constants:
Constants are named values that cannot be changed after they are defined.
They are created using the const keyword.
const pi = 3.14
const daysInWeek = 7
Constants can be of any data type.

### iota:
iota is a special pre-declared identifier in GoLang that's used in constant declarations.
It starts from 0 and increments by 1 for each subsequent constant.
const (
    Sunday    = iota // 0
    Monday           // 1
    Tuesday          // 2
    // and so on
)
iota can be used to create sequential values for enums or bit flags.
