package main

import "fmt"

func main() {
	fmt.Println(f1())
}

func f1() string {
	return "BEES!? " + f2()
}

func f2() string {
	return "BEADS!?"
}
