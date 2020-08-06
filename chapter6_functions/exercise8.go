package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0 //value assigned to pointer here
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}
