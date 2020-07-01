package main

import "fmt"

func main() {
	var x string
	x = "first "
	fmt.Println("x =", x)
	x = x + "second"
	fmt.Println("x = x + \"second\" =", x)
}
