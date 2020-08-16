// chapter 8 exercise 1 - Why do we use packages?

package main

import (
	"fmt"
	m "introducing_go/chapter8_packages/packagetesting" //package with alias here
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
}
