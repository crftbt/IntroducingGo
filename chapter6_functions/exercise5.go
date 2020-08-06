//chapter 6 exercise 5 - The Fibonacci sequence is not a big truck. It is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function that can find fib(n).

package main

import "fmt"

func fibbers(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibbers(n-1) + fibbers(n-2)
	}
}

func main() {
	fmt.Println(fibbers(2))
}
