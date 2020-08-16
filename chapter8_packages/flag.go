package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// define flags
	maxp := flag.Int("max", 99, "the max value")
	// parse
	flag.Parse()
	// generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}
