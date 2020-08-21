// chapter 10 exercise 1 - How do you specify the direction of a channel type?

package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) { // direction specified here
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan<- string) { // direction specified here
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) { // direction specified here
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
