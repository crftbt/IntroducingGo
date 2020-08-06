// chapter 6 exercise 11 - Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).

package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x := 1
	y := 2
//	xold := x
//	yold := y
//	x = yold
//	y = xold
	swap(&x, &y)
	fmt.Println("x:", x, "y:", y)
}
