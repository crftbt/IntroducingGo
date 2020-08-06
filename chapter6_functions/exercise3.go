// chapter 6 exercise 3 - Write a function with one variadic parameter that finds the greatest number in a list of numbers.
package main

import "fmt"

func greatestNumber(args ...int) int {
	highaf := args[0]
	for _, arg := range args {
		if arg > highaf {
			highaf = arg
		}
	}
	return highaf
}

func main() {
	xs := []int{7, 238, 83, 3434, 838383, 31}
	fmt.Println(greatestNumber(xs...))
}
