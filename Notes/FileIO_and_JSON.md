# File I/O and JSON
File I/O (Input/Output) and JSON encoding/decoding are fundamental tasks in software development. GoLang provides efficient and straightforward ways to handle these operations.

## Reading and Writing Files
### Reading Files
- Use the os package to open and read files.
- You can read files byte by byte, or line by line.
```
package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
)

func main() {
	filePath := "example.txt"

	// Read entire file into memory
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
```
### Writing Files
- Use the os package to create and write to files.
- You can write strings, bytes, or other data types.
```
package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	filePath := "output.txt"

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := "Hello, File I/O!"
	_, err = file.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File written successfully.")
}
```
## Encoding and Decoding JSON Data
### Encoding JSON
- Use the encoding/json package to convert Go data structures to JSON.
```
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func main() {
	person := Person{
		Name: "Jigyasa",
		Age:  200,
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}
```
### Decoding JSON
- Use the encoding/json package to convert JSON to Go data structures.
```
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData := `{"name":"Shambhavi","age":300}`

	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
```
