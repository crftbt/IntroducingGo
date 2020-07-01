// Chapter 4 - Exercise 2 - Print all numbers between 1 and 100 that are evenly divicable by 3 (i.e., 3, 6, 9, etc.).

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("We got", i, "rock")
		}
	}
}
