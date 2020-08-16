// chapter 8 exercise 4 - Create Min and Max function that find minimum and maximum values in a slice of float64s in a package.

package main

import (
	"fmt"
	m "introducing_go/chapter8_packages/packagetesting" // package alias made here
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs) // package alias used here
	fmt.Println("Average: ", avg)
	min := m.Minimum(xs)
	fmt.Println("Minimum: ", min)
	max := m.Maximum(xs)
	fmt.Println("Maximum: ", max)
}
