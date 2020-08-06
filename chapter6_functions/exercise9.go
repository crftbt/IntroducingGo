package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int) //new pointer created here
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}
