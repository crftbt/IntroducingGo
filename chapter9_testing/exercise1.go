// chatper 9 exercise 1 - return 0 if an empty list ([]float64{}) is passed

package main

import (
	"fmt"
)

func Average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func main() {
	v := Average([]float64{1,2})
	fmt.Println(v)
}
