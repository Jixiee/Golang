//Write a program to convert temperatures from C to F and vice versa.
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Temperature Converter")

	fmt.Print("Enter temperature value: ")
	var temperature float64
	fmt.Scanln(&temperature)

	fmt.Print("Enter source unit (F/C): ")
	var sourceUnit string
	fmt.Scanln(&sourceUnit)

	var convertedTemperature float64

	switch sourceUnit {
	case "F", "f":
		convertedTemperature = (temperature - 32) * 5 / 9
		fmt.Printf("%.2f°F is equal to %.2f°C\n", temperature, convertedTemperature)
	case "C", "c":
		convertedTemperature = (temperature * 9 / 5) + 32
		fmt.Printf("%.2f°C is equal to %.2f°F\n", temperature, convertedTemperature)
	default:
		fmt.Println("Invalid source unit. Please enter F or C.")
	}
}
