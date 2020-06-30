// Exercise 5 - Convert from Fahrenheit to Celsius

package main

import "fmt"

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	var input_fahrenheit float32
	fmt.Scanf("%f", &input_fahrenheit)

	celsius := (input_fahrenheit - 32) * 5 / 9

	fmt.Println("Celsius:", celsius)
}
