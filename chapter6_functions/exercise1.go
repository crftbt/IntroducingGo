// Introducing Go - Chapter 6 - Exercise 1 - sum is a function that take a slice of numbers and adds them together. What would its function signature look like in go?

package main

import "fmt"

func sum(xs []int) int { // Hello, function signature here.
	total := 0
	for _, v := range xs {
		total += v
	}
	return total
}

func main() {
	xs := []int{1, 2, 3, 4}

//	total := 0

	fmt.Println(sum(xs))
}
