// Introducing Go - Chapter 6 Functions - Testing Closure

package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	for i := 0; i <= 10; i++ {
		fmt.Println(nextEven())
	}
}
