// chapter 8 exercise 3 - What is a package alias? How do you make one?

package main

import (
	"fmt"
	m "introducing_go/chapter8_packages/packagetesting" // package alias made here
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs) // package alias used here
	fmt.Println(avg)
}
