// chapter 6 exercise 2 - Write a function that takes an integer and halves it and returns true if it was even or false if it was odd. For example, half(1) should return (0, false) and half(2) should return (1, true).

package main

import "fmt"

func half(full int) (int, bool) {
	return full / 2, full%2 == 0
}

func main() {
	var i int = 2
	x, y := half(i)
	fmt.Println(x, y)
}
