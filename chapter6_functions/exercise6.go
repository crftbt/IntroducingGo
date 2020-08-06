package main

import "fmt"

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("RECOVERED:", x)
		}
	}()
	panic("PANIC IN THE DISCO")
}
