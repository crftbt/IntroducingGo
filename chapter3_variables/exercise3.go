// What is scope? How do you determine the scope of a variable in Go?

package main

import "fmt"

var levariable int = 1

func main() {
	fmt.Println("levariable global: ", levariable)
	levariable := 2
	fmt.Println("levariable local: ", levariable)
	letsgo()
}

func letsgo() {
	fmt.Println("levariable global still: ", levariable)
}
