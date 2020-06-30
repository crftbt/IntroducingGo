// What is the difference between var and const?

package main

import "fmt"

func main() {
	const myconst int = 1
	var myvar int = 1
	fmt.Println("myconst:", myconst, "myvar:", myvar)
	mysconst = 2
	myvar = 2
	fmt.Println("myconst:", myconst, "myvar:", myvar)
}
