// chapter 6 exercise 2 - Write a function that takes an integer and halves it and returns true if it was even or false if it was odd. For example, half(1) should return (0, false) and half(2) should return (1, true).

package main

import "fmt"

func half(full int) (int, string) {
	halved := full / 2
	eoo := halved % 2
	if eoo == 0 {
		return halved, "even"
	} else {
		return halved, "odd"
	}
}

func main() {
	var i int = 6
	x, y := half(i)
	fmt.Println(x, y)
}
