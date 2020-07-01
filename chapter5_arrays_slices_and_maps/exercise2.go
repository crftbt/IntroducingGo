// Chapter 5 - Exercise 2 - Print the length of the slice

package main

import "fmt"

func main() {
	x := make([]int, 3, 9)
	fmt.Println(len(x))
}
