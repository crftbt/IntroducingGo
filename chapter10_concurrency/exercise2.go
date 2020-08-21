// chatper 10 exercise 2 -

package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Sleeping for")

	Sleep(10000)

	fmt.Println("Bueler")
}
