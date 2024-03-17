package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			processInput(arg)
		}
	} else {
		var input string
		fmt.Print("Enter a number: ")
		fmt.Scanln(&input)
		processInput(input)
	}
}

func processInput(input string) {
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("Invalid input: %s\n", input)
		return
	}

	// Perform unit conversions here
	// Example: temperature conversions
	celsius := value
	fahrenheit := celsius*9/5 + 32
	fmt.Printf("%.2f°C is %.2f°F\n", celsius, fahrenheit)

	// Example: length conversions
	feet := value
	meters := feet * 0.3048
	fmt.Printf("%.2f feet is %.2f meters\n", feet, meters)

	// Example: weight conversions
	pounds := value
	kilograms := pounds * 0.45359237
	fmt.Printf("%.2f pounds is %.2f kilograms\n", pounds, kilograms)

	// Add more unit conversions as needed
}
