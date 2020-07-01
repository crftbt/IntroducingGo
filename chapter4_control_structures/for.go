// Chapter 4 - Test a for loop

package main

import "fmt"

func main() {
	fmt.Println("Verbose boi:")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("Condensed boi:")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
