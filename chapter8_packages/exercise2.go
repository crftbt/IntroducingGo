// chapter 8 exercise 2 - What is the difference between an identifier that starts withat a capital letter and one that doesn't (e.g., Average versus average)?

package main

import (
	"fmt"
	m "introducing_go/chapter8_packages/packagetesting"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs) // public function from package here
	fmt.Println(avg)
}
