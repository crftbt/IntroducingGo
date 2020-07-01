package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 10
	name, ok := y[2]
	fmt.Println(name, ok)
	if name, ok := y[2]; ok {
		fmt.Println(name, ok)
	}
}
