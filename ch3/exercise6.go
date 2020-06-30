// Exercise 6 - Convert user input feet to meters

package main

import "fmt"

func main() {
	fmt.Print("Enter a girth in feet: ")
	var input_girth float32
	fmt.Scanf("%f", &input_girth)

	meters := input_girth * 0.3048

	fmt.Println("Meters:", meters)
}
