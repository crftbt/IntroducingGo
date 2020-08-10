//chapter 6 exercise 1 - what's the difference between a method and a function?

package main

import (
	"fmt"
	"math"
)

func (c *Circle) area() float64 { //method receiver is here
	return math.Pi * c.r * c.r
}

type Circle struct {
	x, y, r float64
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.area())
}
