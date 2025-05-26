package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags for our two numbers
	var num1 = flag.Float64("n1", 0.0, "First number to add")
	var num2 = flag.Float64("n2", 0.0, "Second number to add")

	// Parse the flags
	flag.Parse()

	// Calculate the sum
	result := *num1 + *num2

	// Print the result
	fmt.Printf("%.2f + %.2f = %.2f\n", *num1, *num2, result)
}
