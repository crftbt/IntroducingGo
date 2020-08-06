package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	zero(&x) //Memory address of variable here
	fmt.Println(x)
}
