package main

import "fmt"

func main() {
	fmt.Println("(true && true) || (false && true) || !(false && false) =")
	fmt.Println((true && true) || (false && true) || !(false && false))
}
